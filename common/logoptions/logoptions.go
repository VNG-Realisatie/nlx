// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package logoptions

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogOptions contains go-flags fields which can be used to configure a go-uber/zap config.
type LogOptions struct {
	LogType  string `long:"log-type" env:"LOG_TYPE" default:"production" description:"Set the logging config. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger." choice:"production" choice:"development"`
	LogLevel string `long:"log-level" env:"LOG_LEVEL" description:"Override the default loglevel as set by --log-type." choice:"debug" choice:"info" choice:"warn"`
}

// ZapConfig returns a ZapConfig based on provided LogOptions
func (o *LogOptions) ZapConfig() zap.Config {
	var config zap.Config
	switch o.LogType {
	case "production":
		config = zap.NewProductionConfig()
	case "development":
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		// This should never ocur since choices are limited by go-flags
		panic("invalid value " + o.LogType + " for option log type")
	}
	switch o.LogLevel {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "":
		// ignore, use default loglevel for provided type
	default:
		// This should never ocur since choices are limited by go-flags
		panic("invalid value " + o.LogLevel + " for option log level")
	}

	return config
}
