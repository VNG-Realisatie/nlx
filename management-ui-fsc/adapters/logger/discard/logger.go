// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package discardlogger

import "go.nlx.io/nlx/management-ui-fsc/adapters/logger"

type Logger struct{}

func (l *Logger) Error(string, error) {}

func (l *Logger) Info(string) {}

func (l *Logger) Warn(string, error) {}

func (l *Logger) Fatal(string, error) {}

func New() logger.Logger {
	return &Logger{}
}
