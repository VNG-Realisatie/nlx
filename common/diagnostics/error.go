package diagnostics

import (
	"runtime/debug"
	"strings"

	"google.golang.org/grpc/status"
)

type ErrorCode int

const (
	InternalError ErrorCode = iota
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
		Cause:      err.Error(),
		Code:       InternalError,
		StackTrace: strings.Split(string(debug.Stack()), "\n"),
	}
}

func (err *ErrorDetails) WithCode(code ErrorCode) *ErrorDetails {
	err.Code = code

	return err
}
