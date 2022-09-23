// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package diagnostics

import (
	"fmt"
	"runtime/debug"
	"strings"

	"google.golang.org/grpc/status"
)

type ErrorCode int

const (
	Unspecified ErrorCode = iota
	InternalError
	NoInwaySelectedError
)

type ErrorDetails struct {
	Code       ErrorCode
	Cause      string   `json:"cause"`
	StackTrace []string `json:"stackTrace"`
}

func ParseError(err error) *ErrorDetails {
	if err == nil {
		return nil
	}

	cause := err.Error()

	st, ok := status.FromError(err)
	if ok {
		cause = fmt.Sprintf("%s (%s)", st.Message(), st.Code().String())
	}

	return &ErrorDetails{
		Cause:      cause,
		Code:       InternalError,
		StackTrace: strings.Split(string(debug.Stack()), "\n"),
	}
}

func (err *ErrorDetails) WithCode(code ErrorCode) *ErrorDetails {
	err.Code = code

	return err
}
