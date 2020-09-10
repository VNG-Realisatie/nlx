// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package grpcproxy

import (
	"errors"

	"google.golang.org/grpc/encoding"
)

var (
	_ encoding.Codec = (*passthroughCodec)(nil)
)

type message struct {
	b []byte
}

type passthroughCodec struct{}

func (c *passthroughCodec) Marshal(v interface{}) ([]byte, error) {
	m, ok := v.(*message)
	if !ok {
		return nil, errors.New("failed to marshal")
	}

	return m.b, nil
}

func (c *passthroughCodec) Unmarshal(data []byte, v interface{}) error {
	m, ok := v.(*message)
	if !ok {
		return errors.New("failed to unmarshal")
	}

	m.b = data

	return nil
}

func (c *passthroughCodec) Name() string {
	return "passthrough"
}

func (c *passthroughCodec) String() string {
	return c.Name()
}
