package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ProductionConfig creates a zap config for use in production environments.
func ProductionConfig() zap.Config {
	return Config(false)
}

// DevelopmentConfig creates a zap config for use in development environments.
func DevelopmentConfig() zap.Config {
	return Config(true)
}

// Config creates a new zap config.
func Config(dev bool) zap.Config {
	config := zap.NewProductionConfig()

	if dev {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		config.Development = true
	}

	config.EncoderConfig.LevelKey = "severity"
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	config.DisableStacktrace = true

	return config
}
