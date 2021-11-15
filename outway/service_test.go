// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"net/http/httputil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
)

const (
	mockSerialNumber string = "00000000000000000001"
	mockservicename  string = "mockservicename"
)

func TestNewRoundRobinLoadBalancer(t *testing.T) {
	organizationSerialNumber := mockSerialNumber
	serviceName := mockservicename

	inwayAddresses := []string{"mockaddress1", "mockaddress2"}

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	l, err := NewRoundRobinLoadBalancedHTTPService(
		zap.NewNop(), cert,
		organizationSerialNumber, serviceName,
		[]directoryapi.Inway{
			{
				Address: inwayAddresses[0],
				State:   directoryapi.Inway_UP,
			},
			{
				Address: inwayAddresses[1],
				State:   directoryapi.Inway_UP,
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
