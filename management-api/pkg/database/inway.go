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

type Inway struct {
	Name        string   `json:"name,omitempty"`
	Version     string   `json:"version,omitempty"`
	Hostname    string   `json:"hostname,omitempty"`
	IPAddress   string   `json:"ipAddress,omitempty"`
	SelfAddress string   `json:"selfAddress,omitempty"`
	Services    []string `json:"services,omitempty"`
}

// ListInways returns a list of inways
func (db ETCDConfigDatabase) ListInways(ctx context.Context) ([]*Inway, error) {
	key := path.Join(db.pathPrefix, "inways")
	if !strings.HasSuffix(key, "/") {
		key += "/"
	}

	getResponse, err := db.etcdCli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	inways := []*Inway{}

	for _, kv := range getResponse.Kvs {
		inway := &Inway{}
		err := json.Unmarshal(kv.Value, inway)

		if err != nil {
			return nil, err
		}

		inways = append(inways, inway)
	}

	return inways, nil
}

// GetInway returns a specific inway by name
func (db ETCDConfigDatabase) GetInway(ctx context.Context, name string) (*Inway, error) {
	key := path.Join(db.pathPrefix, "inways", name)

	values, err := db.etcdCli.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if values.Count == 0 {
		return nil, nil
	}

	inway := &Inway{}
	err = json.Unmarshal(values.Kvs[0].Value, inway)

	if err != nil {
		return nil, err
	}

	return inway, nil
}

// CreateInway creates a new inway
func (db ETCDConfigDatabase) CreateInway(ctx context.Context, inway *Inway) error {
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
//nolint:dupl // TODO
func (db ETCDConfigDatabase) UpdateInway(ctx context.Context, name string, inway *Inway) error {
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
