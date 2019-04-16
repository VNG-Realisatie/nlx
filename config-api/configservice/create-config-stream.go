package configservice

import (
	"context"
	"fmt"

	"go.nlx.io/nlx/config-api/configapi"
	"go.uber.org/zap"
)

type connection struct {
	stream configapi.ConfigApi_CreateConfigStreamServer
	id     string
	active bool
	error  chan error
}

// CreateConfigStream will setup a stream
func (s *ConfigService) CreateConfigStream(in *configapi.CreateConfigStreamRequest, stream configapi.ConfigApi_CreateConfigStreamServer) error {
	conn := &connection{
		stream: stream,
		id:     in.ComponentName,
		active: true,
		error:  make(chan error),
	}
	s.connections[conn.id] = conn
	s.logger.Info("rpc request CreateConfigStream")
	wChan := s.etcdCli.Watch(context.Background(), in.ComponentName)

	go func() {
		//TODO: graceful shutdown
		for {
			dbEvent := <-wChan
			s.logger.Info("received DB event")
			stream.Send(&configapi.Config{
				Config: string(dbEvent.Events[0].Kv.Value),
				Kind:   "inway",
			})
			if dbEvent.Canceled {
				s.logger.Info("DB watcher canceled returning")
				return
			}
		}
	}()
	value, err := s.etcdCli.Get(context.Background(), fmt.Sprintf("%s", in.ComponentName))
	if err != nil {
		s.logger.Error("cannot get config from db", zap.Error(err))
	}

	if value.Count > 0 {
		stream.Send(&configapi.Config{
			Kind:   "inway",
			Config: string(value.Kvs[0].Value),
		})
	} else {
		s.logger.Info("no config found")
	}

	return <-conn.error
}
