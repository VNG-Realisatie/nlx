// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

func (s *ManagementService) CreateOutgoingOrder(ctx context.Context, request *api.CreateOutgoingOrderRequest) (*emptypb.Empty, error) {
	err := s.authorize(ctx, permissions.CreateOutgoingOrder)
	if err != nil {
		return nil, err
	}

	s.logger.Info("rpc request CreateOutgoingOrder")

	order := &database.CreateOutgoingOrder{
		Reference:      request.Reference,
		Description:    request.Description,
		PublicKeyPEM:   request.PublicKeyPEM,
		Delegatee:      request.Delegatee,
		ValidFrom:      request.ValidFrom.AsTime(),
		ValidUntil:     request.ValidUntil.AsTime(),
		AccessProofIds: request.AccessProofIds,
	}

	if err := validateCreateOutgoingOrder(order); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid outgoing order: %s", err))
	}

	userInfo, userAgent, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	accessProofs, err := s.configDatabase.GetAccessProofs(ctx, order.AccessProofIds)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not retrieve access proofs")
	}

	// Check if there are no duplicate service entries in the access proofs
	services := make(map[string]bool)

	for _, a := range accessProofs {
		key := fmt.Sprintf("%s/%s", a.OutgoingAccessRequest.Organization.SerialNumber, a.OutgoingAccessRequest.ServiceName)

		_, exists := services[key]
		if exists {
			return nil, status.Error(codes.Internal, "cannot create order with duplicate services")
		}

		services[key] = true
	}

	err = s.auditLogger.OrderCreate(ctx, userInfo.Email, userAgent, order.Delegatee, accessProofsToAuditLogRecordServices(accessProofs))
	if err != nil {
		s.logger.Error("failed to write auditlog", zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to write to auditlog")
	}

	if err := s.configDatabase.CreateOutgoingOrder(ctx, order); err != nil {
		s.logger.Error("failed to create outgoing order", zap.Error(err))

		if err == database.ErrDuplicateOutgoingOrder {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("an order with reference %s for %s already exist", order.Reference, order.Delegatee))
		}

		return nil, status.Errorf(codes.Internal, "failed to create outgoing order")
	}

	return &emptypb.Empty{}, nil
}

func validateOrganizationSerialNumber(value interface{}) error {
	valueAsString, _ := value.(string)

	err := common_tls.ValidateSerialNumber(valueAsString)
	if err != nil {
		return fmt.Errorf("organization serial number must be in a valid format: %s", err)
	}

	return err
}

func validateCreateOutgoingOrder(order *database.CreateOutgoingOrder) error {
	const (
		minLength = 1
		maxLength = 100
	)

	return validation.ValidateStruct(
		order,
		validation.Field(&order.Reference, validation.Required, validation.Length(minLength, maxLength)),
		validation.Field(&order.Description, validation.Required, validation.Length(minLength, maxLength)),
		validation.Field(&order.ValidUntil, validation.Min(order.ValidFrom).Error("order can not expire before the start date")),
		validation.Field(&order.PublicKeyPEM, validation.By(validatePublicKey)),
		validation.Field(&order.Delegatee, validation.By(validateOrganizationSerialNumber)),
		validation.Field(&order.AccessProofIds, validation.Required, validation.Length(1, 0)),
	)
}

func validateUpdateOutgoingOrder(order *database.UpdateOutgoingOrder) error {
	const (
		minLength = 1
		maxLength = 100
	)

	return validation.ValidateStruct(
		order,
		validation.Field(&order.Reference, validation.Required, validation.Length(minLength, maxLength)),
		validation.Field(&order.Description, validation.Required, validation.Length(minLength, maxLength)),
		validation.Field(&order.ValidUntil, validation.Min(order.ValidFrom).Error("order can not expire before the start date")),
		validation.Field(&order.PublicKeyPEM, validation.By(validatePublicKey)),
		validation.Field(&order.AccessProofIds, validation.Required, validation.Length(minLength, 0)),
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

	_, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return ErrInvalidPublicKey
	}

	return nil
}
