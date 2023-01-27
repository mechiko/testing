package entity

import (
	"testing/pkg/config"
	"testing/pkg/logger"

	"github.com/rs/zerolog"
)

type App interface {
	GetLogger() *logger.Logger
	GetConfig() *config.Config
	GetConfiguration() *Configuration
	GetDbService() DbService
	// GetDb2() sqlite3.DbService
	GetPwd() string
	GetOutput() string
	GetExport() string
	GetBrowser() string
	GetBaseUrl() string
	DebugLog() *zerolog.Event
	ErrorLog() *zerolog.Event
	InfoLog() *zerolog.Event
	OpenDir()
	Open(url string)
	DumpSql(s string)
	DumpSqlAppend(s string)
	DumpSqlClear()
	GetRepo() Repo
	GetUtmInfo() (*UTMInfo, error)
	InitUtm() error
	GetMonitor() EventMonitor
	GetTaskScheduler() TaskList
	GetHistory() History
	GetRecovery() RecoverInterface
	// SetGuiService(GuiService)
	// GetGuiService() GuiService
	Shutdown()
	Restart()
}

var AppInterrupt = make(chan int, 2)
