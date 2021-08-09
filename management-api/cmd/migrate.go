// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package cmd

import (
	"fmt"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"

	"go.nlx.io/nlx/management-api/pkg/database"
)

var migrateOpts struct {
	PostgresDSN string
}

//nolint:gochecknoinits // this is the recommended way to use cobra
func init() {
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

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Run the migration tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var migrateUpCommand = &cobra.Command{
	Use:   "up",
	Short: "Up the migrations",
	Run: func(cmd *cobra.Command, args []string) {
		err := database.PostgresPerformMigrations(migrateOpts.PostgresDSN)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	},
}

var migrateStatusCommand = &cobra.Command{
	Use:   "status",
	Short: "Show migration status",
	Run: func(cmd *cobra.Command, args []string) {
		version, dirty, err := database.PostgresMigrationStatus(migrateOpts.PostgresDSN)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("dirty=%v version=%d\n", dirty, version)
	},
}
