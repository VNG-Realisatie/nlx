// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/coreos/etcd/clientv3"
)

type Service struct {
	Name                  string                        `json:"name,omitempty"`
	EndpointURL           string                        `json:"endpointURL,omitempty"`
	DocumentationURL      string                        `json:"documentationURL,omitempty"`
	APISpecificationURL   string                        `json:"apiSpecificationURL,omitempty"`
	Internal              bool                          `json:"internal,omitempty"`
	TechSupportContact    string                        `json:"techSupportContact,omitempty"`
	PublicSupportContact  string                        `json:"publicSupportContact,omitempty"`
	AuthorizationSettings *ServiceAuthorizationSettings `json:"authorizationSettings,omitempty"`
	Inways                []string                      `json:"inways,omitempty"`
}

type ServiceAuthorizationSettings struct {
	Mode string `json:"mode,omitempty"`
}

// ListServices returns a list of services
func (db ETCDConfigDatabase) ListServices(ctx context.Context) ([]*Service, error) {
	key := path.Join(db.pathPrefix, "services")
	if !strings.HasSuffix(key, "/") {
		key += "/"
	}

	getResponse, err := db.etcdCli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	services := []*Service{}

	for _, kv := range getResponse.Kvs {
		service := &Service{}
		err := json.Unmarshal(kv.Value, service)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

// GetService returns a specific service by name
func (db ETCDConfigDatabase) GetService(ctx context.Context, name string) (*Service, error) {
	key := path.Join(db.pathPrefix, "services", name)

	values, err := db.etcdCli.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if values.Count == 0 {
		return nil, nil
	}

	service := &Service{}
	err = json.Unmarshal(values.Kvs[0].Value, service)

	if err != nil {
		return nil, err
	}

	return service, nil
}

// CreateService creates a new service
func (db ETCDConfigDatabase) CreateService(ctx context.Context, service *Service) error {
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
//nolint:dupl // test method
func (db ETCDConfigDatabase) UpdateService(ctx context.Context, name string, service *Service) error {
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

// FilterServices returns an array with only services for the given inway
func FilterServices(services []*Service, inway *Inway) []*Service {
	result := []*Service{}

	for _, service := range services {
		if contains(service.Inways, inway.Name) {
			result = append(result, service)
		}
	}

	return result
}

func contains(s []string, v string) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}

	return false
}
