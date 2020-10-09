// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"path"
)

type Inway struct {
	Name        string   `json:"name,omitempty"`
	Version     string   `json:"version,omitempty"`
	Hostname    string   `json:"hostname,omitempty"`
	IPAddress   string   `json:"ipAddress,omitempty"`
	SelfAddress string   `json:"selfAddress,omitempty"`
	Services    []string `json:"services,omitempty"`
}

const inwaysKey = "inways"

// ListInways returns a list of inways
func (db ETCDConfigDatabase) ListInways(ctx context.Context) ([]*Inway, error) {
	inways := []*Inway{}

	if err := db.list(ctx, inwaysKey, &inways); err != nil {
		return nil, err
	}

	return inways, nil
}

// GetInway returns a specific inway by name
func (db ETCDConfigDatabase) GetInway(ctx context.Context, name string) (*Inway, error) {
	key := path.Join(inwaysKey, name)
	inway := &Inway{}

	if err := db.get(ctx, key, inway); err != nil {
		return nil, err
	}

	return inway, nil
}

// CreateInway creates a new inway
func (db ETCDConfigDatabase) CreateInway(ctx context.Context, inway *Inway) error {
	key := path.Join(inwaysKey, inway.Name)

	if err := db.put(ctx, key, inway); err != nil {
		return err
	}

	return nil
}

// UpdateInway updates an existing inway
func (db ETCDConfigDatabase) UpdateInway(ctx context.Context, name string, inway *Inway) error {
	if _, err := db.GetInway(ctx, name); err != nil {
		return err
	}

	key := path.Join(inwaysKey, name)

	if err := db.put(ctx, key, inway); err != nil {
		return err
	}

	return nil
}

// DeleteInway deletes a specific inway
func (db ETCDConfigDatabase) DeleteInway(ctx context.Context, name string) error {
	key := path.Join(db.pathPrefix, inwaysKey, name)

	_, err := db.etcdCli.Delete(ctx, key)
	if err != nil {
		return err
	}

	return nil
}
