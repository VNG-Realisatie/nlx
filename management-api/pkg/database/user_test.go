// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package database_test

import (
	"context"
	"net/url"
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func Test_User(t *testing.T) {
	configDb := newPostgresConfigDatabase(t)

	t.Run("verify_credentials", func(t *testing.T) {
		t.Parallel()
		testVerifyCredentials(t, configDb)
	})

	t.Run("create_user", func(t *testing.T) {
		t.Parallel()
		testCreateUser(t, configDb)
	})
}

func testVerifyCredentials(t *testing.T, configDb database.ConfigDatabase) {
	t.Helper()

	testCases := map[string]struct {
		prepare       func(configDb database.ConfigDatabase) error
		email         string
		password      string
		expected      bool
		expectedError error
	}{
		"with_invalid_username_password": {
			email:         "invalid@credentials.com",
			password:      "bar",
			expected:      false,
			expectedError: gorm.ErrRecordNotFound,
		},
		"happy_flow": {
			prepare: func(configDb database.ConfigDatabase) error {
				_, err := configDb.CreateUser(context.Background(), "test-verify-credentials@happy-flow.com", "password", []string{})
				return err
			},
			email:         "test-verify-credentials@happy-flow.com",
			password:      "password",
			expected:      true,
			expectedError: nil,
		},
	}

	for name, tt := range testCases {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if tt.prepare != nil {
				err := tt.prepare(configDb)
				require.Nil(t, err)
			}

			actual, err := configDb.VerifyUserCredentials(context.Background(), tt.email, tt.password)
			if tt.expectedError != nil {
				require.ErrorIs(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
			}
		})
	}
}

func testCreateUser(t *testing.T, configDb database.ConfigDatabase) {
	t.Helper()

	testCases := map[string]struct {
		prepare       func(configDb database.ConfigDatabase) error
		email         string
		password      string
		roleNames     []string
		expected      *database.User
		expectedError string
	}{
		"with_existing_email_address": {
			prepare: func(configDb database.ConfigDatabase) error {
				_, err := configDb.CreateUser(
					context.Background(),
					"already@present.com",
					"password",
					[]string{},
				)
				return err
			},
			email:         "already@present.com",
			password:      "foobar",
			expectedError: database.ErrUserAlreadyExists.Error(),
		},
		"without_a_password": {
			email: "john.doe@example.com",
			expected: &database.User{
				Email: "john.doe@example.com",
			},
		},
		"happy_flow": {
			email:    "jane.doe@example.com",
			password: "foobar",
			expected: &database.User{
				Email: "jane.doe@example.com",
			},
		},
	}

	for name, tt := range testCases {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if tt.prepare != nil {
				err := tt.prepare(configDb)
				require.Nil(t, err)
			}

			actual, err := configDb.CreateUser(context.Background(), tt.email, tt.password, tt.roleNames)
			if len(tt.expectedError) > 1 {
				require.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.NotNil(t, actual)
				require.Equal(t, tt.expected.Email, actual.Email)
				require.Equal(t, tt.expected.Roles, actual.Roles)

				if len(tt.expected.Password) > 0 {
					require.Equal(t, tt.expected.Password, actual.Password)
				}
			}
		})
	}
}

func newPostgresConfigDatabase(t *testing.T) database.ConfigDatabase {
	dsn := os.Getenv("POSTGRES_DSN")
	configDb, err := database.NewPostgresConfigDatabase(dsn)
	require.NoError(t, err)

	dsnForMigrations := addQueryParamToAddress(dsn, "x-migrations-table", "management_migrations")

	err = database.PostgresPerformMigrations(dsnForMigrations)
	require.NoError(t, err)

	return configDb
}

func addQueryParamToAddress(address, key, value string) string {
	u, _ := url.Parse(address)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add(key, value)
	u.RawQuery = q.Encode()
	return u.String()
}
