package logger

import (
	"runtime/debug"

	"github.com/wayneashleyberry/zaperations/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewDevelopment creates a new logger for use in development environments.
func NewDevelopment() (*zap.Logger, error) {
	config := config.Config(true)

	return config.Build(wrap())
}

// NewDevelopment creates a new logger for use in production environments.
func NewProduction() (*zap.Logger, error) {
	config := config.Config(false)

	return config.Build(wrap())
}

func wrap() zap.Option {
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return &core{c}
	})
}

type core struct {
	zapcore.Core
}

func (c *core) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}

	return ce
}

func (c *core) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	if ent.Level >= zapcore.ErrorLevel {
		ent.Message = ent.Message + "\n" + string(debug.Stack())
	}

	return c.Core.Write(ent, fields)
}

func (c *core) With(fields []zap.Field) zapcore.Core {
	return &core{c.Core.With(fields)}
}
