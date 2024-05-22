# moonbase-logger

A [slog](https://pkg.go.dev/log/slog) handler that logs to [Moonbase](https://moonbasehq.com).

## Usage

```go
package main

import (
	"log/slog"

	"github.com/kevmo314/moonbase-logger"
)

func main() {
	handler, err := logger.NewMoonbaseLogger("<project id>", "<token>", nil)
	if err != nil {
		panic(err)
	}
	logger := slog.New(handler)
	logger.Warn("test")
}
```
