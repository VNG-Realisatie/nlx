package server

import (
	context "context"
	"errors"
	"fmt"

	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

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
