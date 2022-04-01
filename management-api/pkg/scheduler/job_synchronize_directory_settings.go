// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package scheduler

import (
	"context"
	"errors"
	"time"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
)

type SynchronizeDirectorySettingsJob struct {
	ctx             context.Context
	directoryClient directory.Client
	configDatabase  database.ConfigDatabase
	pollInterval    time.Duration
}

func NewSynchronizeDirectorySettingsJob(ctx context.Context, pollInterval time.Duration, directoryClient directory.Client, configDatabase database.ConfigDatabase) *SynchronizeDirectorySettingsJob {
	return &SynchronizeDirectorySettingsJob{
		ctx:             ctx,
		directoryClient: directoryClient,
		configDatabase:  configDatabase,
		pollInterval:    pollInterval,
	}
}

func (job *SynchronizeDirectorySettingsJob) Synchronize(ctx context.Context) error {
	settings, err := job.configDatabase.GetSettings(ctx)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil
		}

		return err
	}

	_, err = job.directoryClient.SetOrganizationContactDetails(ctx, &directoryapi.SetOrganizationContactDetailsRequest{
		EmailAddress: settings.OrganizationEmailAddress(),
	})
	if err != nil {
		return err
	}

	return nil
}
