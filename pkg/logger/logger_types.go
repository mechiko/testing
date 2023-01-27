package logger

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
	Interface
	logfile *os.File
}

// truncate output file names
// https://github.com/rs/zerolog
var tmy = mylog{}

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
		e.Str("caller:", `[`+strings.TrimPrefix(path.Dir(file), `E:/src/goproj/test/testing/`)+`][`+path.Base(file)+`]:`+strconv.Itoa(line))
	}
}
