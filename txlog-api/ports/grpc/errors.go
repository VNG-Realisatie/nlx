// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/txlog-api/app/errors"
)

func ResponseFromError(err error) error {
	appError, ok := err.(errors.Error)
	if !ok {
		return status.Error(codes.Internal, "internal")
	}

	switch appError.ErrorType() {
	case errors.ErrorTypeIncorrectInput:
		return status.Error(codes.InvalidArgument, appError.Error())
	default:
		return status.Error(codes.Internal, "internal")
	}
}
