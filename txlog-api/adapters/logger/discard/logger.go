package discard

import (
	"go.nlx.io/nlx/txlog-api/ports/logger"
)

type Logger struct{}

func (l *Logger) Error(string, error) {}

func (l *Logger) Info(string) {}

func (l *Logger) Warn(string, error) {}

func (l *Logger) Fatal(string, error) {}

func New() logger.Logger {
	return &Logger{}
}
