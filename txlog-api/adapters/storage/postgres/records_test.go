// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package postgresadapter_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/txlog-api/domain/record"
)

func TestCreateRecord(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		loadFixtures bool
		args         *record.NewRecordArgs
		wantErr      error
	}{
		"happy_flow_without_order": {
			loadFixtures: false,
			args: &record.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               record.IN,
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
			args: &record.NewRecordArgs{
				SourceOrganization:      "0001",
				DestinationOrganization: "0002",
				Direction:               record.IN,
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

			repo, closeDb := new(t, tt.loadFixtures)
			defer closeDb()

			model, err := record.NewRecord(tt.args)
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
		want         []*record.Record
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			want: []*record.Record{
				mustNewRecord(t, &record.NewRecordArgs{
					SourceOrganization:      "0001",
					DestinationOrganization: "0002",
					Direction:               record.IN,
					ServiceName:             "test-service",
					OrderReference:          "test-reference",
					Delegator:               "0003",
					Data:                    []byte(`{"test": "data"}`),
					TransactionID:           "abcde",
					CreatedAt:               time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
				}),
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

			got, err := repo.ListRecords(context.Background(), 100)
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, tt.want, got)
			}
		})
	}
}

func assertRecordInRepository(t *testing.T, repo record.Repository, r *record.Record) {
	require.NotNil(t, r)

	records, err := repo.ListRecords(context.Background(), 100)
	require.NoError(t, err)

	var found bool
	for _, recordModel := range records {
		if recordModel.TransactionID() == r.TransactionID() {
			found = true
			break
		}
	}

	require.Equal(t, true, found)
}

func mustNewRecord(t *testing.T, args *record.NewRecordArgs) *record.Record {
	result, err := record.NewRecord(args)
	assert.NoError(t, err)

	return result
}
