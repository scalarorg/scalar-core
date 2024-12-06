package clog

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

const (
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorReset   = "\033[0m"
	colorGray    = "\033[90m"
)

const (
	levelGreen  = "🌳"
	levelYellow = "🚌"
	levelRed    = "🔥"
	levelBlue   = "🌀"
	levelGray   = "🎲"
)

var (
	logger zerolog.Logger
	once   sync.Once
)

type ColorWriter struct {
	Writer io.Writer
}

func (w ColorWriter) Write(p []byte) (n int, err error) {
	return w.Writer.Write(p)
}

// NewLogger creates a new colored logger
func NewLogger() zerolog.Logger {
	once.Do(func() {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			FormatLevel: func(i interface{}) string {
				level := fmt.Sprintf("%s", i)
				switch level {
				case "info": // Changed from "info"
					return fmt.Sprintf("%s%s%s", colorGreen, levelGreen, colorReset)
				case "warn": // Changed from "warn"
					return fmt.Sprintf("%s%s%s", colorYellow, levelYellow, colorReset)
				case "error": // Changed from "error"
					return fmt.Sprintf("%s%s%s", colorRed, levelRed, colorReset)
				case "debug": // Changed from "debug"
					return fmt.Sprintf("%s%s%s", colorBlue, levelBlue, colorReset)
				case "trace": // Added cyan
					return fmt.Sprintf("%s%s%s", colorGray, levelGray, colorReset)
				default:
					return fmt.Sprintf("%s%s%s", colorMagenta, level, colorReset)
				}
			},
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("%s%v%s", colorCyan, i, colorReset)
			},
			FormatFieldName: func(i interface{}) string {
				return fmt.Sprintf("%s%s:%s", colorMagenta, i, colorReset)
			},
			FormatFieldValue: func(i interface{}) string {
				return fmt.Sprintf("%s%v%s", colorYellow, i, colorReset)
			},
		}

		logger = zerolog.New(output).With().Timestamp().Logger()
	})

	return logger
}

func logLevel(event *zerolog.Event, msg string, logs ...interface{}) {
	logger := NewLogger()

	if event == nil {
		event = logger.Info()
	}

	for _, log := range logs {
		typeOf := reflect.TypeOf(log)
		event.Interface(typeOf.String(), log)
	}

	event.Msg(msg)
}

func Green(msg string, fields ...interface{}) {
	logLevel(nil, msg, fields...)
}

func Yellow(msg string, fields ...interface{}) {
	logLevel(logger.Warn(), msg, fields...)
}

func Red(msg string, fields ...interface{}) {
	logLevel(logger.Error(), msg, fields...)
}

func Blue(msg string, fields ...interface{}) {
	logLevel(logger.Debug(), msg, fields...)
}

func Cyan(msg string, fields ...interface{}) {
	logLevel(logger.Trace(), msg, fields...)
}
