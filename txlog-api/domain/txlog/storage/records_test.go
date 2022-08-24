// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package storage_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/txlog-api/domain"
)

func TestCreateRecord(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		loadFixtures bool
		args         *domain.NewRecordArgs
		wantErr      error
	}{
		"happy_flow_without_order": {
			loadFixtures: false,
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
				DataSubjects:            map[string]string{"foo": "bar"},
			},
			wantErr: nil,
		},
		"happy_flow": {
			loadFixtures: false,
			args: &domain.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               domain.IN,
				ServiceName:             "test-service",
				OrderReference:          "test-reference",
				Delegator:               "0003",
				Data:                    []byte(`{"test": "value"}`),
				TransactionID:           "abcde",
				CreatedAt:               time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
				DataSubjects:            map[string]string{"foo": "bar"},
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo, close := new(t, tt.loadFixtures)
			defer close()

			model, err := domain.NewRecord(tt.args)
			require.NoError(t, err)

			err = repo.CreateRecord(context.Background(), model)
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assertRecordInRepository(t, repo, model)
			}
		})
	}
}

func TestListRecords(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		loadFixtures bool
		want         []*domain.NewRecordArgs
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			want: []*domain.NewRecordArgs{
				{
					SourceOrganization:      "0001",
					DestinationOrganization: "0002",
					Direction:               domain.IN,
					ServiceName:             "test-service",
					OrderReference:          "test-reference",
					Delegator:               "0003",
					Data:                    []byte(`{"test": "data"}`),
					TransactionID:           "abcde",
					CreatedAt:               time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo, close := new(t, tt.loadFixtures)
			defer close()

			want := make([]*domain.Record, len(tt.want))

			for i, s := range tt.want {
				var err error
				want[i], err = domain.NewRecord(s)
				require.NoError(t, err)
			}

			got, err := repo.ListRecords(context.Background(), 100)
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, want, got)
			}
		})
	}
}
