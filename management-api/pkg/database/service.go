// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"path"
)

type Service struct {
	Name                 string   `json:"name,omitempty"`
	EndpointURL          string   `json:"endpointURL,omitempty"`
	DocumentationURL     string   `json:"documentationURL,omitempty"`
	APISpecificationURL  string   `json:"apiSpecificationURL,omitempty"`
	Internal             bool     `json:"internal,omitempty"`
	TechSupportContact   string   `json:"techSupportContact,omitempty"`
	PublicSupportContact string   `json:"publicSupportContact,omitempty"`
	Inways               []string `json:"inways,omitempty"`
}

const servicesKey = "services"

// ListServices returns a list of services
func (db ETCDConfigDatabase) ListServices(ctx context.Context) ([]*Service, error) {
	services := []*Service{}

	if err := db.list(ctx, servicesKey, &services); err != nil {
		return nil, err
	}

	return services, nil
}

// GetService returns a specific service by name
func (db ETCDConfigDatabase) GetService(ctx context.Context, name string) (*Service, error) {
	key := path.Join(servicesKey, name)
	service := &Service{}

	if err := db.get(ctx, key, service); err != nil {
		return nil, err
	}

	return service, nil
}

// CreateService creates a new service
func (db ETCDConfigDatabase) CreateService(ctx context.Context, service *Service) error {
	key := path.Join("services", service.Name)

	if err := db.put(ctx, key, service); err != nil {
		return err
	}

	return nil
}

// UpdateService updates an existing service
func (db ETCDConfigDatabase) UpdateService(ctx context.Context, name string, service *Service) error {
	if _, err := db.GetService(ctx, name); err != nil {
		return err
	}

	key := path.Join(servicesKey, name)

	if err := db.put(ctx, key, service); err != nil {
		return err
	}

	return nil
}

// DeleteService deletes a specific service
func (db ETCDConfigDatabase) DeleteService(ctx context.Context, name string) error {
	key := path.Join(db.pathPrefix, servicesKey, name)

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
