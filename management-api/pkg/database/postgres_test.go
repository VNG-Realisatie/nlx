// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	common_tls "go.nlx.io/nlx/common/tls"
)

const nonFixturesStartID = 1
const fixturesStartID = 10001

func newFixtureCertificateBundle() (*common_tls.CertificateBundle, error) {
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	return common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)
}

func getFixtureTime(t *testing.T) time.Time {
	return getCustomFixtureTime(t, "2021-01-02T01:02:03Z")
}

func getCustomFixtureTime(t *testing.T, input string) time.Time {
	fixtureTime, err := time.Parse(time.RFC3339, input)
	require.NoError(t, err)

	return fixtureTime
}
