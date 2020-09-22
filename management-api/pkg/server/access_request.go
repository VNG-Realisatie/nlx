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
			OrganizationName: req.OrganizationName,
			ServiceName:      req.ServiceName,
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

	request, err := s.configDatabase.GetLatestOutgoingAccessRequest(ctx, md.OrganizationName, req.ServiceName)
	if err != nil {
		s.logger.Error("failed to retrieve latest outgoing access request", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to retrieve access request")
	}

	return &external.GetAccessRequestStateResponse{
		State: api.AccessRequestState(request.State),
	}, nil
}

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
