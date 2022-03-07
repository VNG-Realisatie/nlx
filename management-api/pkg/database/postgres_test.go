// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	common_tls "go.nlx.io/nlx/common/tls"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

const nonFixturesStartID = 1
const fixturesStartID = 10001

func newFixtureCertificateBundle() (*common_tls.CertificateBundle, error) {
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	return common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
}

func getFixtureTime(t *testing.T) time.Time {
	return getCustomFixtureTime(t, "2021-01-02T01:02:03Z")
}

func getCustomFixtureTime(t *testing.T, input string) time.Time {
	fixtureTime, err := time.Parse(time.RFC3339, input)
	require.NoError(t, err)

	return fixtureTime
}
