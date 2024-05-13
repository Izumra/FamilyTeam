package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return slog.New(logHandler)
}
