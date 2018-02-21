package data

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	bindata "github.com/mattes/migrate/source/go-bindata"
	"github.com/nicklanng/carpark/logging"
)

// MakeBinDataMigration creates a migration source from files packed into the binary with go-bindata
func MakeBinDataMigration(assetNames []string, assetLoader func(name string) ([]byte, error)) *bindata.AssetSource {
	return bindata.Resource(assetNames, assetLoader)
}

// PerformMigration will update the database using migrations in the asset source
func PerformMigration(database *sql.DB, sourceType string, s *bindata.AssetSource) error {
	sourceDriver, err := bindata.WithInstance(s)
	if err != nil {
		logging.Fatal(err.Error())
	}

	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithInstance(sourceType, sourceDriver, "mysql", driver)
	if err != nil {
		return err
	}

	// get version of schema currently in the database
	currentVersion, dirty, err := migration.Version()
	if err != nil {
		switch err.Error() {
		case "no migration":
			logging.Info("No existing schema")
		default:
			return err
		}
	} else {
		logging.Info(fmt.Sprintf("Schema is at version %d", currentVersion))
	}

	if dirty {
		return errors.New("Dirty schema - please manually fix")
	}

	for {
		// find the next available version in the code
		nextVersion, err := sourceDriver.First()
		if err != nil {
			return err
		}

		// if no update required, break out
		if currentVersion == nextVersion {
			logging.Info("Schema up-to-date")
			break
		}

		logging.Info(fmt.Sprintf("Found schema migration %d", nextVersion))

		// apply next migration
		err = migration.Steps(int(nextVersion))
		if err != nil {
			return err
		}
		logging.Info(fmt.Sprintf("Applied schema migration %d", nextVersion))

		currentVersion = nextVersion
	}

	return nil
}
