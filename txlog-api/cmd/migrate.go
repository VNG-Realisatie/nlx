// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package cmd

import (
	"fmt"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"

	pgadapter "go.nlx.io/nlx/txlog-api/adapters/storage/postgres"
)

var migrateOpts struct {
	PostgresDSN string
}

//nolint:gochecknoinits // this is the recommended way to use cobra
func init() {
	migrateUpCommand.Flags().StringVarP(&migrateOpts.PostgresDSN, "postgres-dsn", "", "", "Postgres Connection URL")

	if err := migrateUpCommand.MarkFlagRequired("postgres-dsn"); err != nil {
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
		err := pgadapter.PostgreSQLPerformMigrations(migrateOpts.PostgresDSN)
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
		version, dirty, err := pgadapter.PostgresMigrationStatus(migrateOpts.PostgresDSN)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("dirty=%v version=%d\n", dirty, version)
	},
}
