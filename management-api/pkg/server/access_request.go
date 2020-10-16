package server

import (
	context "context"
	"errors"
	"fmt"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

type proxyMetadata struct {
	OrganizationName     string
	PublicKeyFingerprint string
}

var accessRequestState = map[database.AccessRequestState]api.AccessRequestState{
	database.AccessRequestUnspecified: api.AccessRequestState_UNSPECIFIED,
	database.AccessRequestFailed:      api.AccessRequestState_FAILED,
	database.AccessRequestCreated:     api.AccessRequestState_CREATED,
	database.AccessRequestReceived:    api.AccessRequestState_RECEIVED,
	database.AccessRequestApproved:    api.AccessRequestState_APPROVED,
	database.AccessRequestRejected:    api.AccessRequestState_REJECTED,
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
		if accessRequest.ServiceName == req.ServiceName {
			responseAccessRequest, err := convertIncomingAccessRequest(accessRequest)
			if err != nil {
				s.logger.Error(
					"converting incoming access request",
					zap.String("id", accessRequest.ID),
					zap.String("service", accessRequest.ServiceName),
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

func (s *ManagementService) ApproveIncomingAccessRequest(ctx context.Context, req *api.ApproveIncomingAccessRequestRequest) (*types.Empty, error) {
	accessRequest, err := s.newIncomingAccessRequestFromRequest(ctx, req.ServiceName, req.AccessRequestID)
	if err != nil {
		return nil, err
	}

	if accessRequest.State == database.AccessRequestApproved {
		return nil, status.Error(codes.AlreadyExists, "access request is already approved")
	}

	if _, err := s.configDatabase.CreateAccessGrant(ctx, accessRequest); err != nil {
		if errors.Is(err, database.ErrAccessRequestModified) {
			s.logger.Warn("creating access grant", zap.Error(err))
			return nil, status.Error(codes.Aborted, "access request modified")
		}

		s.logger.Error("creating access grant", zap.Error(err))

		return nil, status.Error(codes.Internal, "creating access grant")
	}

	return &types.Empty{}, nil
}

func (s *ManagementService) RejectIncomingAccessRequest(ctx context.Context, req *api.RejectIncomingAccessRequestRequest) (*types.Empty, error) {
	accessRequest, err := s.newIncomingAccessRequestFromRequest(ctx, req.ServiceName, req.AccessRequestID)
	if err != nil {
		s.logger.Error("getting incoming access request of request", zap.String("serviceName", req.ServiceName), zap.String("accessRequestID", req.AccessRequestID), zap.Error(err))
		return nil, err
	}

	err = s.configDatabase.UpdateIncomingAccessRequestState(ctx, accessRequest, database.AccessRequestRejected)
	if err != nil {
		s.logger.Error("error updating incoming access request to rejected", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &types.Empty{}, nil
}

func (s *ManagementService) newIncomingAccessRequestFromRequest(ctx context.Context, serviceName, accessRequestID string) (*database.IncomingAccessRequest, error) {
	_, err := s.configDatabase.GetService(ctx, serviceName)
	if err != nil {
		if errIsNotFound(err) {
			return nil, status.Error(codes.NotFound, "service not found")
		}

		s.logger.Error("error fetching service", zap.String("serviceName", serviceName), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	accessRequest, err := s.configDatabase.GetIncomingAccessRequest(ctx, accessRequestID)
	if err != nil {
		if errIsNotFound(err) {
			return nil, status.Error(codes.NotFound, "access request not found")
		}

		s.logger.Error("error fetching access request", zap.String("id", accessRequestID), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	if accessRequest.ServiceName != serviceName {
		return nil, status.Error(codes.InvalidArgument, "service name does not match the one from access request")
	}

	return accessRequest, nil
}

func (s *ManagementService) ListOutgoingAccessRequests(ctx context.Context, req *api.ListOutgoingAccessRequestsRequest) (*api.ListOutgoingAccessRequestsResponse, error) {
	requests, err := s.configDatabase.ListOutgoingAccessRequests(ctx, req.OrganizationName, req.ServiceName)
	if err != nil {
		return nil, err
	}

	response := &api.ListOutgoingAccessRequestsResponse{}
	response.AccessRequests = make([]*api.OutgoingAccessRequest, len(requests))

	for i, request := range requests {
		ra, err := convertOutgoingAccessRequest(request)
		if err != nil {
			return nil, err
		}

		response.AccessRequests[i] = ra
	}

	return response, nil
}

func (s *ManagementService) CreateAccessRequest(ctx context.Context, req *api.CreateAccessRequestRequest) (*api.OutgoingAccessRequest, error) {
	ar := &database.OutgoingAccessRequest{
		AccessRequest: database.AccessRequest{
			OrganizationName:     req.OrganizationName,
			ServiceName:          req.ServiceName,
			PublicKeyFingerprint: s.orgCert.PublicKeyFingerprint(),
		},
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
	accessRequest, err := s.configDatabase.GetOutgoingAccessRequest(ctx, req.AccessRequestID)
	if err != nil {
		if errIsNotFound(err) {
			return nil, status.Error(codes.NotFound, "access request not found")
		}

		s.logger.Error("fetching access request", zap.String("id", req.AccessRequestID), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	if accessRequest.OrganizationName != req.OrganizationName {
		s.logger.Error("organization mismatch", zap.String("organization", req.OrganizationName))
		return nil, status.Error(codes.NotFound, "organization not found")
	}

	if accessRequest.ServiceName != req.ServiceName {
		s.logger.Error("service mismatch", zap.String("service", req.ServiceName))
		return nil, status.Error(codes.NotFound, "service not found")
	}

	if !accessRequest.Sendable() {
		return nil, status.Error(codes.AlreadyExists, "access request is not in a sendable state")
	}

	err = s.configDatabase.UpdateOutgoingAccessRequestState(ctx, accessRequest, database.AccessRequestCreated)
	if err != nil {
		s.logger.Error("access request cannot be updated", zap.String("id", accessRequest.ID), zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response, err := convertOutgoingAccessRequest(accessRequest)
	if err != nil {
		s.logger.Error(
			"converting outgoing access request",
			zap.String("id", accessRequest.ID),
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

	organizationName := md.Get("nlx-organization")
	if len(organizationName) != 1 {
		return nil, status.Error(codes.Internal, "invalid metadata: organization name missing")
	}

	publicKeyFingerprint := md.Get("nlx-public-key-fingerprint")
	if len(publicKeyFingerprint) != 1 {
		return nil, status.Error(codes.Internal, "invalid metadata: public key fingerprint missing")
	}

	return &proxyMetadata{
		OrganizationName:     organizationName[0],
		PublicKeyFingerprint: publicKeyFingerprint[0],
	}, nil
}

func (s *ManagementService) RequestAccess(ctx context.Context, req *external.RequestAccessRequest) (*types.Empty, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		return nil, err
	}

	request := &database.IncomingAccessRequest{
		AccessRequest: database.AccessRequest{
			ServiceName:          req.ServiceName,
			OrganizationName:     md.OrganizationName,
			PublicKeyFingerprint: md.PublicKeyFingerprint,
			State:                database.AccessRequestReceived,
		},
	}

	_, err = s.configDatabase.CreateIncomingAccessRequest(ctx, request)
	if err != nil {
		if errors.Is(err, database.ErrActiveAccessRequest) {
			return nil, status.Error(codes.AlreadyExists, "an active access request already exists")
		}

		s.logger.Error("create access request failed", zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to create access request")
	}

	return &types.Empty{}, nil
}

func (s *ManagementService) GetAccessRequestState(ctx context.Context, req *external.GetAccessRequestStateRequest) (*external.GetAccessRequestStateResponse, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		return nil, err
	}

	request, err := s.configDatabase.GetLatestIncomingAccessRequest(ctx, md.OrganizationName, req.ServiceName)
	if err != nil {
		s.logger.Error("failed to retrieve latest outgoing access request", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to retrieve access request")
	}

	return &external.GetAccessRequestStateResponse{
		State: api.AccessRequestState(request.State),
	}, nil
}

// nolint:dupl // incoming access request looks like outgoing access request
func convertIncomingAccessRequest(accessRequest *database.IncomingAccessRequest) (*api.IncomingAccessRequest, error) {
	createdAt, err := types.TimestampProto(accessRequest.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := types.TimestampProto(accessRequest.UpdatedAt)
	if err != nil {
		return nil, err
	}

	state, ok := accessRequestState[accessRequest.State]
	if !ok {
		return nil, fmt.Errorf("unsupported state: %v", accessRequest.State)
	}

	return &api.IncomingAccessRequest{
		Id:               accessRequest.ID,
		OrganizationName: accessRequest.OrganizationName,
		ServiceName:      accessRequest.ServiceName,
		State:            state,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}, nil
}

// nolint:dupl // outgoing access request looks like incoming access request
func convertOutgoingAccessRequest(a *database.OutgoingAccessRequest) (*api.OutgoingAccessRequest, error) {
	createdAt, err := types.TimestampProto(a.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := types.TimestampProto(a.UpdatedAt)
	if err != nil {
		return nil, err
	}

	aState, ok := accessRequestState[a.State]
	if !ok {
		return nil, fmt.Errorf("unsupported state: %v", a.State)
	}

	return &api.OutgoingAccessRequest{
		Id:               a.ID,
		OrganizationName: a.OrganizationName,
		ServiceName:      a.ServiceName,
		State:            aState,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}, nil
}
