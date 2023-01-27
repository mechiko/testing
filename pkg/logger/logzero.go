package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zerolog.Logger
}

var tmy = mylog{}

func GetLogger() Interface {
	if instance == nil {
		fmt.Println("Logger not configured! Packed: testinglogger")
		fmt.Println("Logger will be configured default config! Packed: testinglogger")
		instance = InitLogger(defaultConfig)
	}
	return instance
}

func ZeroLog() *zerolog.Logger {
	if instance == nil {
		fmt.Println("Logger not configured! Packed: testinglogger")
		fmt.Println("Logger will be configured default config! Packed: testinglogger")
		instance = InitLogger(defaultConfig)
	}
	return instance.Logger
}

var defaultConfig = &ConfigZero{
	ConsoleLoggingEnabled: false,
	NoColor:               true,
	Debug:                 false,
	FileLoggingEnabled:    true,
	Rolling:               false,
	Directory:             "",
	// Filename:              cfg.Logging.LogFilename,
	Filename:   "log.txt",
	MaxSize:    10, // megabytes
	MaxBackups: 10, // files
	MaxAge:     5,  //days
	Truncate:   true,
}

// Configuration for logging
type ConfigZero struct {
	// Enable console logging
	ConsoleLoggingEnabled bool
	NoColor               bool
	Debug                 bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// lumberjack
	Rolling bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge   int
	Truncate bool
}

type mylog struct{}

func (h mylog) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if _, file, line, ok := runtime.Caller(3); ok {
		e.Str("caller:", `[`+strings.TrimPrefix(path.Dir(file), `E:/src/goproj/go_clean_architech/`)+`][`+path.Base(file)+`]:`+strconv.Itoa(line))
	}
}

func (config *ConfigZero) configure() *Logger {
	var writers []io.Writer
	var logger zerolog.Logger

	// fmt.Println("LOGGER configure()")
	if config.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true})
	}
	if config.FileLoggingEnabled && (config.Filename != "") {
		f, err := os.OpenFile(config.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("error opening file logging: %s %v", config.Filename, err.Error())
		} else {
			writers = append(writers, f)
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
	logger.Info().Msg("Logger Started")
	return &Logger{
		Logger: &logger,
	}
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

// func recoverFunc(str ...string) {
// 	if r := recover(); r != nil {
// 		err := fmt.Sprintf("%s %v", str, r)
// 		messageErr(err)
// 		d := []byte(err)
// 		_ = ioutil.WriteFile("error.txt", d, 0644)
// 		os.Exit(1)
// 	}
// }

// func messageErr(str string) {
// 	fmt.Println(str)
// }
