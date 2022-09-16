// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package httperrors_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"

	"go.nlx.io/nlx/common/grpcerrors"
	grpc_common_errors "go.nlx.io/nlx/common/grpcerrors/errors"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/outway/pkg/httperrors"
)

func TestEqual(t *testing.T) {
	t.Parallel()

	type args struct {
		err error
	}

	requestClaim := &external.RequestClaimRequest{OrderReference: "test", ServiceOrganizationSerialNumber: "", ServiceName: ""}
	validationError := requestClaim.Validate()

	tests := map[string]struct {
		args args
		want string
	}{
		"when_not_a_grpc_error": {
			args: args{
				err: errors.New("arbitrary error"),
			},
			want: "something went wrong",
		},
		"happy_flow_validation_error": {
			args: args{
				err: grpcerrors.NewFromValidationError("test", validationError),
			},
			want: `rpc error: code = InvalidArgument desc = request has invalid fields
"service_name": cannot be blank"service_organization_serial_number": cannot be blank`,
		},
		"happy_flow_internal_error": {
			args: args{
				err: grpcerrors.NewInternal("test", "internal error", nil),
			},
			want: "rpc error: code = Internal desc = internal error",
		},
		"happy_flow": {
			args: args{
				err: grpcerrors.New("test", codes.Canceled, grpc_common_errors.ErrorReason_ERROR_REASON_UNSPECIFIED, "test message", nil),
			},
			want: "rpc error: code = Canceled desc = test message",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := httperrors.NewFromGRPCError(tt.args.err)

			assert.Error(t, got)
			assert.Equal(t, tt.want, got.Error())
		})
	}
}
