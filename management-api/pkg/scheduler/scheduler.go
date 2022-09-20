// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package scheduler

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
)

const (
	pollIntervalSynchronizeDirectorySettingsJob = 10 * time.Second
)

type scheduler struct {
	logger                          *zap.Logger
	synchronizeDirectorySettingsJob *SynchronizeDirectorySettingsJob
}

func NewScheduler(logger *zap.Logger, directoryClient directory.Client, configDatabase database.ConfigDatabase) *scheduler {
	synchronizeDirectorySettingsJob := NewSynchronizeDirectorySettingsJob(context.Background(), pollIntervalSynchronizeDirectorySettingsJob, directoryClient, configDatabase)

	return &scheduler{
		logger:                          logger,
		synchronizeDirectorySettingsJob: synchronizeDirectorySettingsJob,
	}
}

func (s *scheduler) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func() {
		s.RunSynchronizeDirectorySettings(ctx)
		wg.Done()
	}()

	wg.Wait()
}

func (s *scheduler) RunSynchronizeDirectorySettings(ctx context.Context) {
	ticker := time.NewTicker(s.synchronizeDirectorySettingsJob.pollInterval)

	defer ticker.Stop()

schedulingLoop:
	for {
		select {
		case <-ctx.Done():
			break schedulingLoop
		case <-ticker.C:
			err := s.synchronizeDirectorySettingsJob.Synchronize(ctx)
			if err != nil {
				s.logger.Error("failed to schedule synchronize directory settings", zap.Error(err))
			}
		}
	}
}
