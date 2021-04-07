// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
)

// CreateOrder creates a new order
func (s *ManagementService) CreateOrder(_ context.Context, _ *api.CreateOrderRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request CreateOrder")
	return &emptypb.Empty{}, nil
}
