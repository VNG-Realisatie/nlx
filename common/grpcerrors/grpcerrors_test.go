// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package grpcerrors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"

	"go.nlx.io/nlx/common/grpcerrors"
	"go.nlx.io/nlx/common/grpcerrors/errors"
)

func TestEqual(t *testing.T) {
	t.Parallel()

	type args struct {
		err  error
		code grpcerrors.Code
	}

	tests := map[string]struct {
		args args
		want bool
	}{
		"when_not_equal": {
			args: args{
				err:  grpcerrors.New("test", codes.Canceled, errors.ErrorReason_ERROR_REASON_UNSPECIFIED, "test message", nil),
				code: errors.ErrorReason_INVALID_REQUEST,
			},
			want: false,
		},
		"happy_flow_internal_error": {
			args: args{
				err:  grpcerrors.NewInternal("test", "internal error", nil),
				code: errors.ErrorReason_INTERNAL_SERVER_ERROR,
			},
			want: true,
		},
		"happy_flow_normal_error": {
			args: args{
				err:  grpcerrors.New("test", codes.Canceled, errors.ErrorReason_ERROR_REASON_UNSPECIFIED, "test message", nil),
				code: errors.ErrorReason_ERROR_REASON_UNSPECIFIED,
			},
			want: true,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := grpcerrors.Equal(tt.args.err, tt.args.code)
			assert.Equal(t, tt.want, got)
		})
	}
}
