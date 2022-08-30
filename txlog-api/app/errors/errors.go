// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package errors

type Type struct {
	t string
}

var (
	ErrorTypeIncorrectInput = Type{"incorrect-input"}
)

type Error struct {
	message   string
	errorType Type
}

func (s Error) Error() string {
	return s.message
}

func (s Error) ErrorType() Type {
	return s.errorType
}

func NewIncorrectInputError(message string) Error {
	return Error{
		message:   message,
		errorType: ErrorTypeIncorrectInput,
	}
}
