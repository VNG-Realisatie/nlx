// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package scheduler

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/management"
)

const (
	maxRetries                                  = 3
	pollIntervalSynchronizeAccessRequestJob     = 1500 * time.Millisecond
	pollIntervalSynchronizeDirectorySettingsJob = 10 * time.Second
	maxConcurrencyAccessRequests                = 50
	synchronizationIntervalAccessRequests       = 15 * time.Second

	// jobs are unlocked after 5 minutes, let's wait at least one minute before retrying
	jobTimeout = 4 * time.Minute
)

type scheduler struct {
	logger                              *zap.Logger
	synchronizeOutgoingAccessRequestJob *SynchronizeOutgoingAccessRequestJob
	synchronizeDirectorySettingsJob     *SynchronizeDirectorySettingsJob
}

func NewScheduler(logger *zap.Logger, directoryClient directory.Client, configDatabase database.ConfigDatabase, orgCert *common_tls.CertificateBundle) *scheduler {
	synchronizeOutgoingAccessRequestJob := NewSynchronizeOutgoingAccessRequestJob(
		context.Background(),
		logger,
		pollIntervalSynchronizeAccessRequestJob,
		directoryClient,
		configDatabase,
		orgCert,
		management.NewClient,
	)

	synchronizeDirectorySettingsJob := NewSynchronizeDirectorySettingsJob(context.Background(), pollIntervalSynchronizeDirectorySettingsJob, directoryClient, configDatabase)

	return &scheduler{
		logger:                              logger,
		synchronizeOutgoingAccessRequestJob: synchronizeOutgoingAccessRequestJob,
		synchronizeDirectorySettingsJob:     synchronizeDirectorySettingsJob,
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

func (s *scheduler) RunSynchronizeOutgoingAccessRequest(ctx context.Context) {
	wgRequests := &sync.WaitGroup{}
	ticker := time.NewTicker(s.synchronizeOutgoingAccessRequestJob.pollInterval)
	sem := semaphore.NewWeighted(int64(maxConcurrencyAccessRequests))

	defer ticker.Stop()

schedulingLoop:
	for {
		select {
		case <-ctx.Done():
			break schedulingLoop
		case <-ticker.C:
			requests, err := s.synchronizeOutgoingAccessRequestJob.configDatabase.TakePendingOutgoingAccessRequests(ctx)
			if err != nil {
				return
			}

			for _, request := range requests {
				requestToSync := request
				go func() {
					jobCtx, cancel := context.WithTimeout(ctx, jobTimeout)
					defer cancel()

					if !sem.TryAcquire(1) {
						s.unlockOutgoingAccessRequest(jobCtx, requestToSync)
						return
					}

					defer sem.Release(1)

					wgRequests.Add(1)
					defer wgRequests.Done()

					err := s.synchronizeOutgoingAccessRequestJob.Synchronize(jobCtx, requestToSync)
					if err != nil {
						s.logger.Error("failed to schedule pending request", zap.Error(err))
						return
					}

					s.unlockOutgoingAccessRequest(jobCtx, requestToSync)
				}()
			}
		}
	}

	wgRequests.Wait()
}

func (s *scheduler) unlockOutgoingAccessRequest(ctx context.Context, request *database.OutgoingAccessRequest) {
	err := s.synchronizeOutgoingAccessRequestJob.configDatabase.UnlockOutgoingAccessRequest(ctx, request)
	if err != nil {
		s.logger.Error("failed to unlock pending request", zap.Error(err))
	}
}
