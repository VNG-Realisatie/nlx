// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package grpcerrors

import (
	"google.golang.org/grpc/codes"

	"go.nlx.io/nlx/common/grpcerrors"
)

const domain = "management"

func New(grpcCode codes.Code, code grpcerrors.Code, msg string, md *grpcerrors.Metadata) error {
	return grpcerrors.New(domain, grpcCode, code, msg, md)
}

func NewFromValidationError(err error) error {
	return grpcerrors.NewFromValidationError(domain, err)
}

func NewInternal(msg string, md *grpcerrors.Metadata) error {
	return grpcerrors.NewInternal(domain, msg, md)
}
