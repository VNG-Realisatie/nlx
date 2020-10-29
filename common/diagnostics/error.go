package diagnostics

import (
	"runtime/debug"
	"strings"
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

	code := InternalError
	stackTrace := strings.Split(string(debug.Stack()), "\n")

	return &ErrorDetails{
		Cause:      err.Error(),
		Code:       code,
		StackTrace: stackTrace,
	}
}

func (err *ErrorDetails) WithCode(code ErrorCode) *ErrorDetails {
	err.Code = code

	return err
}
