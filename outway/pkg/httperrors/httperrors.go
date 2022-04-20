// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package httperrors

import (
	"errors"
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func NewFromGRPCError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return errors.New("something went wrong")
	}

	var errString string

	for i, detail := range st.Details() {
		if i != len(st.Details())-1 {
			errString += "\n"
		}

		//nolint:gocritic // Type switches can't be if statements in Go
		switch t := detail.(type) {
		case *errdetails.BadRequest:
			for _, violation := range t.GetFieldViolations() {
				errString += fmt.Sprintf("%q: %s", violation.GetField(), violation.GetDescription())

				if i != len(t.GetFieldViolations())-1 {
					errString += "\n"
				}
			}
		}
	}

	return errors.New(err.Error() + errString)
}
