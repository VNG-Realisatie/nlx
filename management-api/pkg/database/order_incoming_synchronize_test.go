// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestSynchronizeOrders(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	type args struct {
		incomingOrders []*database.IncomingOrder
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         []*domain.IncomingOrder
		wantErr      error
	}{
		"happy_flow_update_existing_without_services": {
			loadFixtures: true,
			args: args{
				incomingOrders: []*database.IncomingOrder{
					{
						Reference:   "fixture-reference",
						Description: "new-description",
						Delegator:   "00000000000000000001",
						RevokedAt:   sql.NullTime{},
						ValidFrom:   fixtureTime,
						ValidUntil:  fixtureTime,
						Services:    []database.IncomingOrderService{},
					},
				},
			},
			want: []*domain.IncomingOrder{
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference-two",
					Description: "fixture-description",
					Delegator:   "00000000000000000002",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services: []domain.IncomingOrderService{
						domain.NewIncomingOrderService("fixture-service-two", "10000000000000000002", "fixture-organization-two"),
					},
				}),
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference-three",
					Description: "fixture-description",
					Delegator:   "00000000000000000003",
					RevokedAt:   &fixtureTime,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services:    []domain.IncomingOrderService{},
				}),
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference",
					Description: "new-description",
					Delegator:   "00000000000000000001",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services:    []domain.IncomingOrderService{},
				}),
			},
			wantErr: nil,
		},
		"happy_flow_update_existing": {
			loadFixtures: true,
			args: args{
				incomingOrders: []*database.IncomingOrder{
					{
						Reference:   "fixture-reference",
						Description: "new-description",
						Delegator:   "00000000000000000001",
						RevokedAt:   sql.NullTime{},
						ValidFrom:   fixtureTime,
						ValidUntil:  fixtureTime,
						Services: []database.IncomingOrderService{
							{
								Service: "new-service-one",
								Organization: database.IncomingOrderServiceOrganization{
									Name:         "new-organization-one",
									SerialNumber: "10000000000000000001",
								},
							},
						},
					},
				},
			},
			want: []*domain.IncomingOrder{
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference-two",
					Description: "fixture-description",
					Delegator:   "00000000000000000002",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services: []domain.IncomingOrderService{
						domain.NewIncomingOrderService("fixture-service-two", "10000000000000000002", "fixture-organization-two"),
					},
				}),
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference-three",
					Description: "fixture-description",
					Delegator:   "00000000000000000003",
					RevokedAt:   &fixtureTime,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services:    []domain.IncomingOrderService{},
				}),
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference",
					Description: "new-description",
					Delegator:   "00000000000000000001",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services: []domain.IncomingOrderService{
						domain.NewIncomingOrderService("new-service-one", "10000000000000000001", "new-organization-one"),
					},
				}),
			},
			wantErr: nil,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				incomingOrders: []*database.IncomingOrder{
					{
						Reference:   "reference-one",
						Description: "description",
						Delegator:   "20000000000000000001",
						RevokedAt:   sql.NullTime{},
						ValidFrom:   fixtureTime,
						ValidUntil:  fixtureTime,
						Services: []database.IncomingOrderService{
							{
								Service: "service-one",
								Organization: database.IncomingOrderServiceOrganization{
									Name:         "organization-one",
									SerialNumber: "10000000000000000001",
								},
							},
						},
					},
					{
						Reference:   "reference-two",
						Description: "description",
						Delegator:   "20000000000000000002",
						RevokedAt:   sql.NullTime{},
						ValidFrom:   fixtureTime,
						ValidUntil:  fixtureTime,
						Services: []database.IncomingOrderService{
							{
								Service: "service-two",
								Organization: database.IncomingOrderServiceOrganization{
									Name:         "organization-two",
									SerialNumber: "10000000000000000002",
								},
							},
						},
					},
				},
			},
			want: []*domain.IncomingOrder{
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference",
					Description: "fixture-description",
					Delegator:   "00000000000000000001",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services: []domain.IncomingOrderService{
						domain.NewIncomingOrderService("fixture-service", "10000000000000000001", "fixture-organization"),
					},
				}),
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference-two",
					Description: "fixture-description",
					Delegator:   "00000000000000000002",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services: []domain.IncomingOrderService{
						domain.NewIncomingOrderService("fixture-service-two", "10000000000000000002", "fixture-organization-two"),
					},
				}),
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference-three",
					Description: "fixture-description",
					Delegator:   "00000000000000000003",
					RevokedAt:   &fixtureTime,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services:    []domain.IncomingOrderService{},
				}),
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "reference-one",
					Description: "description",
					Delegator:   "20000000000000000001",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services: []domain.IncomingOrderService{
						domain.NewIncomingOrderService("service-one", "10000000000000000001", "organization-one"),
					},
				}),
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "reference-two",
					Description: "description",
					Delegator:   "20000000000000000002",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services: []domain.IncomingOrderService{
						domain.NewIncomingOrderService("service-two", "10000000000000000002", "organization-two"),
					},
				}),
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			err := configDb.SynchronizeOrders(context.Background(), tt.args.incomingOrders)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertIncomingOrderList(t, configDb, tt.want)
			}
		})
	}
}

func assertIncomingOrderList(t *testing.T, repo database.ConfigDatabase, want []*domain.IncomingOrder) {
	got, err := repo.ListIncomingOrders(context.Background())
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.EqualValues(t, want, got)
}
