// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package pgsessionstore

import (
	"context"
	"fmt"
	"time"
)

var defaultInterval = time.Minute * 5

// Cleanup runs a background goroutine every interval that deletes expired
// sessions from the database.
//
// The design is based on https://github.com/yosssi/boltstore
func (db *PGStore) StartCleanup(interval time.Duration) (quit chan<- struct{}, done <-chan struct{}) {
	if interval <= 0 {
		interval = defaultInterval
	}

	q, d := make(chan struct{}), make(chan struct{})
	go db.startCleanup(interval, q, d)

	return q, d
}

// StopCleanup stops the background cleanup from running.
func (db *PGStore) StopCleanup(quit chan<- struct{}, done <-chan struct{}) {
	quit <- struct{}{}

	<-done
}

// cleanup deletes expired sessions at set intervals.
func (db *PGStore) startCleanup(interval time.Duration, quit <-chan struct{}, done chan<- struct{}) {
	ticker := time.NewTicker(interval)

	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case <-quit:
			// Handle the quit signal.
			done <- struct{}{}
			return
		case <-ticker.C:
			// Delete expired sessions on each tick.
			err := db.Querier.DeleteExpiredSessions(context.Background())
			if err != nil {
				db.logger.Error(fmt.Sprintf("oidc_session_store: unable to delete expired sessions: %v", err))
			}
		}
	}
}
