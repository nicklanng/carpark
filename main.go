package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/nicklanng/carpark/api"
	"github.com/nicklanng/carpark/config"
	"github.com/nicklanng/carpark/data"
	"github.com/nicklanng/carpark/data/migrations"
	"github.com/nicklanng/carpark/events"
	"github.com/nicklanng/carpark/logging"
	"github.com/nicklanng/carpark/metrics"
)

const (
	serviceName = "carpark"
)

var (
	version           string // set by LDFlags
	serverGracePeriod = 5 * time.Second
)

func main() {
	// check for version command
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 && args[0] == "version" {
		fmt.Println(version)
		return
	}

	// set up logging
	logging.SetStandardFields(serviceName, version)

	// load config
	conf, err := config.Load()
	if err != nil {
		logging.Fatal(err.Error())
	}

	// set up metrics
	metrics.Initialize(conf.StatsdEndpoint, serviceName)

	// connect to database
	db, err := data.OpenConnection(conf.DatabaseUser, conf.DatabasePassword, serviceName, conf.DatabaseHost)
	if err != nil {
		logging.Fatal(err.Error())
	}

	// migrate database
	migrationAsset := data.MakeBinDataMigration(migrations.AssetNames(), migrations.Asset)
	err = data.PerformMigration(db, "go-bindata", migrationAsset)
	if err != nil {
		log.Fatal(err.Error())
	}

	// event dispatcher
	eventChan := events.NewDispatcher(db)

	// create https server
	addr := fmt.Sprintf("%s:%s", conf.Address, conf.Port)
	routes := api.BuildRoutes(eventChan)
	server := api.StartHTTPSServer(addr, conf.TLSCertPath, conf.TLSKeyPath, routes)

	api.GracefulShutdownOnSignal([]syscall.Signal{syscall.SIGINT, syscall.SIGTERM}, func() {
		api.ShutdownHTTPServer(server, serverGracePeriod)
		db.Close()
	})
}
