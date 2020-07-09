package configapi_test

import (
	"context"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/pkg/configapi"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
)

func newService(t *testing.T) (s *configapi.ConfigService, ctrl *gomock.Controller, db *mock_database.MockConfigDatabase) {
	logger := zaptest.Logger(t)
	proc := process.NewProcess(logger)

	ctrl = gomock.NewController(t)

	db = mock_database.NewMockConfigDatabase(ctrl)
	s = configapi.New(logger, proc, registrationapi.NewDirectoryRegistrationClient(nil), db)

	return
}

func TestCreateAccessRequest(t *testing.T) {
	service, ctrl, db := newService(t)
	ctrl.Finish()

	ctx := context.Background()

	createTimestamp := func(ti time.Time) *types.Timestamp {
		return &types.Timestamp{
			Seconds: ti.Unix(),
			Nanos:   int32(ti.Nanosecond()),
		}
	}

	tests := []struct {
		name        string
		req         *configapi.CreateAccessRequestRequest
		ar          *database.AccessRequest
		returnReq   *database.AccessRequest
		returnErr   error
		expectedReq *configapi.AccessRequest
		expectedErr error
	}{
		{
			"without active access request",
			&configapi.CreateAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			&database.AccessRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			&database.AccessRequest{
				ID:               "12345abcde",
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				Status:           database.AccessRequestCreated,
				CreatedAt:        time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
				UpdatedAt:        time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
			},
			nil,
			&configapi.AccessRequest{
				Id:               "12345abcde",
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				Status:           configapi.AccessRequest_CREATED,
				CreatedAt:        createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
				UpdatedAt:        createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
			},
			nil,
		},
		{
			"with active access request",
			&configapi.CreateAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			&database.AccessRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			nil,
			database.ErrActiveAccessRequest,
			nil,
			status.New(codes.AlreadyExists, "there is already an active access request").Err(),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db.EXPECT().CreateAccessRequest(ctx, tt.ar).Return(tt.returnReq, tt.returnErr)
			actual, err := service.CreateAccessRequest(ctx, tt.req)

			assert.Equal(t, tt.expectedReq, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
