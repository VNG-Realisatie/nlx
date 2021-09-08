// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

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
		args          args
		expected      bool
		expectedError error
	}{
		"with_invalid_username_password": {
			args: args{
				email:    "invalid@credentials.com",
				password: "bar",
			},
			expected:      false,
			expectedError: gorm.ErrRecordNotFound,
		},
		"happy_flow": {
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

			configDb, close := newConfigDatabase(t, t.Name())
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
		args        args
		expected    *database.User
		expectedErr error
	}{
		"with_existing_email_address": {
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

			configDb, close := newConfigDatabase(t, t.Name())
			defer close()

			actual, err := configDb.CreateUser(context.Background(), tt.args.email, tt.args.password, tt.args.roleNames)
			require.ErrorIs(t, err, tt.expectedErr)

			if err == nil {
				require.Equal(t, tt.expected.Email, actual.Email)

				for i, role := range tt.expected.Roles {
					require.Equal(t, role.Code, actual.Roles[i].Code)
				}
			}
		})
	}
}
