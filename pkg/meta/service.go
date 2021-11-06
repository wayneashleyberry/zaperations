package meta

import "go.uber.org/zap/zapcore"

// ServiceContext holds the name and version of our service. Using a service
// context enables rich error reporting as described in the documentation here:
// https://cloud.google.com/error-reporting/reference/rest/v1beta1/ServiceContext
type ServiceContext struct {
	Service string
	Version string
}

// MarshalLogObject satisfies the zapcore.ObjectMarshaler interface.
func (s ServiceContext) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("service", s.Service)
	enc.AddString("version", s.Version)

	return nil
}
