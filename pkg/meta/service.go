package meta

import "go.uber.org/zap/zapcore"

type ServiceContext struct {
	Service string
	Version string
}

func (s ServiceContext) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("service", s.Service)
	enc.AddString("version", s.Version)

	return nil
}
