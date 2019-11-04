// nolint:dupl
package configservice

import (
	"context"
	"testing"
	"time"

	"go.nlx.io/nlx/config-api/configapi"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/integration"
	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/common/process"
	"go.uber.org/zap"
)

type TestCluster struct {
	cluster *integration.ClusterV3
	DB      ConfigDatabase
	Addrs   []string
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

	db, err := NewEtcdConfigDatabase(logger, testProcess, addrs)
	if err != nil {
		t.Fatal("error constructing etcd config database", err)
	}

	return TestCluster{
		cluster: cluster,
		DB:      db,
		Addrs:   addrs,
	}
}

func TestNewEtcdConfigDatabase(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	assert.NotNil(t, cluster.DB)
}

func TestListServices(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockService := &configapi.Service{
		Name: "my-service",
	}

	anotherMockService := &configapi.Service{
		Name: "another-service",
	}

	err := cluster.DB.CreateService(ctx, mockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	err = cluster.DB.CreateService(ctx, anotherMockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	services, err := cluster.DB.ListServices(ctx)
	if err != nil {
		t.Fatal("error listing services", err)
	}

	assert.Equal(t, []*configapi.Service{anotherMockService, mockService}, services)
}

func TestCreateGetService(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockService := &configapi.Service{
		Name: "my-service",
	}

	err := cluster.DB.CreateService(ctx, mockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	service, err := cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Equal(t, service, mockService)
}

func TestUpdateService(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockService := &configapi.Service{
		Name:        "my-service",
		EndpointURL: "https://somewhere/",
	}

	updatedMockService := &configapi.Service{
		Name:        "my-service",
		EndpointURL: "https://somewhere-else/",
	}

	err := cluster.DB.CreateService(ctx, mockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	service, err := cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Equal(t, service.EndpointURL, mockService.EndpointURL)

	err = cluster.DB.UpdateService(ctx, "my-service", updatedMockService)
	if err != nil {
		t.Fatal("error updating service", err)
	}

	service, err = cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Equal(t, service.EndpointURL, updatedMockService.EndpointURL)
}

func TestDeleteService(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockService := &configapi.Service{
		Name:        "my-service",
		EndpointURL: "https://somewhere/",
	}

	err := cluster.DB.CreateService(ctx, mockService)
	if err != nil {
		t.Fatal("error creating service", err)
	}

	service, err := cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Equal(t, service, mockService)

	err = cluster.DB.DeleteService(ctx, "my-service")
	if err != nil {
		t.Fatal("error deleting service", err)
	}

	service, err = cluster.DB.GetService(ctx, "my-service")
	if err != nil {
		t.Fatal("error getting service", err)
	}

	assert.Nil(t, service)
}

func TestListInways(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockInway := &configapi.Inway{
		Name: "inway42.test",
	}

	anotherMockInway := &configapi.Inway{
		Name: "inway43.test",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	err = cluster.DB.CreateInway(ctx, anotherMockInway)
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	inways, err := cluster.DB.ListInways(ctx)
	if err != nil {
		t.Fatal("error listing inways", err)
	}

	assert.Equal(t, []*configapi.Inway{mockInway, anotherMockInway}, inways)
}

func TestCreateGetInway(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockInway := &configapi.Inway{
		Name: "my-inway",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	service, err := cluster.DB.GetInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error getting inway", err)
	}

	assert.Equal(t, service, mockInway)
}

func TestUpdateInway(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockInway := &configapi.Inway{
		Name: "my-inway",
	}

	mockUpdatedInway := &configapi.Inway{
		Name: "my-inway",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	inway, err := cluster.DB.GetInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error getting inway", err)
	}

	assert.Equal(t, mockInway, inway)

	err = cluster.DB.UpdateInway(ctx, "my-inway", mockUpdatedInway)
	if err != nil {
		t.Fatal("error updating inway", err)
	}

	inway, err = cluster.DB.GetInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error getting inway", err)
	}

	assert.Equal(t, inway, mockUpdatedInway)
}

func TestDeleteInway(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockInway := &configapi.Inway{
		Name: "my-inway",
	}

	err := cluster.DB.CreateInway(ctx, mockInway)
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	inway, err := cluster.DB.GetInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error getting inway", err)
	}

	assert.Equal(t, inway, mockInway)

	err = cluster.DB.DeleteInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error deleting inway", err)
	}

	inway, err = cluster.DB.GetInway(ctx, "my-inway")
	if err != nil {
		t.Fatal("error getting inway", err)
	}

	assert.Nil(t, inway)
}

func TestPutGetInsight(t *testing.T) {
	cluster := newTestCluster(t)
	defer cluster.Terminate(t)

	ctx := context.Background()

	mockInsightConfiguration := &configapi.InsightConfiguration{
		IrmaServerURL: "http://irma-url.com",
		InsightAPIURL: "http://insight-url.com",
	}

	err := cluster.DB.PutInsightConfiguration(ctx, mockInsightConfiguration)
	if err != nil {
		t.Fatal("error putting insight configuration", err)
	}

	insightConfig, err := cluster.DB.GetInsightConfiguration(ctx)
	if err != nil {
		t.Fatal("error getting insight configuration", err)
	}

	assert.Equal(t, mockInsightConfiguration, insightConfig)

}
