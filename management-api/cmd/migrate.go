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
	sessionstore_migrations "go.nlx.io/nlx/management-api/pkg/oidc/pgsessionstore/migrations"
)

var migrateOpts struct {
	PostgresDSN     string
	EnableBasicAuth bool
}

//nolint:gochecknoinits // this is the recommended way to use cobra
func init() {
	migrateUpCommand.Flags().StringVarP(&migrateOpts.PostgresDSN, "postgres-dsn", "", "", "Postgres Connection URL")
	migrateUpCommand.Flags().BoolVarP(&migrateOpts.EnableBasicAuth, "enable-basic-auth", "", false, "Enable HTTP basic authentication and disable OIDC")
	migrateStatusCommand.Flags().StringVarP(&migrateOpts.PostgresDSN, "postgres-dsn", "", "", "Postgres Connection URL")
	migrateStatusCommand.Flags().BoolVarP(&migrateOpts.EnableBasicAuth, "enable-basic-auth", "", false, "Enable HTTP basic authentication and disable OIDC")

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

		if !migrateOpts.EnableBasicAuth {
			err = sessionstore_migrations.PerformMigrations(migrateOpts.PostgresDSN)
			if err != nil {
				log.Fatal(err)
			}
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

		fmt.Printf("nlx_management: dirty=%v version=%d\n", dirty, version)

		if !migrateOpts.EnableBasicAuth {
			version, dirty, err = sessionstore_migrations.MigrationStatus(migrateOpts.PostgresDSN)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("http_sessions: dirty=%v version=%d\n", dirty, version)
		}
	},
}
