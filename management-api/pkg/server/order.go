// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
)

// CreateOrder creates a new order
func (s *ManagementService) CreateOrder(_ context.Context, request *api.CreateOrderRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request CreateOrder")

	err := validateInput(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid order: %s", err))
	}

	return &emptypb.Empty{}, nil
}

func validateInput(request *api.CreateOrderRequest) error {
	var (
		maxLengthReference   = 100
		minLengthReference   = 1
		maxLengthDescription = 100
		minLengthDescription = 1
	)

	if len(request.Reference) < minLengthReference {
		return errors.New("the reference must be provided")
	}

	if len(request.Reference) > maxLengthReference {
		return errors.New("the reference should not exceed 100 characters")
	}

	if len(request.Description) < minLengthDescription {
		return errors.New("the description must be provided")
	}

	if len(request.Description) > maxLengthDescription {
		return errors.New("the description should not exceed 100 characters")
	}

	// regex from https://gitlab.com/commonground/nlx/nlx/-/blob/4c7f0be8b2c9c980351b6202fbd2106bf4acdab0/directory-registration-api/pkg/database/inway.go#L16
	if !regexp.MustCompile(`^[a-zA-Z0-9-._\s]{1,100}$`).MatchString(request.Delegatee) {
		return errors.New("the delegatee should be a valid organization name (alphanumeric, max. 100 chars)")
	}

	if request.ValidUntil.AsTime().Before(request.ValidFrom.AsTime()) {
		return errors.New("valid from should be a timestamp before the valid until timestamp")
	}

	if len(request.Services) < 1 {
		return errors.New("at least one service should be specified")
	}

	// regex from https://gitlab.com/commonground/nlx/nlx/-/blob/4c7f0be8b2c9c980351b6202fbd2106bf4acdab0/directory-registration-api/pkg/database/inway.go#L19
	for _, serviceName := range request.Services {
		if !regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`).MatchString(serviceName) {
			return fmt.Errorf("service '%s' is not a valid service name", serviceName)
		}
	}

	return nil
}
