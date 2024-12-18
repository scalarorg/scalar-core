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

	// Build the final message with colors for each part
	coloredMsg := fmt.Sprintf("%s%s%s", colorReset, msg, colorReset)

	for _, log := range logs {
		if log == nil {
			coloredMsg += fmt.Sprintf(" %s<nil>%s", colorYellow, colorReset)
			continue
		}

		// Get the value and type of the log
		v := reflect.ValueOf(log)
		t := reflect.TypeOf(log)

		if v.Kind() == reflect.Struct || v.Kind() == reflect.Pointer {
			// Add colored struct type and content
			coloredMsg += fmt.Sprintf(" %s%s%s{%+v}", colorMagenta, t.Name(), colorYellow, log)
		} else {
			// Add other types with the same color
			coloredMsg += fmt.Sprintf(" %s%v%s", colorMagenta, log, colorYellow)
		}
	}

	// Log the final formatted message
	event.Msg(coloredMsg)
}

func logLevelf(event *zerolog.Event, format string, logs ...interface{}) {
	logger := NewLogger()

	if event == nil {
		event = logger.Info()
	}

	event.Msg(fmt.Sprintf(format, logs...))
}

func Green(msg string, fields ...interface{}) {
	msg = colorGreen + msg + colorReset
	logLevel(logger.Info(), msg, fields...)
}

func Greenf(format string, args ...interface{}) {
	msg := colorGreen + fmt.Sprintf(format, args...) + colorReset
	logLevelf(logger.Info(), msg)
}

func Yellow(msg string, fields ...interface{}) {
	logLevel(logger.Warn(), colorYellow+msg+colorReset, fields...)
}

func Yellowf(format string, args ...interface{}) {
	msg := colorYellow + fmt.Sprintf(format, args...) + colorReset
	logLevelf(logger.Warn(), msg)
}

func Red(msg string, fields ...interface{}) {
	msg = colorRed + msg + colorReset
	logLevel(logger.Error(), msg, fields...)
}

func Redf(format string, args ...interface{}) {
	msg := colorRed + fmt.Sprintf(format, args...) + colorReset
	logLevelf(logger.Error(), msg)
}

func Blue(msg string, fields ...interface{}) {
	msg = colorBlue + msg + colorReset
	logLevel(logger.Debug(), msg, fields...)
}

func Bluef(format string, args ...interface{}) {
	msg := colorBlue + fmt.Sprintf(format, args...) + colorReset
	logLevelf(logger.Debug(), msg)
}

func Cyan(msg string, fields ...interface{}) {
	msg = colorGray + msg + colorReset
	logLevel(logger.Trace(), msg, fields...)
}

func Cyanf(format string, args ...interface{}) {
	msg := colorGray + fmt.Sprintf(format, args...) + colorReset
	logLevelf(logger.Trace(), msg)
}

func Magenta(msg string, fields ...interface{}) {
	msg = colorMagenta + msg + colorReset
	logLevel(logger.Info(), msg, fields...)
}

func Magentaf(format string, args ...interface{}) {
	msg := colorMagenta + fmt.Sprintf(format, args...) + colorReset
	logLevelf(logger.Info(), msg)
}
