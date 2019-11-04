// nolint:dupl
package configservice

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/config-api/configapi"
)

const PREFIX = "nlx"

// ConfigDatabase is the interface for a configuration database
type ConfigDatabase interface {
	ListServices(ctx context.Context) ([]*configapi.Service, error)
	GetService(ctx context.Context, name string) (*configapi.Service, error)
	CreateService(ctx context.Context, service *configapi.Service) error
	UpdateService(ctx context.Context, name string, service *configapi.Service) error
	DeleteService(ctx context.Context, name string) error
	ListInways(ctx context.Context) ([]*configapi.Inway, error)
	GetInway(ctx context.Context, name string) (*configapi.Inway, error)
	CreateInway(ctx context.Context, inway *configapi.Inway) error
	UpdateInway(ctx context.Context, name string, inway *configapi.Inway) error
	DeleteInway(ctx context.Context, name string) error
	PutInsightConfiguration(ctx context.Context, configuration *configapi.InsightConfiguration) error
	GetInsightConfiguration(ctx context.Context) (*configapi.InsightConfiguration, error)
}

// ETCDConfigDatabase is the etcd implementation of ConfigDatabase
type ETCDConfigDatabase struct {
	pathPrefix string
	etcdCli    *clientv3.Client
	logger     *zap.Logger
}

// NewEtcdConfigDatabase constructs a new EtcdConfigDatabase
func NewEtcdConfigDatabase(logger *zap.Logger, p *process.Process, connectionStrings []string) (ConfigDatabase, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   connectionStrings,
		DialTimeout: time.Second,
	})
	if err != nil {
		return nil, err
	}

	pathPrefix := PREFIX
	if !strings.HasPrefix(PREFIX, "/") {
		pathPrefix = "/" + pathPrefix
	}

	p.CloseGracefully(cli.Close)
	return &ETCDConfigDatabase{
		pathPrefix: pathPrefix,
		etcdCli:    cli,
		logger:     logger,
	}, nil

}

// ListServices returns a list of services
func (db ETCDConfigDatabase) ListServices(ctx context.Context) ([]*configapi.Service, error) {
	key := path.Join(db.pathPrefix, "services")
	if !strings.HasSuffix(key, "/") {
		key += "/"
	}

	getResponse, err := db.etcdCli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	services := []*configapi.Service{}
	for _, kv := range getResponse.Kvs {
		service := &configapi.Service{}
		err := json.Unmarshal(kv.Value, service)
		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

// GetService returns a specific service by name
func (db ETCDConfigDatabase) GetService(ctx context.Context, name string) (*configapi.Service, error) {
	key := path.Join(db.pathPrefix, "services", name)

	values, err := db.etcdCli.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if values.Count == 0 {
		return nil, nil
	}

	service := &configapi.Service{}
	err = json.Unmarshal(values.Kvs[0].Value, service)
	if err != nil {
		return nil, err
	}

	return service, nil
}

// CreateService creates a new service
func (db ETCDConfigDatabase) CreateService(ctx context.Context, service *configapi.Service) error {
	key := path.Join(db.pathPrefix, "services", service.Name)

	data, err := json.Marshal(&service)
	if err != nil {
		return err
	}

	_, err = db.etcdCli.Put(ctx, key, string(data))
	if err != nil {
		return err
	}

	return nil
}

// UpdateService updates an existing service
func (db ETCDConfigDatabase) UpdateService(ctx context.Context, name string, service *configapi.Service) error {
	key := path.Join(db.pathPrefix, "services", name)

	value, err := db.etcdCli.Get(ctx, key)
	if err != nil {
		return err
	}

	if value.Count == 0 {
		return fmt.Errorf("not found")
	}

	data, err := json.Marshal(&service)
	if err != nil {
		return err
	}

	_, err = db.etcdCli.Put(ctx, key, string(data))
	if err != nil {
		return err
	}

	return nil
}

// DeleteService deletes a specific service
func (db ETCDConfigDatabase) DeleteService(ctx context.Context, name string) error {
	key := path.Join(db.pathPrefix, "services", name)

	_, err := db.etcdCli.Delete(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

// ListInways returns a list of inways
func (db ETCDConfigDatabase) ListInways(ctx context.Context) ([]*configapi.Inway, error) {
	key := path.Join(db.pathPrefix, "inways")
	if !strings.HasSuffix(key, "/") {
		key += "/"
	}

	getResponse, err := db.etcdCli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	inways := []*configapi.Inway{}
	for _, kv := range getResponse.Kvs {
		inway := &configapi.Inway{}
		err := json.Unmarshal(kv.Value, inway)
		if err != nil {
			return nil, err
		}

		inways = append(inways, inway)
	}

	return inways, nil
}

// GetInway returns a specific inway by name
func (db ETCDConfigDatabase) GetInway(ctx context.Context, name string) (*configapi.Inway, error) {
	key := path.Join(db.pathPrefix, "inways", name)

	values, err := db.etcdCli.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if values.Count == 0 {
		return nil, nil
	}

	inway := &configapi.Inway{}
	err = json.Unmarshal(values.Kvs[0].Value, inway)
	if err != nil {
		return nil, err
	}

	return inway, nil
}

// CreateInway creates a new inway
func (db ETCDConfigDatabase) CreateInway(ctx context.Context, inway *configapi.Inway) error {
	key := path.Join(db.pathPrefix, "inways", inway.Name)

	data, err := json.Marshal(&inway)
	if err != nil {
		return err
	}

	_, err = db.etcdCli.Put(ctx, key, string(data))
	if err != nil {
		return err
	}

	return nil
}

// UpdateInway updates an existing inway
func (db ETCDConfigDatabase) UpdateInway(ctx context.Context, name string, inway *configapi.Inway) error {
	key := path.Join(db.pathPrefix, "inways", name)

	value, err := db.etcdCli.Get(ctx, key)
	if err != nil {
		return err
	}

	if value.Count == 0 {
		return fmt.Errorf("not found")
	}

	data, err := json.Marshal(&inway)
	if err != nil {
		return err
	}

	_, err = db.etcdCli.Put(ctx, key, string(data))
	if err != nil {
		return err
	}

	return nil
}

// DeleteInway deletes a specific inway
func (db ETCDConfigDatabase) DeleteInway(ctx context.Context, name string) error {
	key := path.Join(db.pathPrefix, "inways", name)

	_, err := db.etcdCli.Delete(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

// PutInsight sets the insight configuration
func (db ETCDConfigDatabase) PutInsightConfiguration(ctx context.Context, insightConfiguration *configapi.InsightConfiguration) error {
	key := path.Join(db.pathPrefix, "insight-configuration")

	data, err := json.Marshal(&insightConfiguration)
	if err != nil {
		return err
	}

	_, err = db.etcdCli.Put(ctx, key, string(data))
	if err != nil {
		return err
	}

	return nil
}

// GetInsight returns the insight configuration
func (db ETCDConfigDatabase) GetInsightConfiguration(ctx context.Context) (*configapi.InsightConfiguration, error) {
	key := path.Join(db.pathPrefix, "insight-configuration")

	values, err := db.etcdCli.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if values.Count == 0 {
		return nil, nil
	}

	insightConfiguration := &configapi.InsightConfiguration{}
	err = json.Unmarshal(values.Kvs[0].Value, insightConfiguration)
	if err != nil {
		return nil, err
	}

	return insightConfiguration, nil
}
