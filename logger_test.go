package logger

import (
	"log/slog"
)

func ExampleLogger() {
	handler, err := NewMoonbaseLogger("<project id>", "<token>", nil)
	if err != nil {
		panic(err)
	}
	logger := slog.New(handler)
	logger.Warn("test")
}
