package configapi

import (
	context "context"
	"errors"
	"fmt"

	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/pkg/database"
)

var accessRequestState = map[database.AccessRequestState]AccessRequest_State{
	database.AccessRequestFailed:   AccessRequest_FAILED,
	database.AccessRequestCreated:  AccessRequest_CREATED,
	database.AccessRequestReceived: AccessRequest_RECEIVED,
}

func (s *ConfigService) ListOutgoingAccessRequests(ctx context.Context, req *ListOutgoingAccessRequestsRequest) (*ListOutgoingAccessRequestsResponse, error) {
	l, err := s.configDatabase.ListOutgoingAccessRequests(ctx, req.OrganizationName, req.ServiceName)
	if err != nil {
		return nil, err
	}

	response := &ListOutgoingAccessRequestsResponse{}
	response.AccessRequests = make([]*AccessRequest, len(l))

	for i, a := range l {
		ra, err := convertAccessRequest(a)
		if err != nil {
			return nil, err
		}

		response.AccessRequests[i] = ra
	}

	return response, nil
}

func (s *ConfigService) CreateAccessRequest(ctx context.Context, req *CreateAccessRequestRequest) (*AccessRequest, error) {
	ar := &database.AccessRequest{
		OrganizationName: req.OrganizationName,
		ServiceName:      req.ServiceName,
	}

	a, err := s.configDatabase.CreateAccessRequest(ctx, ar)
	if err != nil {
		if errors.Is(err, database.ErrActiveAccessRequest) {
			return nil, status.Errorf(codes.AlreadyExists, "there is already an active access request")
		}

		return nil, err
	}

	response, err := convertAccessRequest(a)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func convertAccessRequest(a *database.AccessRequest) (*AccessRequest, error) {
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

	return &AccessRequest{
		Id:               a.ID,
		OrganizationName: a.OrganizationName,
		ServiceName:      a.ServiceName,
		State:            aState,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}, nil
}
