package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"

	"go.nlx.io/nlx/management-api/db"
)

var migrateOpts struct {
	PostgresDSN string
}

//nolint:gochecknoinits // this is the recommended way to use cobra
func init() {
	rootCmd.AddCommand(migrateCommand)
	migrateUpCommand.Flags().StringVarP(&migrateOpts.PostgresDSN, "postgres-dsn", "", "", "Postgres Connection URL")
	migrateStatusCommand.Flags().StringVarP(&migrateOpts.PostgresDSN, "postgres-dsn", "", "", "Postgres Connection URL")

	if err := migrateUpCommand.MarkFlagRequired("postgres-dsn"); err != nil {
		log.Fatal(err)
	}

	if err := migrateStatusCommand.MarkFlagRequired("postgres-dsn"); err != nil {
		log.Fatal(err)
	}

	migrateCommand.AddCommand(migrateUpCommand)
	migrateCommand.AddCommand(migrateStatusCommand)
}

func setupMigrator() *migrate.Migrate {
	resource := bindata.Resource(
		db.AssetNames(),
		db.Asset,
	)

	source, err := bindata.WithInstance(resource)
	if err != nil {
		log.Fatal(err)
	}

	migrator, err := migrate.NewWithSourceInstance("go-bindata", source, migrateOpts.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	return migrator
}

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Run the migration tool",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var migrateUpCommand = &cobra.Command{
	Use:   "up",
	Short: "Up the migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrator := setupMigrator()

		if err := migrator.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("migrations are up-to-date")

				os.Exit(0)
			}

			log.Fatal(err)
		}
	},
}

var migrateStatusCommand = &cobra.Command{
	Use:   "status",
	Short: "Show migration status",
	Run: func(cmd *cobra.Command, args []string) {
		migrator := setupMigrator()

		version, dirty, err := migrator.Version()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("dirty=%v version=%d\n", dirty, version)
	},
}
