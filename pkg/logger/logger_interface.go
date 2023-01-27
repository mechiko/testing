package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// Interface -.
type Interface interface {
	IDebug(message interface{}, args ...interface{})
	IInfo(message string, args ...interface{})
	IWarn(message string, args ...interface{})
	IError(message interface{}, args ...interface{})
	IFatal(message interface{}, args ...interface{})
	Dispose()
}

// Debug -.
func (l *Logger) IDebug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

// Info -.
func (l *Logger) IInfo(message string, args ...interface{}) {
	l.log(message, args...)
}

// Warn -.
func (l *Logger) IWarn(message string, args ...interface{}) {
	l.log(message, args...)
}

// Error -.
func (l *Logger) IError(message interface{}, args ...interface{}) {
	if l.Logger.GetLevel() == zerolog.DebugLevel {
		l.IDebug(message, args...)
	}

	l.msg("error", message, args...)
}

// Fatal -.
func (l *Logger) LFatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)

	os.Exit(1)
}

func (l *Logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.Logger.Info().Msg(message)
	} else {
		l.Logger.Info().Msgf(message, args...)
	}
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}

// пытаемся закрыть файл лога
func (l *Logger) Dispose() {
	if l.logfile == nil {
		return
	}
	if err := l.logfile.Close(); err != nil {
		fmt.Printf("ошибка закрытия файла лога %s\n", err.Error())
	}
}
