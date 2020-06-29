package configapi

import (
	context "context"

	"github.com/gogo/protobuf/types"

	"go.nlx.io/nlx/management-api/pkg/database"
)

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

	var status AccessRequest_Status

	switch a.Status {
	case database.AccessRequestFailed:
		status = AccessRequest_FAILED
	case database.AccessRequestCreated:
		status = AccessRequest_CREATED
	case database.AccessRequestSent:
		status = AccessRequest_SENT
	}

	return &AccessRequest{
		Id:               a.ID,
		OrganizationName: a.OrganizationName,
		ServiceName:      a.ServiceName,
		Status:           status,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}, nil
}
