package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"slices"
	"strings"
)

var logger *slog.Logger

var (
	LogLevel            = ""
	LogOutput io.Writer = os.Stdout
)

func init() {
	Initialize()
}

func Initialize() {
	logger = slog.New(createLogHandler())
}

func createLogHandler() slog.Handler {
	LogLevel = os.Getenv("LOG_LEVEL")
	if !slices.Contains([]string{"debug", "info", "warn", "error"}, LogLevel) {
		LogLevel = "info"
	}

	opts := &slog.HandlerOptions{Level: parseLevel(LogLevel)}
	return slog.NewTextHandler(LogOutput, opts)
}

func parseLevel(lvl string) slog.Level {
	switch strings.ToLower(lvl) {
	case "error":
		return slog.LevelError
	case "warn":
		return slog.LevelWarn
	case "debug":
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}

func Info(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	logger.Info(msg)
}

func Fatal(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	logger.Error(msg)
	panic(msg)
}

func Error(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	logger.Error(msg)
}

func Warn(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	logger.Warn(msg)
}

func Debug(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	logger.Debug(msg)
}
