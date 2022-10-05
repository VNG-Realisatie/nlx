// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"net/http/httputil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

const (
	mockSerialNumber string = "00000000000000000001"
	mockservicename  string = "mockservicename"
)

func TestNewRoundRobinLoadBalancer(t *testing.T) {
	organizationSerialNumber := mockSerialNumber
	serviceName := mockservicename

	inwayAddresses := []string{"mockaddress1", "mockaddress2"}

	orgCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	l, err := NewRoundRobinLoadBalancedHTTPService(
		zap.NewNop(), orgCert,
		organizationSerialNumber, serviceName,
		[]directoryapi.Inway{
			{
				Address: inwayAddresses[0],
				State:   directoryapi.Inway_STATE_UP,
			},
			{
				Address: inwayAddresses[1],
				State:   directoryapi.Inway_STATE_UP,
			},
		})

	assert.Nil(t, err)
	assert.Equal(t, inwayAddresses, l.GetInwayAddresses())

	proxy1 := l.getProxy()
	proxy2 := l.getProxy()
	proxy3 := l.getProxy()

	// Test if round robin loadbalancing algoritme is working
	assert.NotEqualf(t, proxy1, proxy2, "proxy1 and proxy2 are the same, expected proxy1 and proxy2 to be different proxies")
	assert.Equal(t, proxy1, proxy3, "proxy1 and proxy3 are not the same, expected proxy1 and proxy3 to be the same proxies")

	// Test if getProxy returns nil if slice of proxies is empty
	l.proxies = []*httputil.ReverseProxy{}
	proxy := l.getProxy()
	assert.Nilf(t, proxy, "proxy is not nil, expected proxy to be nil if there are no proxies available")
}
