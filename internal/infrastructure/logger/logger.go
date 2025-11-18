package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	Log *slog.Logger
}

func InitLogger(env string) *Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return &Logger{
		Log: log,
	}
}
