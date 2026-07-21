package logger

import (
	"os"
	"log/slog"
)

var globalLogger *slog.Logger

func Setup(isProd bool) {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: isProd,
		Level:     slog.LevelDebug,
	})
	globalLogger = slog.New(handler)
	slog.SetDefault(globalLogger)
}
