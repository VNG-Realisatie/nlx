// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
)

const PREFIX = "nlx"

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
