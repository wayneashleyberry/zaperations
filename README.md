> Package zaperations provides a [Google Cloud Operations](https://cloud.google.com/products/operations) (formerly Stackdriver) compatible config for the [uber-go/zap](https://github.com/uber-go/zap) logger.

[![Go](https://github.com/wayneashleyberry/zaperations/actions/workflows/go.yml/badge.svg)](https://github.com/wayneashleyberry/zaperations/actions/workflows/go.yml)

## Example

_This example is using unreleased functionality from go tip that is set to be released in 1.18._

```go
package main

import (
	"runtime/debug"

	"github.com/wayneashleyberry/zaperations/pkg/logger"
	"github.com/wayneashleyberry/zaperations/pkg/meta"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.NewProduction()
	if err != nil {
		panic(err)
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		panic("could not read build info")
	}

	var version string

	for _, setting := range info.Settings {
		if setting.Key == "gitrevision" {
			version = setting.Value
		}
	}

	log = log.With(zap.Object("serviceContext", meta.ServiceContext{
		Service: "logger-demo",
		Version: version,
	}))

	log.Info("Hello, World!")
}
```
