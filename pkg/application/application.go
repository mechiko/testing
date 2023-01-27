package application

import (
	"fmt"
	"os"
	"sync"

	"testing/internal/entity"
	"testing/pkg/config"
	"testing/pkg/logger"
	"testing/pkg/sqlite3"

	"github.com/rs/zerolog"
)

type Application interface {
	GetLogger() *logger.Logger
	GetConfig() *config.Config
	GetConfiguration() *entity.Configuration
	GetDbService() entity.DbService
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
	SetRepo(repo entity.Repo)
	GetRepo() entity.Repo
	GetUtmInfo() (*entity.UTMInfo, error)
	InitUtm() error
	InitDb() error
	SetMonitor(entity.EventMonitor) error
	GetMonitor() entity.EventMonitor
	InitLogger() error
	GetTaskScheduler() entity.TaskList
	SetTaskScheduler(entity.TaskList)
	GetHistory() entity.History
	SetHistory(entity.History)
	GetRecovery() entity.RecoverInterface
	Shutdown()
	Restart()
}

type applicationType struct {
	Name          string
	Config        *config.Config
	Logger        *logger.Logger
	Configuration *entity.Configuration
	Db            entity.DbService
	Repo          entity.Repo
	History       entity.History
	Recovery      entity.RecoverInterface
	// Settings
	pwd     string
	output  string
	export  string
	browser string
	monitor entity.EventMonitor
	tasks   entity.TaskList
}

var applicationInstance *applicationType = nil

var _ Application = &applicationType{}

var doOnce sync.Once

func init() {
	defer RecoverFmtFunc("pkg:application.init()")
	// инициализируем приложение тут
	if err := InstanceApp(); err != nil {
		fmt.Printf("Ошибка app.init() инициализации приложения %s\n", err.Error())
		os.Exit(1)
	}
	// к этому времени логгер уже должен работать
	Get().InfoLog().Msg("app.init() ok")
}

// проверка создан объект или нет
func IsInit() bool {
	return applicationInstance != nil
}

func InstanceApp() error {
	var err error

	doOnce.Do(func() {
		applicationInstance, err = initApplication()
	})
	return err
}

func Get() Application {
	return applicationInstance
}

func initApplication() (*applicationType, error) {
	applicationInstance = &applicationType{
		Name:          "default",
		Configuration: &entity.Configuration{},
	}
	if err := applicationInstance.initConfig(); err != nil {
		return applicationInstance, fmt.Errorf("initConfig() %w", err)
	}

	if err := applicationInstance.initLogger(); err != nil {
		return applicationInstance, fmt.Errorf("initLogger() %w", err)
	}
	applicationInstance.Recovery = NewRecoveryInterface(applicationInstance.Logger.Logger)

	if err := applicationInstance.initUtm(); err != nil {
		return applicationInstance, fmt.Errorf("initUtm() %w", err)
	}
	if err := applicationInstance.initDb(); err != nil {
		return applicationInstance, fmt.Errorf("initDb() %w", err)
	}
	// if err := appInst.initDb2(); err != nil {
	// 	return appInst, err
	// }
	if err := applicationInstance.Config.Set("application.console", isConsole(), true); err != nil {
		logger.ZeroLog().Error().AnErr("app error", err).Send()
	}
	return applicationInstance, nil
}

func (a *applicationType) initLogger() error {
	cfg := a.Configuration
	cc := &logger.ConfigZero{
		// ConsoleLoggingEnabled: cfg.Logging.ConsoleLoggingEnabled,
		ConsoleLoggingEnabled: isConsole(),
		NoColor:               cfg.Logging.NoColor,
		Debug:                 cfg.Debug,
		FileLoggingEnabled:    cfg.Logging.FileLoggingEnabled,
		Directory:             cfg.Logging.Directory,
		Filename:              cfg.Logging.LogFilename,
		// Filename:              "log.txt",
		MaxSize:    cfg.Logging.MaxSize,    // megabytes
		MaxBackups: cfg.Logging.MaxBackups, // files
		MaxAge:     cfg.Logging.MaxAge,     //days
		Truncate:   cfg.Logging.LogTruncate,
	}

	a.Logger = logger.InitLogger(cc)
	return nil
}

func (a *applicationType) GetLogger() *logger.Logger {
	return a.Logger
}

func (a *applicationType) initConfig() error {
	// Configuration
	config.TomlConfig = entity.TomlConfig
	// tt, err := config.GetInstance("config", a.Configuration, "debug")
	tt, err := config.GetInstance("config", a.Configuration)
	if err != nil {
		fmt.Printf("app initConfig() error = %v\n", err.Error())
		os.Exit(1)
	}
	a.Config = tt
	return nil
}

func (a *applicationType) GetConfig() *config.Config {
	return a.Config
}

func (a *applicationType) GetConfiguration() *entity.Configuration {
	return a.Configuration
}

func (a *applicationType) initDb() error {
	cfg := a.Configuration
	dbs, err := sqlite3.NewDbService(cfg.Database.DbName, sqlite3.RwModeWithCreate, sqlite3.Version)
	if err != nil {
		return fmt.Errorf("initDb %w", err)
	}
	a.Db = dbs
	// return fmt.Errorf("test error")
	return nil
}

func (a *applicationType) GetDbService() entity.DbService {
	// тут бы ошибку может какую а так это пустой код
	// if a.Db == nil {
	// 	return nil
	// }
	return a.Db
}

func (a *applicationType) initUtm() error {
	cfg := a.Configuration
	// оставляем так в начале приложения без пинга что короче, просто так пока надежней
	if infoUtm, err := a.GetUtmInfo(); err != nil {
		a.ErrorLog().AnErr("initUtm()", err).Send()
	} else {
		if infoUtm.OwnerId != "" {
			a.GetConfig().Set("application.fsrarid", infoUtm.OwnerId, true)
		}
	}
	if cfg.Application.Fsrarid == "" {
		MessageBox("Ошибка ФСРАР ИД", "FSRAR ID пустой. Функциональность программы ограничена.")
		// os.Exit(1)
	}
	// имя БД установим тут если он не прописано, это делаю чтобы один и тот же файл под разные фсрар ид пользовать
	if cfg.Database.DbName == "" {
		a.GetConfig().Set("database.dbname", cfg.Application.Fsrarid+".db", true)
	}
	return nil
}

func (a *applicationType) SetRepo(repo entity.Repo) {
	a.Repo = repo
}

func (a *applicationType) GetRepo() entity.Repo {
	return a.Repo
}

func (a *applicationType) InitUtm() error {
	return a.initUtm()
}

func (a *applicationType) InitDb() error {
	return a.initDb()
}

func (a *applicationType) InitLogger() error {
	return a.initLogger()
}

func (a *applicationType) GetTaskScheduler() entity.TaskList {
	return a.tasks
}
func (a *applicationType) SetTaskScheduler(t entity.TaskList) {
	a.tasks = t
}

func (a *applicationType) GetHistory() entity.History {
	return a.History
}
func (a *applicationType) SetHistory(h entity.History) {
	a.History = h
}
func (a *applicationType) GetRecovery() entity.RecoverInterface {
	return a.Recovery
}

func (a *applicationType) Shutdown() {
	entity.AppInterrupt <- 0
}

func (a *applicationType) Restart() {
	entity.AppInterrupt <- 1
}
