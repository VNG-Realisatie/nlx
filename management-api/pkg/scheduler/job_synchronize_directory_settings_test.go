// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// nolint funlen: these are tests
package scheduler_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/scheduler"
)

func TestSynchronizeDirectorySettingsJob(t *testing.T) {
	pollInterval := 10 * time.Second

	tests := map[string]struct {
		setupMocks func(schedulerMocks)
		wantErr    error
	}{
		"when_database_error": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					GetSettings(gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			wantErr: errors.New("arbitrary error"),
		},
		"when_directory_error": {
			setupMocks: func(mocks schedulerMocks) {
				settings, err := domain.NewSettings("inway-1", "test@test.com")
				assert.NoError(t, err)

				mocks.db.
					EXPECT().
					GetSettings(gomock.Any()).
					Return(settings, nil)

				mocks.directory.
					EXPECT().
					SetOrganizationContactDetails(gomock.Any(), &directoryapi.SetOrganizationContactDetailsRequest{
						EmailAddress: "test@test.com",
					}).
					Return(nil, errors.New("arbitrary error"))
			},
			wantErr: errors.New("arbitrary error"),
		},
		"happy_flow_no_email_set": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					GetSettings(gomock.Any()).
					Return(nil, database.ErrNotFound)
			},
			wantErr: nil,
		},
		"happy_flow": {
			setupMocks: func(mocks schedulerMocks) {
				settings, err := domain.NewSettings("inway-1", "test@test.com")
				assert.NoError(t, err)

				mocks.db.
					EXPECT().
					GetSettings(gomock.Any()).
					Return(settings, nil)

				mocks.directory.
					EXPECT().
					SetOrganizationContactDetails(gomock.Any(), &directoryapi.SetOrganizationContactDetailsRequest{
						EmailAddress: "test@test.com",
					}).
					Return(&directoryapi.SetOrganizationContactDetailsResponse{}, nil)
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mocks := newMocks(t)

			tt.setupMocks(mocks)

			job := scheduler.NewSynchronizeDirectorySettingsJob(
				context.Background(),
				pollInterval,
				mocks.directory,
				mocks.db,
			)

			err := job.Synchronize(context.Background())

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
