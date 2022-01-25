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
	maxRetries              = 3
	pollInterval            = 1500 * time.Millisecond
	maxConcurrency          = 50
	synchronizationInterval = 15 * time.Second

	// jobs are unlocked after 5 minutes, let's wait at least one minute before retrying
	jobTimeout = 4 * time.Minute
)

type scheduler struct {
	logger                              *zap.Logger
	synchronizeOutgoingAccessRequestJob *SynchronizeOutgoingAccessRequestJob
}

func NewOutgoingAccessRequestScheduler(logger *zap.Logger, directoryClient directory.Client, configDatabase database.ConfigDatabase, orgCert *common_tls.CertificateBundle) *scheduler {
	job := NewSynchronizeOutgoingAccessRequestJob(
		context.Background(),
		directoryClient,
		configDatabase,
		orgCert,
		management.NewClient,
	)

	return &scheduler{
		logger:                              logger,
		synchronizeOutgoingAccessRequestJob: job,
	}
}

func (s *scheduler) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}
	ticker := time.NewTicker(pollInterval)
	sem := semaphore.NewWeighted(int64(maxConcurrency))

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

					wg.Add(1)
					defer wg.Done()

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

	wg.Wait()
}

func (s *scheduler) unlockOutgoingAccessRequest(ctx context.Context, request *database.OutgoingAccessRequest) {
	err := s.synchronizeOutgoingAccessRequestJob.configDatabase.UnlockOutgoingAccessRequest(ctx, request)
	if err != nil {
		s.logger.Error("failed to unlock pending request", zap.Error(err))
	}
}
