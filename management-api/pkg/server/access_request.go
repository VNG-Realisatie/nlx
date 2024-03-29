// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/common/diagnostics"
	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

type proxyMetadata struct {
	OrganizationName         string
	OrganizationSerialNumber string
	PublicKeyFingerprint     string
	PublicKeyPEM             string
}

const accessRequestStatePrefix = "ACCESS_REQUEST_STATE_%s"

func incomingAccessRequestStateToProto(state database.IncomingAccessRequestState) external.AccessRequestState {
	name := fmt.Sprintf(accessRequestStatePrefix, strings.ToUpper(string(state)))

	protoState, ok := external.AccessRequestState_value[name]
	if !ok {
		return external.AccessRequestState_ACCESS_REQUEST_STATE_UNSPECIFIED
	}

	return external.AccessRequestState(protoState)
}

func outgoingAccessRequestStateToProto(state database.OutgoingAccessRequestState) external.AccessRequestState {
	name := fmt.Sprintf(accessRequestStatePrefix, strings.ToUpper(string(state)))

	protoState, ok := external.AccessRequestState_value[name]
	if !ok {
		return external.AccessRequestState_ACCESS_REQUEST_STATE_UNSPECIFIED
	}

	return external.AccessRequestState(protoState)
}

func (s *ManagementService) ListIncomingAccessRequests(ctx context.Context, req *api.ListIncomingAccessRequestsRequest) (*api.ListIncomingAccessRequestsResponse, error) {
	err := s.authorize(ctx, permissions.ReadIncomingAccessRequests)
	if err != nil {
		return nil, err
	}

	_, err = s.configDatabase.GetService(ctx, req.ServiceName)
	if err != nil {
		if errIsNotFound(err) {
			return nil, ErrServiceDoesNotExist
		}

		s.logger.Error("fetching service", zap.String("name", req.ServiceName), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	accessRequests, err := s.configDatabase.ListIncomingAccessRequests(ctx, req.ServiceName)
	if err != nil {
		s.logger.Error("fetching incoming access requests", zap.String("service name", req.ServiceName), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	convertedAccessRequests := make([]*api.IncomingAccessRequest, len(accessRequests))

	for i, accessRequest := range accessRequests {
		convertedAccessRequests[i] = convertIncomingAccessRequest(accessRequest)
	}

	return &api.ListIncomingAccessRequestsResponse{
		AccessRequests: convertedAccessRequests,
	}, nil
}

func (s *ManagementService) ApproveIncomingAccessRequest(ctx context.Context, req *api.ApproveIncomingAccessRequestRequest) (*api.ApproveIncomingAccessRequestResponse, error) {
	err := s.authorize(ctx, permissions.ApproveIncomingAccessRequest)
	if err != nil {
		return nil, err
	}

	accessRequest, err := s.getIncomingAccessRequest(ctx, req.AccessRequestId)
	if err != nil {
		return nil, err
	}

	if accessRequest.State == database.IncomingAccessRequestApproved {
		return nil, status.Error(codes.AlreadyExists, "access request is already approved")
	}

	userInfo, userAgent, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.IncomingAccessRequestAccept(ctx, userInfo.Email, userAgent, accessRequest.Organization.SerialNumber, accessRequest.Organization.Name, req.ServiceName)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	err = s.configDatabase.UpdateIncomingAccessRequestState(ctx, accessRequest.ID, database.IncomingAccessRequestApproved)
	if err != nil {
		s.logger.Error("error updating incoming access request to aproved", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if _, err := s.configDatabase.CreateAccessGrant(ctx, accessRequest); err != nil {
		s.logger.Error("creating access grant", zap.Error(err))

		return nil, status.Error(codes.Internal, "creating access grant")
	}

	return &api.ApproveIncomingAccessRequestResponse{}, nil
}

func (s *ManagementService) RejectIncomingAccessRequest(ctx context.Context, req *api.RejectIncomingAccessRequestRequest) (*api.RejectIncomingAccessRequestResponse, error) {
	err := s.authorize(ctx, permissions.RejectIncomingAccessRequest)
	if err != nil {
		return nil, err
	}

	accessRequest, err := s.getIncomingAccessRequest(ctx, req.AccessRequestId)
	if err != nil {
		s.logger.Error(
			"getting incoming access request of request",
			zap.String("serviceName", req.ServiceName),
			zap.Uint("accessRequestID", uint(req.AccessRequestId)),
			zap.Error(err),
		)

		return nil, err
	}

	userInfo, userAgent, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.IncomingAccessRequestReject(ctx, userInfo.Email, userAgent, accessRequest.Organization.SerialNumber, accessRequest.Organization.Name, req.ServiceName)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	err = s.configDatabase.UpdateIncomingAccessRequestState(ctx, accessRequest.ID, database.IncomingAccessRequestRejected)
	if err != nil {
		s.logger.Error("error updating incoming access request to rejected", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &api.RejectIncomingAccessRequestResponse{}, nil
}

func (s *ManagementService) getIncomingAccessRequest(ctx context.Context, accessRequestID uint64) (*database.IncomingAccessRequest, error) {
	accessRequest, err := s.configDatabase.GetIncomingAccessRequest(ctx, uint(accessRequestID))
	if err != nil {
		if errIsNotFound(err) {
			return nil, status.Error(codes.NotFound, "access request not found")
		}

		s.logger.Error(
			"error fetching access request",
			zap.Uint("id", uint(accessRequestID)),
			zap.Error(err),
		)

		return nil, status.Error(codes.Internal, "database error")
	}

	return accessRequest, nil
}

func (s *ManagementService) parseProxyMetadata(ctx context.Context) (*proxyMetadata, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "missing metadata from the management proxy")
	}

	organizationName := md.Get("nlx-organization-name")
	if len(organizationName) != 1 && organizationName[0] == "" {
		return nil, status.Error(codes.Internal, "invalid metadata: organization name missing")
	}

	organizationSerialNumber := md.Get("nlx-organization-serial-number")
	if len(organizationSerialNumber) != 1 && organizationSerialNumber[0] == "" {
		return nil, status.Error(codes.Internal, "invalid metadata: organization serial number missing")
	}

	publicKeyFingerprint := md.Get("nlx-public-key-fingerprint")
	if len(publicKeyFingerprint) != 1 && publicKeyFingerprint[0] == "" {
		return nil, status.Error(codes.Internal, "invalid metadata: public key fingerprint missing")
	}

	publicKeyString := md.Get("nlx-public-key-der")
	if len(publicKeyString) != 1 && publicKeyString[0] == "" {
		return nil, status.Error(codes.Internal, "invalid metadata: public key missing")
	}

	publicKeyDER, err := base64.StdEncoding.DecodeString(publicKeyString[0])
	if err != nil {
		return nil, status.Error(codes.Internal, "invalid metadata: invalid public key")
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDER,
	})
	if publicKeyPEM == nil {
		return nil, status.Error(codes.Internal, "invalid metadata: invalid public key")
	}

	return &proxyMetadata{
		OrganizationName:         organizationName[0],
		OrganizationSerialNumber: organizationSerialNumber[0],
		PublicKeyPEM:             string(publicKeyPEM),
		PublicKeyFingerprint:     publicKeyFingerprint[0],
	}, nil
}

func (s *ManagementService) RequestAccess(ctx context.Context, req *external.RequestAccessRequest) (*external.RequestAccessResponse, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		return nil, err
	}

	service, err := s.configDatabase.GetService(ctx, req.ServiceName)
	if err != nil {
		if errIsNotFound(err) {
			s.logger.Error("getting service by name failed. service does not exist", zap.String("name", req.ServiceName), zap.Error(err))
			return nil, ErrServiceDoesNotExist
		}

		s.logger.Error("getting service by name failed", zap.String("name", req.ServiceName), zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to retrieve service")
	}

	if req.PublicKeyPem == "" {
		s.logger.Error("request missing public key pem", zap.String("service-name", req.ServiceName), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "missing public key pem")
	}

	publicKeyFingerPrint, err := tls.PemPublicKeyFingerprint([]byte(req.PublicKeyPem))
	if err != nil {
		s.logger.Error("cannot parse public key fingerprint", zap.Error(err), zap.String("public-key-pem", req.PublicKeyPem))
		return nil, status.Error(codes.InvalidArgument, "invalid public key pem")
	}

	request := &database.IncomingAccessRequest{
		ServiceID: service.ID,
		Organization: database.IncomingAccessRequestOrganization{
			Name:         md.OrganizationName,
			SerialNumber: md.OrganizationSerialNumber,
		},
		PublicKeyFingerprint: publicKeyFingerPrint,
		PublicKeyPEM:         req.PublicKeyPem,
		State:                database.IncomingAccessRequestReceived,
	}

	existingIncomingAccessRequest, err := s.configDatabase.GetLatestIncomingAccessRequest(ctx, md.OrganizationSerialNumber, req.GetServiceName(), request.PublicKeyFingerprint)
	if err != nil {
		if !errIsNotFound(err) {
			s.logger.Error("getting latest incoming access request failed", zap.String("organization-serial-number", md.OrganizationSerialNumber), zap.String("service-name", req.ServiceName), zap.String("public-key-pem", req.PublicKeyPem), zap.Error(err))
			return nil, status.Error(codes.Internal, "failed to create access request")
		}
	}

	var accessGrant *database.AccessGrant

	if existingIncomingAccessRequest != nil && existingIncomingAccessRequest.State == database.IncomingAccessRequestApproved {
		accessGrant, err = s.configDatabase.GetLatestAccessGrantForService(ctx, md.OrganizationSerialNumber, req.ServiceName, request.PublicKeyFingerprint)
		if err != nil {
			s.logger.Error("cannot get latest access grant from database ", zap.String("organization-serial-number", md.OrganizationSerialNumber), zap.String("service-name", req.ServiceName), zap.String("public-key-pem", req.PublicKeyPem), zap.Error(err))
			return nil, status.Error(codes.Internal, "internal")
		}
	}

	if isIncomingAccessRequestStillActive(existingIncomingAccessRequest, accessGrant) {
		return &external.RequestAccessResponse{
			ReferenceId:        uint64(existingIncomingAccessRequest.ID),
			AccessRequestState: incomingAccessRequestStateToProto(existingIncomingAccessRequest.State),
		}, nil
	}

	createdRequest, err := s.configDatabase.CreateIncomingAccessRequest(ctx, request)
	if err != nil {
		if errors.Is(err, database.ErrActiveAccessRequest) {
			return nil, status.Error(codes.AlreadyExists, "an active access request already exists")
		}

		s.logger.Error("create access request failed", zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to create access request")
	}

	return &external.RequestAccessResponse{
		ReferenceId:        uint64(createdRequest.ID),
		AccessRequestState: incomingAccessRequestStateToProto(createdRequest.State),
	}, nil
}

func (s *ManagementService) GetAccessRequestState(ctx context.Context, req *external.GetAccessRequestStateRequest) (*external.GetAccessRequestStateResponse, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		return nil, err
	}

	_, err = s.configDatabase.GetService(ctx, req.ServiceName)
	if err != nil {
		s.logger.Error("failed to get service for access request state", zap.Error(err))

		if errIsNotFound(err) {
			return nil, ErrServiceDoesNotExist
		}

		return nil, status.Error(codes.Internal, "database error")
	}

	request, err := s.configDatabase.GetLatestIncomingAccessRequest(ctx, md.OrganizationSerialNumber, req.ServiceName, req.PublicKeyFingerprint)
	if err != nil {
		s.logger.Error("failed to retrieve latest outgoing access request", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to retrieve access request")
	}

	return &external.GetAccessRequestStateResponse{
		State: incomingAccessRequestStateToProto(request.State),
	}, nil
}

func isIncomingAccessRequestStillActive(incomingAccessRequest *database.IncomingAccessRequest, accessGrant *database.AccessGrant) bool {
	if incomingAccessRequest == nil {
		return false
	}

	if incomingAccessRequest.State == database.IncomingAccessRequestReceived {
		return true
	}

	if incomingAccessRequest.State == database.IncomingAccessRequestApproved {
		return !accessGrant.TerminatedAt.Valid && !accessGrant.RevokedAt.Valid
	}

	return false
}

// nolint:dupl // incoming access request looks like outgoing access request
func convertIncomingAccessRequest(accessRequest *database.IncomingAccessRequest) *api.IncomingAccessRequest {
	return &api.IncomingAccessRequest{
		Id: uint64(accessRequest.ID),
		Organization: &external.Organization{
			Name:         accessRequest.Organization.Name,
			SerialNumber: accessRequest.Organization.SerialNumber,
		},
		ServiceName:          accessRequest.Service.Name,
		State:                incomingAccessRequestStateToProto(accessRequest.State),
		PublicKeyFingerprint: accessRequest.PublicKeyFingerprint,
		CreatedAt:            timestamppb.New(accessRequest.CreatedAt),
		UpdatedAt:            timestamppb.New(accessRequest.UpdatedAt),
	}
}

// nolint:dupl // outgoing access request looks like incoming access request
func convertOutgoingAccessRequest(request *database.OutgoingAccessRequest) *api.OutgoingAccessRequest {
	var details *api.ErrorDetails

	if request.ErrorCause != "" {
		code := api.ErrorCode_ERROR_CODE_INTERNAL

		if request.ErrorCode == int(diagnostics.NoInwaySelectedError) {
			code = api.ErrorCode_ERROR_CODE_NO_INWAY_SELECTED
		}

		details = &api.ErrorDetails{
			Code:        code,
			Cause:       request.ErrorCause,
			StackTraces: request.ErrorStackTrace,
		}
	}

	return &api.OutgoingAccessRequest{
		Id: uint64(request.ID),
		Organization: &external.Organization{
			SerialNumber: request.Organization.SerialNumber,
			Name:         request.Organization.Name,
		},
		ServiceName:          request.ServiceName,
		State:                outgoingAccessRequestStateToProto(request.State),
		ErrorDetails:         details,
		CreatedAt:            timestamppb.New(request.CreatedAt),
		UpdatedAt:            timestamppb.New(request.UpdatedAt),
		PublicKeyFingerprint: request.PublicKeyFingerprint,
	}
}
