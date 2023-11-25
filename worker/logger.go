package worker

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Print(level zerolog.Level, args ...interface{}) {
	log.WithLevel(level).Msg(fmt.Sprint(args...))
}

func (logger *Logger) Printf(ctx context.Context, format string, v ...interface{}) {
	log.WithLevel(zerolog.DebugLevel).Msgf(format, v...)
}

// Debug logs a message at Debug level.
func (l *Logger) Debug(args ...interface{}) {
	l.Print(zerolog.DebugLevel, args...)
}

// Info logs a message at Info level.
func (l *Logger) Info(args ...interface{}) {
	l.Print(zerolog.InfoLevel, args...)
}

// Warn logs a message at Warning level.
func (l *Logger) Warn(args ...interface{}) {
	l.Print(zerolog.WarnLevel, args...)
}

// Error logs a message at Error level.
func (l *Logger) Error(args ...interface{}) {
	l.Print(zerolog.ErrorLevel, args...)
}

// Fatal logs a message at Fatal level
// and process will exit with status set to 1.
func (l *Logger) Fatal(args ...interface{}) {
	l.Print(zerolog.FatalLevel, args...)
}
