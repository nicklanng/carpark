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
	"github.com/nicklanng/carpark/projection"
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
	db, eventListener, err := data.OpenConnection(conf.DatabaseUser, conf.DatabasePassword, serviceName, conf.DatabaseHost)
	if err != nil {
		logging.Fatal(err.Error())
	}

	// migrate database
	migrationAsset := data.MakeBinDataMigration(migrations.AssetNames(), migrations.Asset)
	err = data.PerformMigration(db, "go-bindata", migrationAsset)
	if err != nil {
		log.Fatal(err.Error())
	}

	// TODO: Read in all events from DB

	// event dispatcher
	eventChan := events.NewDispatcher(db)

	// create state in memory
	state := projection.NewState()

	// listen to notifications from database and process event
	projection.CreateEventListener(state, eventListener)

	// create https server
	addr := fmt.Sprintf("%s:%s", conf.Address, conf.Port)
	routes := api.BuildRoutes(state, eventChan)
	server := api.StartHTTPSServer(addr, conf.TLSCertPath, conf.TLSKeyPath, routes)

	api.GracefulShutdownOnSignal([]syscall.Signal{syscall.SIGINT, syscall.SIGTERM}, func() {
		api.ShutdownHTTPServer(server, serverGracePeriod)
		db.Close()
		eventListener.Close()
	})
}
