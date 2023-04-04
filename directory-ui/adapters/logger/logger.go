// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package logger

type Logger interface {
	Info(message string)
	Warn(message string, error error)
	Error(message string, error error)
	Fatal(message string, error error)
}
