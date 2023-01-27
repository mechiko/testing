package entity

var TomlConfig = []byte(`
# This is a TOML document.
name = "Request UTM"
hostname = "localhost"
hostport = "3600"
debug = false
isadmin = false
browser = ""
utmhost = "localhost"
utmport = "8080"
timelayout = "2006-01-02T15:04:05-0700"

[application]
scantimer = 30
utmstatetimer = 10
console = false
disconnected = true
fsrarid = ""
scanutm = false
requestsendtimer = 60
RequestSendCount = 20
tasktimer = 10

[gui]
limitforemptyfilter = 1000
hidemainwindow = false
output = "output"
export = "xlsx"
nouseperiodforfilter = false
tray = false

[logging]
logtruncate = true
logfilename = "logs.log"
logfilenameecho = ""
nocolor = false
fileloggingenabled = true
directory = ""
maxsize = 10
maxbackups = 10
maxage = 10

[database]
timeout = 2
driver = "sqlite3"
connectionuri = "?cache=shared&mode=rw"
dbname = "requests.db"
`)

type Configuration struct {
	Name        string           `json:"name"`
	Hostname    string           `json:"hostname"`
	HostPort    string           `json:"hostport"`
	UtmHost     string           `json:"utmhost"`
	UtmPort     string           `json:"utmport"`
	Debug       bool             `json:"debug"`
	IsAdmin     bool             `json:"isadmin"`
	Browser     string           `json:"browser"`
	TimeLayout  string           `json:"timelayout"`
	Application appConfiguration `json:"application"`
	// Alcohelp2 bool                  `json:"alcohelp2"`
	// Database   DatabaseConfiguration   `json:"database"`
	Database databaseConfiguration `json:"database"`
	Logging  loggingConfiguration  `json:"logging"`
	// Web        WebConfiguration        `json:"web"`
	// Version    Version                 `json:"version"`
	Gui guiGonfiguration `json:"gui"`
	// Declaracia DeclaraciaConfiguration `json:"declaracia"`
}

type loggingConfiguration struct {
	// Enable console logging
	LogFilename     string
	LogFilenameEcho string
	LogTruncate     bool
	// ConsoleLoggingEnabled bool
	NoColor bool
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

type appConfiguration struct {
	Console          bool   `json:"console"`
	Disconnected     bool   `json:"disconnected"`
	Fsrarid          string `json:"fsrarid"`
	ScanTimer        int    `json:"scantimer"`
	UtmStateTimer    int    `json:"utmstatetimer"`
	TaskTimer        int    `json:"tasktimer"`
	ScanUtm          bool   `json:"scanutm"`
	RequestSendTimer int    `json:"requestsendtimer"`
	RequestSendCount int    `json:"requestsendcount"`
	HistoryEventSize int    `json:"historyeventsize"`
}
