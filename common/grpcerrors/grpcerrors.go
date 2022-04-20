// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package grpcerrors

import (
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/runtime/protoiface"

	"go.nlx.io/nlx/common/grpcerrors/errors"
)

type Code interface {
	String() string
}

type Metadata struct {
	Metadata            map[string]string
	RetryInfo           *errdetails.RetryInfo
	DebugInfo           *errdetails.DebugInfo
	QuotaFailure        *errdetails.QuotaFailure
	PreconditionFailure *errdetails.PreconditionFailure
	BadRequest          *errdetails.BadRequest
	RequestInfo         *errdetails.RequestInfo
	ResourceInfo        *errdetails.ResourceInfo
	Help                *errdetails.Help
	LocalizedMessage    *errdetails.LocalizedMessage
}

// Don't use this function directly, wrap this function in a new function supplying the domain for that component.
// E.g. grpcerrors.New("inway", grpcCode, code, msg, md)
func New(domain string, grpcCode codes.Code, code Code, msg string, md *Metadata) error {
	if md == nil {
		md = &Metadata{}
	}

	details := []protoiface.MessageV1{&errdetails.ErrorInfo{
		Domain:   domain,
		Reason:   code.String(),
		Metadata: md.Metadata,
	}}

	details = appendMetadata(details, md)

	st, err := status.New(grpcCode, msg).
		WithDetails(details...)

	if err != nil {
		log.Println(err)
		return status.Errorf(codes.Internal, "error attaching metadata to error")
	}

	return st.Err()
}

func NewFromValidationError(domain string, err error) error {
	if e, ok := err.(validation.InternalError); ok {
		return New(domain, codes.Internal, errors.ErrorReason_VALIDATION_ERROR, e.InternalError().Error(), nil)
	}

	if e, ok := err.(validation.Errors); ok {
		fieldViolations := []*errdetails.BadRequest_FieldViolation{}

		for field, err := range e {
			fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{
				Field:       field,
				Description: err.Error(),
			})
		}

		return New(domain, codes.InvalidArgument, errors.ErrorReason_INVALID_REQUEST, "request has invalid fields", &Metadata{
			BadRequest: &errdetails.BadRequest{
				FieldViolations: fieldViolations,
			},
		})
	}

	return New(domain, codes.Internal, errors.ErrorReason_VALIDATION_ERROR, "invalid validation error type", nil)
}

func NewInternal(domain, msg string, md *Metadata) error {
	return New(domain, codes.Internal, errors.ErrorReason_INTERNAL_SERVER_ERROR, msg, md)
}

func Equal(err error, code Code) bool {
	st, ok := status.FromError(err)
	if !ok {
		return false
	}

	for _, detail := range st.Details() {
		//nolint:gocritic // Type switches can't be if statements in Go
		switch t := detail.(type) {
		case *errdetails.ErrorInfo:
			return t.GetReason() == code.String()
		}
	}

	return false
}

func appendMetadata(details []protoiface.MessageV1, md *Metadata) []protoiface.MessageV1 {
	if md.RetryInfo != nil {
		details = append(details, md.RetryInfo)
	}

	if md.DebugInfo != nil {
		details = append(details, md.DebugInfo)
	}

	if md.QuotaFailure != nil {
		details = append(details, md.QuotaFailure)
	}

	if md.PreconditionFailure != nil {
		details = append(details, md.PreconditionFailure)
	}

	if md.BadRequest != nil {
		details = append(details, md.BadRequest)
	}

	if md.RequestInfo != nil {
		details = append(details, md.RequestInfo)
	}

	if md.ResourceInfo != nil {
		details = append(details, md.ResourceInfo)
	}

	if md.Help != nil {
		details = append(details, md.Help)
	}

	if md.LocalizedMessage != nil {
		details = append(details, md.LocalizedMessage)
	}

	return details
}
