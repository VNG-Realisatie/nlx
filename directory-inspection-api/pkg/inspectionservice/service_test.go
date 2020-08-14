// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package inspectionservice_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-inspection-api/pkg/inspectionservice"
)

func generateMockDirectoryDatabase(t *testing.T) *mock.MockDirectoryDatabase {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return mock.NewMockDirectoryDatabase(mockCtrl)
}

func TestInspectionService_ListServices(t *testing.T) {
	type fields struct {
		database database.DirectoryDatabase
	}
	type args struct {
		ctx context.Context
		req *inspectionapi.ListServicesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *inspectionapi.ListServicesResponse
		wantErr bool
	}{
		{
			name: "happy flow",
			fields: fields{
				database: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					db.EXPECT().RegisterOutwayVersion(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					db.EXPECT().ListServices(gomock.Any(), "TODO").Return([]*database.Service{
						{
							Name: "Dummy Service Name",
						},
					}, nil).AnyTimes()

					return db
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: &inspectionapi.ListServicesRequest{},
			},
			want: &inspectionapi.ListServicesResponse{
				Services: []*inspectionapi.ListServicesResponse_Service{
					{
						ServiceName: "Dummy Service Name",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := zap.NewNop()
			service := inspectionservice.New(logger, tt.fields.database)
			got, err := service.ListServices(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListServices() got = %v, want %v", got, tt.want)
			}
		})
	}
}
