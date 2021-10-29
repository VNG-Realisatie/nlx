// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/common/diagnostics"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

type proxyMetadata struct {
	OrganizationName         string
	OrganizationSerialNumber string
	PublicKeyFingerprint     string
	PublicKeyPEM             string
}

func outgoingAccessRequestStateToProto(state database.OutgoingAccessRequestState) api.AccessRequestState {
	name := strings.ToUpper(string(state))

	protoState, ok := api.AccessRequestState_value[name]
	if ok {
		return api.AccessRequestState(protoState)
	}

	return api.AccessRequestState_UNSPECIFIED
}

func incomingAccessRequestStateToProto(state database.IncomingAccessRequestState) api.AccessRequestState {
	name := strings.ToUpper(string(state))

	protoState, ok := api.AccessRequestState_value[name]
	if ok {
		return api.AccessRequestState(protoState)
	}

	return api.AccessRequestState_UNSPECIFIED
}

func (s *ManagementService) ListIncomingAccessRequest(ctx context.Context, req *api.ListIncomingAccessRequestsRequests) (*api.ListIncomingAccessRequestsResponse, error) {
	_, err := s.configDatabase.GetService(ctx, req.ServiceName)
	if err != nil {
		if errIsNotFound(err) {
			return nil, status.Error(codes.NotFound, "service not found")
		}

		s.logger.Error("fetching service", zap.String("name", req.ServiceName), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	accessRequests, err := s.configDatabase.ListAllIncomingAccessRequests(ctx)
	if err != nil {
		s.logger.Error("fetching incoming access requests", zap.String("service name", req.ServiceName), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	filtered := []*api.IncomingAccessRequest{}

	for _, accessRequest := range accessRequests {
		if accessRequest.Service.Name == req.ServiceName {
			responseAccessRequest, err := convertIncomingAccessRequest(accessRequest)
			if err != nil {
				s.logger.Error(
					"converting incoming access request",
					zap.Uint("id", accessRequest.ID),
					zap.String("service", accessRequest.Service.Name),
					zap.Error(err),
				)

				return nil, status.Error(codes.Internal, "converting incoming access request")
			}

			filtered = append(filtered, responseAccessRequest)
		}
	}

	response := &api.ListIncomingAccessRequestsResponse{
		AccessRequests: filtered,
	}

	return response, nil
}

func (s *ManagementService) ApproveIncomingAccessRequest(ctx context.Context, req *api.ApproveIncomingAccessRequestRequest) (*emptypb.Empty, error) {
	accessRequest, err := s.getIncomingAccessRequest(ctx, req.AccessRequestID)
	if err != nil {
		return nil, err
	}

	if accessRequest.State == database.IncomingAccessRequestApproved {
		return nil, status.Error(codes.AlreadyExists, "access request is already approved")
	}

	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.IncomingAccessRequestAccept(ctx, userInfo.username, userInfo.userAgent, accessRequest.Organization.Name, req.ServiceName)
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

	return &emptypb.Empty{}, nil
}

func (s *ManagementService) RejectIncomingAccessRequest(ctx context.Context, req *api.RejectIncomingAccessRequestRequest) (*emptypb.Empty, error) {
	accessRequest, err := s.getIncomingAccessRequest(ctx, req.AccessRequestID)
	if err != nil {
		s.logger.Error(
			"getting incoming access request of request",
			zap.String("serviceName", req.ServiceName),
			zap.Uint("accessRequestID", uint(req.AccessRequestID)),
			zap.Error(err),
		)

		return nil, err
	}

	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.IncomingAccessRequestReject(ctx, userInfo.username, userInfo.userAgent, accessRequest.Organization.Name, req.ServiceName)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	err = s.configDatabase.UpdateIncomingAccessRequestState(ctx, accessRequest.ID, database.IncomingAccessRequestRejected)
	if err != nil {
		s.logger.Error("error updating incoming access request to rejected", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &emptypb.Empty{}, nil
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

func (s *ManagementService) CreateAccessRequest(ctx context.Context, req *api.CreateAccessRequestRequest) (*api.OutgoingAccessRequest, error) {
	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.OutgoingAccessRequestCreate(ctx, userInfo.username, userInfo.userAgent, req.OrganizationName, req.ServiceName)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	publicKey, err := s.orgCert.PublicKeyPEM()
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to parse public key")
	}

	ar := &database.OutgoingAccessRequest{
		OrganizationName:     req.OrganizationName,
		ServiceName:          req.ServiceName,
		PublicKeyPEM:         publicKey,
		PublicKeyFingerprint: s.orgCert.PublicKeyFingerprint(),
		State:                database.OutgoingAccessRequestCreated,
	}

	request, err := s.configDatabase.CreateOutgoingAccessRequest(ctx, ar)
	if err != nil {
		if errors.Is(err, database.ErrActiveAccessRequest) {
			return nil, status.Errorf(codes.AlreadyExists, "there is already an active access request")
		}

		return nil, err
	}

	response, err := convertOutgoingAccessRequest(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ManagementService) SendAccessRequest(ctx context.Context, req *api.SendAccessRequestRequest) (*api.OutgoingAccessRequest, error) {
	accessRequest, err := s.configDatabase.GetOutgoingAccessRequest(ctx, uint(req.AccessRequestID))
	if err != nil {
		if errIsNotFound(err) {
			return nil, status.Error(codes.NotFound, "access request not found")
		}

		s.logger.Error("fetching access request", zap.Uint("id", uint(req.AccessRequestID)), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	if !accessRequest.IsSendable() {
		return nil, status.Error(codes.AlreadyExists, "access request is not in a sendable state")
	}

	err = s.configDatabase.UpdateOutgoingAccessRequestState(ctx, accessRequest.ID, database.OutgoingAccessRequestCreated, 0, nil)
	if err != nil {
		s.logger.Error("access request cannot be updated", zap.Uint("id", accessRequest.ID), zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response, err := convertOutgoingAccessRequest(accessRequest)
	if err != nil {
		s.logger.Error(
			"converting outgoing access request",
			zap.Uint("id", accessRequest.ID),
			zap.String("organization", accessRequest.OrganizationName),
			zap.String("service", accessRequest.ServiceName),
			zap.Error(err),
		)

		return nil, status.Error(codes.Internal, "converting outgoing access request")
	}

	return response, nil
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

	request := &database.IncomingAccessRequest{
		ServiceID: service.ID,
		Organization: database.IncomingAccessRequestOrganization{
			Name:         md.OrganizationName,
			SerialNumber: md.OrganizationSerialNumber,
		},
		PublicKeyPEM:         md.PublicKeyPEM,
		PublicKeyFingerprint: md.PublicKeyFingerprint,
		State:                database.IncomingAccessRequestReceived,
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
		ReferenceId: uint64(createdRequest.ID),
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

	request, err := s.configDatabase.GetLatestIncomingAccessRequest(ctx, md.OrganizationName, req.ServiceName)
	if err != nil {
		s.logger.Error("failed to retrieve latest outgoing access request", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to retrieve access request")
	}

	return &external.GetAccessRequestStateResponse{
		State: incomingAccessRequestStateToProto(request.State),
	}, nil
}

// nolint:dupl // incoming access request looks like outgoing access request
func convertIncomingAccessRequest(accessRequest *database.IncomingAccessRequest) (*api.IncomingAccessRequest, error) {
	return &api.IncomingAccessRequest{
		Id:               uint64(accessRequest.ID),
		OrganizationName: accessRequest.Organization.Name,
		ServiceName:      accessRequest.Service.Name,
		State:            incomingAccessRequestStateToProto(accessRequest.State),
		CreatedAt:        timestamppb.New(accessRequest.CreatedAt),
		UpdatedAt:        timestamppb.New(accessRequest.UpdatedAt),
	}, nil
}

// nolint:dupl // outgoing access request looks like incoming access request
func convertOutgoingAccessRequest(request *database.OutgoingAccessRequest) (*api.OutgoingAccessRequest, error) {
	var details *api.ErrorDetails

	if request.ErrorCause != "" {
		code := api.ErrorCode_INTERNAL

		if request.ErrorCode == int(diagnostics.NoInwaySelectedError) {
			code = api.ErrorCode_NO_INWAY_SELECTED
		}

		details = &api.ErrorDetails{
			Code:       code,
			Cause:      request.ErrorCause,
			StackTrace: request.ErrorStackTrace,
		}
	}

	return &api.OutgoingAccessRequest{
		Id:               uint64(request.ID),
		OrganizationName: request.OrganizationName,
		ServiceName:      request.ServiceName,
		State:            outgoingAccessRequestStateToProto(request.State),
		ErrorDetails:     details,
		CreatedAt:        timestamppb.New(request.CreatedAt),
		UpdatedAt:        timestamppb.New(request.UpdatedAt),
	}, nil
}
