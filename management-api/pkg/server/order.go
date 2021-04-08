// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s *ManagementService) CreateOrder(_ context.Context, request *api.CreateOrderRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request CreateOrder")

	order := convertOrder(request)

	if err := validateOrder(order); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid order: %s", err))
	}

	return &emptypb.Empty{}, nil
}

func convertOrder(request *api.CreateOrderRequest) *database.Order {
	return &database.Order{
		Reference:    request.Reference,
		Description:  request.Description,
		PublicKeyPEM: request.PublicKeyPEM,
		Delegatee:    request.Delegatee,
		ValidFrom:    request.ValidFrom.AsTime(),
		ValidUntil:   request.ValidUntil.AsTime(),
		Services:     request.Services,
	}
}

func validateOrder(order *database.Order) error {
	serviceNameRegex := regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`)
	organizationNameRegex := regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`)

	return validation.ValidateStruct(
		order,
		validation.Field(&order.Reference, validation.Required, validation.Length(1, 100)),
		validation.Field(&order.Description, validation.Required, validation.Length(1, 100)),
		validation.Field(&order.ValidUntil, validation.Min(order.ValidFrom).Error("order can not expire before the start date")),
		validation.Field(&order.PublicKeyPEM, validation.By(validatePublicKey)),
		validation.Field(&order.Services, validation.Required, validation.Each(validation.Match(serviceNameRegex).Error("service must be in a valid format"))),
		validation.Field(&order.Delegatee, validation.Match(organizationNameRegex)),
	)
}

func validatePublicKey(value interface{}) error {
	var (
		ErrInvalidPublicKeyFormat = errors.New("invalid public key format")
		ErrExpectPublicKeyAsPEM   = errors.New("expect public key as pem")
		ErrInvalidPublicKey       = errors.New("invalid public key")
	)

	publicKey, ok := value.(string)
	if !ok {
		return ErrInvalidPublicKeyFormat
	}

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return ErrExpectPublicKeyAsPEM
	}

	_, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return ErrInvalidPublicKey
	}

	return nil
}
