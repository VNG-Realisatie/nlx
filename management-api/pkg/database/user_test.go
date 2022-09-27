// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestVerifyCredentials(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		email    string
		password string
	}

	testCases := map[string]struct {
		loadFixtures  bool
		args          args
		expected      bool
		expectedError error
	}{
		"with_invalid_username_password": {
			loadFixtures: false,
			args: args{
				email:    "invalid@credentials.com",
				password: "bar",
			},
			expected:      false,
			expectedError: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				email:    "fixture@example.com",
				password: "password",
			},
			expected:      true,
			expectedError: nil,
		},
	}

	for name, tt := range testCases {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			actual, err := configDb.VerifyUserCredentials(context.Background(), tt.args.email, tt.args.password)
			if tt.expectedError != nil {
				require.ErrorIs(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		email     string
		password  string
		roleNames []string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		expected     *database.User
		expectedErr  error
	}{
		"with_existing_email_address": {
			loadFixtures: true,
			args: args{
				email:    "fixture@example.com",
				password: "foobar",
				roleNames: []string{
					"admin",
				},
			},
			expectedErr: database.ErrUserAlreadyExists,
		},
		"happy_flow_without_a_password": {
			loadFixtures: true,
			args: args{
				email: "john.doe@example.com",
				roleNames: []string{
					"admin",
				},
			},
			expected: &database.User{
				Email: "john.doe@example.com",
				Roles: []database.Role{
					{
						Code: "admin",
					},
				},
			},
			expectedErr: nil,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				email:    "jane.doe@example.com",
				password: "foobar",
				roleNames: []string{
					"admin",
				},
			},
			expected: &database.User{
				Email: "jane.doe@example.com",
				Roles: []database.Role{
					{
						Code: "admin",
					},
				},
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			_, err := configDb.CreateUser(context.Background(), tt.args.email, tt.args.password, tt.args.roleNames)
			require.ErrorIs(t, err, tt.expectedErr)

			if err == nil {
				newUser, err := configDb.GetUser(context.Background(), tt.args.email)
				require.NoError(t, err)

				require.Equal(t, tt.expected.Email, newUser.Email)

				for i, role := range tt.expected.Roles {
					require.Equal(t, role.Code, newUser.Roles[i].Code)
				}
			}
		})
	}
}
