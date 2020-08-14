// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package inspectionservice_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database/mock"
	"go.nlx.io/nlx/directory-inspection-api/pkg/inspectionservice"
)

func TestInspectionService_ListOrganizations(t *testing.T) {
	type fields struct {
		logger   *zap.Logger
		database database.DirectoryDatabase
	}
	type args struct {
		ctx context.Context
		req *inspectionapi.ListOrganizationsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *inspectionapi.ListOrganizationsResponse
		wantErr bool
	}{
		{
			name: "failed to get organizations from the database",
			fields: fields{
				logger: zap.NewNop(),
				database: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					db.EXPECT().ListOrganizations(gomock.Any()).Return(nil, errors.New("arbitrary error")).AnyTimes()

					return db
				}(),
			},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "happy flow",
			fields: fields{
				logger: zap.NewNop(),
				database: func() *mock.MockDirectoryDatabase {
					db := generateMockDirectoryDatabase(t)
					db.EXPECT().ListOrganizations(gomock.Any()).Return([]*database.Organization{
						{
							Name: "Dummy Organization Name",
						},
					}, nil).AnyTimes()

					return db
				}(),
			},
			args: args{},
			want: &inspectionapi.ListOrganizationsResponse{
				Organizations: []*inspectionapi.ListOrganizationsResponse_Organization{
					{
						Name: "Dummy Organization Name",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := inspectionservice.New(tt.fields.logger, tt.fields.database)
			got, err := h.ListOrganizations(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListOrganizations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListOrganizations() got = %v, want %v", got, tt.want)
			}
		})
	}
}
