// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package zaplogger

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.nlx.io/nlx/directory-ui-htmx/adapters/logger"
)

type Logger struct {
	logger *zap.Logger
}

func (l *Logger) Error(message string, err error) {
	l.logger.Error(message, zap.Error(err))
}

func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

func (l *Logger) Warn(message string, err error) {
	l.logger.Warn(message, zap.Error(err))
}

func (l *Logger) Fatal(message string, err error) {
	l.logger.Fatal(message, zap.Error(err))
}

func New(logLevel, logType string) (logger.Logger, error) {
	var config zap.Config

	switch logType {
	case "live":
		config = zap.Config{
			Development: true,
			Encoding:    "json",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:     "message",
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "name",
				CallerKey:      "caller",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}
	case "local":
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		return nil, errors.New("invalid value " + logType + " for option log type")
	}

	switch logLevel {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "":
		// ignore, use default loglevel for provided type
	default:
		return nil, errors.New("invalid value " + logLevel + " for option log level")
	}

	lgr, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{
		logger: lgr,
	}, nil
}
