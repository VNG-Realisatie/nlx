// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database_test

import (
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/integration"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/util/clock"
)

type TestCluster struct {
	cluster *integration.ClusterV3
	DB      database.ConfigDatabase
	Addrs   []string
	Clock   *clock.FakeClock
}

func (tc TestCluster) Terminate(t *testing.T) {
	tc.cluster.Terminate(t)
}

func (tc TestCluster) GetClient(t *testing.T) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   tc.Addrs,
		DialTimeout: time.Second,
	})

	if err != nil {
		t.Fatal("could not construct etcd client", err)
	}

	return cli
}

func newTestCluster(t *testing.T) TestCluster {
	cluster := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	addrs := []string{cluster.Members[0].GRPCAddr()}

	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	c := clock.NewFakeClock(time.Now())

	db, err := database.NewEtcdConfigDatabase(logger, testProcess, addrs, c)
	if err != nil {
		t.Fatal("error constructing etcd config database", err)
	}

	return TestCluster{
		cluster: cluster,
		DB:      db,
		Addrs:   addrs,
		Clock:   c,
	}
}

func TestNewEtcdConfigDatabase(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	assert.NotNil(t, cluster.DB)
}
