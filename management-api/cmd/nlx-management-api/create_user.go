package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"go.nlx.io/nlx/management-api/pkg/database"
)

const invalidArgumentCode = 128

var createUserOpts struct {
	Email       string
	Password    string
	Roles       []string
	PostgresDSN string
}

//nolint:gochecknoinits // this is the recommended way to use cobra
func init() {
	rootCmd.AddCommand(createUserCommand)

	createUserCommand.Flags().StringVarP(&createUserOpts.Email, "email", "e", "", "User email")
	createUserCommand.Flags().StringVarP(&createUserOpts.Password, "password", "", "", "User password")
	createUserCommand.Flags().StringArrayVarP(&createUserOpts.Roles, "role", "r", []string{}, "User roles")
	createUserCommand.Flags().StringVarP(&createUserOpts.PostgresDSN, "postgres-dsn", "p", "", "Postgres Connection URL")

	if err := createUserCommand.MarkFlagRequired("email"); err != nil {
		log.Fatal(err)
	}

	if err := createUserCommand.MarkFlagRequired("role"); err != nil {
		log.Fatal(err)
	}

	if err := createUserCommand.MarkFlagRequired("postgres-dsn"); err != nil {
		log.Fatal(err)
	}
}

var createUserCommand = &cobra.Command{
	Use:   "create-user",
	Short: "Create a single user with one or multiple roles",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.NewPostgresConfigDatabase(createUserOpts.PostgresDSN)
		if err != nil {
			log.Fatal(err)
		}

		user, err := db.CreateUser(
			context.Background(),
			createUserOpts.Email,
			createUserOpts.Password,
			createUserOpts.Roles,
		)
		if err != nil {
			if errors.Is(err, database.ErrUserAlreadyExists) {
				log.Println(err)
				os.Exit(invalidArgumentCode)
			}

			if errors.Is(err, &database.RoleNotFoundError{}) {
				log.Fatalf(
					"unable to create user because role '%s' was not found\n",
					err.(*database.RoleNotFoundError).RoleName,
				)
			}

			log.Fatal(err)
		}

		fmt.Printf("user created with ID: %d\n", user.ID)
	},
}
