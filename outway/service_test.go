// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"net/http/httputil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

const mockorg string = "mockorg"
const mockservicename string = "mockservicename"

func TestNewRoundRobinLoadBalancerExceptions(t *testing.T) {
	organisationName := mockorg
	serviceName := mockservicename
	inwayAddresses := []string{"mockaddress1", "mockaddress2"}
	certFile := filepath.Join("..", "testing", "org-nlx-test.crt")
	keyFile := filepath.Join("..", "testing", "org-nlx-test.key")
	// Test possible exceptions during RoundRoblinLoadBalancerCreation
	_, err := NewRoundRobinLoadBalancedHTTPService(zap.NewNop(), nil, certFile, keyFile, organisationName, serviceName, []string{})
	assert.Equal(t, errNoInwaysAvailable, err)

	_, err = NewRoundRobinLoadBalancedHTTPService(zap.NewNop(), nil, "invalid certfile", keyFile, organisationName, serviceName, inwayAddresses)
	if err == nil {
		t.Fatalf("result: err is nil, expected err to be set when providing invalid cert file")
	}
	assert.EqualError(t, err, "invalid certificate provided: open invalid certfile: no such file or directory")

	_, err = NewRoundRobinLoadBalancedHTTPService(zap.NewNop(), nil, certFile, "invalid key file", organisationName, serviceName, inwayAddresses)
	assert.EqualError(t, err, "invalid certificate provided: open invalid key file: no such file or directory")
}

func TestNewRoundRobinLoadBalancer(t *testing.T) {
	organisationName := mockorg
	serviceName := mockservicename
	inwayAddresses := []string{"mockaddress1", "mockaddress2"}
	certFile := filepath.Join("..", "testing", "org-nlx-test.crt")
	keyFile := filepath.Join("..", "testing", "org-nlx-test.key")
	l, err := NewRoundRobinLoadBalancedHTTPService(zap.NewNop(), nil, certFile, keyFile, organisationName, serviceName, inwayAddresses)
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
