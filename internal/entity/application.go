package entity

import (
	"testing/pkg/logger"

	"github.com/rs/zerolog"
)

type App interface {
	GetLogger() logger.Logger
	GetConfig() ConfigInterface
	GetConfiguration() *Configuration
	GetDbService() DbService
	// GetDb2() sqlite3.DbService
	GetPwd() string
	GetOutput() string
	SetOutput(e string)
	SaveConfig() error
	GetExport() string
	SetExport(s string) error
	GetBrowser() string
	SetBrowser(s string) error
	GetBaseUrl() string
	DebugLog() *zerolog.Event
	ErrorLog() *zerolog.Event
	InfoLog() *zerolog.Event
	OpenDir()
	Open(url string)
	DumpSql(s string)
	DumpSqlAppend(s string)
	DumpSqlClear()
	SetRepo(Repo)
	GetRepo() Repo
	GetUtmInfo() (*UTMInfo, error)
	InitUtm() error
	InitDb() error
	SetMonitor(EventMonitor) error
	GetMonitor() EventMonitor
	InitLogger() error
	GetTaskScheduler() TaskList
	SetTaskScheduler(TaskList)
	GetHistory() History
	SetHistory(History)
	GetRecovery() RecoverInterface
	Shutdown()
	Restart()
	SetGuiService(GuiService)
	GetGuiService() GuiService
}

var AppInterrupt = make(chan int, 2)
