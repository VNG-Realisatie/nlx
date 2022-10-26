// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package pgsessionstore_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCleanup(t *testing.T) {
	t.Parallel()

	ss := New(t, secret)
	defer ss.Close()

	// Start the cleanup goroutine.
	interval := time.Millisecond * 100
	defer ss.StopCleanup(ss.StartCleanup(interval))

	req, err := http.NewRequest("GET", "http://www.example.com", http.NoBody)
	assert.NoError(t, err, "failed to create request")

	session, err := ss.Get(req, "newsess")
	assert.NoError(t, err, "failed to create session")

	// Expire the session.
	session.Options.MaxAge = 1

	m := make(http.Header)

	err = ss.Save(req, headerOnlyResponseWriter(m), session)
	assert.NoError(t, err, "failed to save session")

	// Give the ticker a moment to run.
	time.Sleep(interval * 2)

	// Count expired sessions. We should get a count of zero back.
	count, err := ss.Querier.CountExpiredSessions(context.Background())
	assert.NoError(t, err, "failed to retrieve expired sessions")

	assert.Zero(t, count, "ticker did not delete expired sessions")
}
