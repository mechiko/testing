package config

// конфигурация на случай сбоя инициализации приложения
// или как пример конфига

var tomlConfig2 = []byte(`
# This is a TOML document.
name = "App merge declaration"
hostname = "localhost"
hostport = "3600"
debug = true
isadmin = false
browser = ""
timelayout = '2006-01-02T15:04:05-0700'

[gui]
limitforemptyfilter = 1000
hidemainwindow = false
output = "output"
export = "xlsx"
nouseperiodforfilter = false
tray = false

[logging]
logtruncate = true
logfilename = "alcogo.logs"
logfilenameecho = ""
nocolor = false
fileloggingenabled = false
consoleloggingenabled = false
directory = ""
maxsize = 10
maxbackups = 10
maxage = 10

[database]
timeout = 2
driver = "sqlite3"
connectionuri = "?cache=shared&mode=rw"
dbname = ""
`)

type DefaultConfiguration struct {
	Name     string                `json:"name"`
	Hostname string                `json:"hostname"`
	HostPort string                `json:"hostport"`
	Debug    bool                  `json:"debug"`
	IsAdmin  bool                  `json:"isadmin"`
	Browser  string                `json:"browser"`
	Database databaseConfiguration `json:"database"`
	Logging  loggingConfiguration  `json:"logging"`
	Gui      guiGonfiguration      `json:"gui"`
}

type loggingConfiguration struct {
	// Enable console logging
	LogFilename           string
	LogFilenameEcho       string
	LogTruncate           bool
	ConsoleLoggingEnabled bool
	NoColor               bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

type guiGonfiguration struct {
	Tray                 bool   `json:"tray"`
	LimitForEmptyFilter  int    `json:"limitforemptyfilter"`
	NoUsePeriodForFilter bool   `json:"nouseperiodforfilter"`
	Output               string `json:"output"`
	HideMainWindow       bool   `json:"hidemainwindow"`
	Export               string `json:"xlsx"`
}

type databaseConfiguration struct {
	ConnectionUri string `json:"connectionuri"`
	Driver        string `json:"driver"`
	DbName        string `json:"dbname"`
	Timeout       int    `json:"timeout"`
}
