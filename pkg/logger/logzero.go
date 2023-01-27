package logger

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(cfg *ConfigZero) Logger {
	if cfg == nil {
		cfg = defaultConfig
	}
	// instance := InitLogger(cfg)
	if cfg.Truncate {
		os.Remove(cfg.Filename)
	}
	instance := cfg.configure()
	return instance
}

func (config *ConfigZero) configure() Logger {
	var writers []io.Writer
	var logger zerolog.Logger

	returnLogger := Logger{}

	// fmt.Println("LOGGER configure()")
	if config.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true})
	}
	if config.FileLoggingEnabled && (config.Filename != "") {
		if config.Rolling {
			f := config.newRollingFile()
			writers = append(writers, f)
		} else {
			if f, err := os.OpenFile(config.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
				fmt.Printf("error opening file logging: %s %v", config.Filename, err.Error())
			} else {
				returnLogger.logfile = f
				writers = append(writers, f)
			}
		}
	}
	mw := io.MultiWriter(writers...)
	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		// logger = zerolog.New(mw).With().Caller().Timestamp().Logger()
		logger = zerolog.New(mw).Hook(tmy).With().Timestamp().Logger()
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		logger = zerolog.New(mw).Hook(tmy)
	}
	returnLogger.Logger = logger
	return returnLogger
}

func (config *ConfigZero) newRollingFile() io.Writer {
	if config.Directory != "" {
		if err := os.MkdirAll(config.Directory, 0744); err != nil {
			fmt.Printf("can't create log directory %v\n\r", config.Directory)
			return nil
		}
	}
	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
	}
}
