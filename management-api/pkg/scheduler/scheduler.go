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
	maxRetries     = 3
	maxConcurrency = 4
	pollInterval   = 1500 * time.Millisecond

	// jobs are unlocked after 5 minutes, let's wait at least one minute before retrying
	jobTimeout = 4 * time.Minute
)

type scheduler struct {
	logger                              *zap.Logger
	requests                            chan *database.OutgoingAccessRequest
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
		requests:                            make(chan *database.OutgoingAccessRequest),
		synchronizeOutgoingAccessRequestJob: job,
	}
}

func (s *scheduler) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}
	sem := semaphore.NewWeighted(int64(maxConcurrency))
	ticker := time.NewTicker(pollInterval)

	defer ticker.Stop()

schedulingLoop:
	for {
		select {
		case <-ctx.Done():
			break schedulingLoop
		case <-ticker.C:
			go func() {
				if !sem.TryAcquire(1) {
					return
				}

				wg.Add(1)

				defer sem.Release(1)
				defer wg.Done()

				jobCtx, cancel := context.WithTimeout(ctx, jobTimeout)

				defer cancel()

				err := s.synchronizeOutgoingAccessRequestJob.Run(jobCtx)
				if err != nil {
					s.logger.Error("failed to schedule pending request", zap.Error(err))
				}
			}()
		}
	}

	wg.Wait()
}
