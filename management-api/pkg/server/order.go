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
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s *ManagementService) CreateOrder(ctx context.Context, request *api.CreateOrderRequest) (*emptypb.Empty, error) {
	s.logger.Info("rpc request CreateOrder")

	order := convertOrder(request)

	if err := validateOrder(order); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid order: %s", err))
	}

	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	services := []auditlog.RecordService{}

	for _, service := range order.Services {
		services = append(services, auditlog.RecordService{
			Organization: service.Organization,
			Service:      service.Service,
		})
	}

	err = s.auditLogger.OrderCreate(ctx, userInfo.username, userInfo.userAgent, order.Delegatee, services)
	if err != nil {
		s.logger.Error("failed to write auditlog", zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to write to auditlog")
	}

	if err := s.configDatabase.CreateOrder(ctx, order); err != nil {
		s.logger.Error("failed to create order", zap.Error(err))

		return nil, status.Errorf(codes.Internal, "failed to create order")
	}

	return &emptypb.Empty{}, nil
}

func convertOrder(request *api.CreateOrderRequest) *database.Order {
	services := make([]database.OrderService, len(request.Services))

	for i, service := range request.Services {
		services[i] = database.OrderService{
			Organization: service.Organization,
			Service:      service.Service,
		}
	}

	return &database.Order{
		Reference:    request.Reference,
		Description:  request.Description,
		PublicKeyPEM: request.PublicKeyPEM,
		Delegatee:    request.Delegatee,
		ValidFrom:    request.ValidFrom.AsTime(),
		ValidUntil:   request.ValidUntil.AsTime(),
		Services:     services,
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
		validation.Field(&order.Services, validation.Required, validation.Each(validation.By(func(value interface{}) error {
			orderService, ok := value.(database.OrderService)
			if !ok {
				return errors.New("expecting an order-service")
			}

			return validation.ValidateStruct(
				&orderService,
				validation.Field(&orderService.Organization, validation.Match(organizationNameRegex).
					Error("organization must be in a valid format")),
				validation.Field(&orderService.Service, validation.Match(serviceNameRegex).
					Error("service must be in a valid format")),
			)
		}))),
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
