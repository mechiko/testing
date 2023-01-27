package application

import (
	"fmt"
	"os"

	"testing/internal/entity"
	"testing/pkg/config"
	"testing/pkg/logger"
	"testing/pkg/sqlite3"
)

type applicationType struct {
	entity.App
	name          string
	config        entity.ConfigInterface
	logger        logger.Logger
	configuration *entity.Configuration
	db            entity.DbService
	repo          entity.Repo
	history       entity.History
	recovery      entity.RecoverInterface
	gui           entity.GuiService
	layout        string
	// Settings
	pwd     string
	output  string
	export  string
	browser string
	monitor entity.EventMonitor
	tasks   entity.TaskList
}

var _ entity.App = &applicationType{}

func NewApplication() (entity.App, error) {
	defer RecoverFmt("NewApplication()")
	if applicationInstance, err := initApplication(); err != nil {
		return nil, fmt.Errorf("application:NewApplication() %w", err)
	} else {
		return applicationInstance, nil
	}
}

// func Get() entity.App {
// 	return applicationInstance
// }

func initApplication() (*applicationType, error) {
	defer RecoverFmt("initApplication()")
	applicationInstance := &applicationType{
		name:          "default",
		configuration: &entity.Configuration{},
	}
	if err := applicationInstance.initConfig(); err != nil {
		return applicationInstance, fmt.Errorf("initConfig() %w", err)
	}
	applicationInstance.layout = applicationInstance.configuration.TimeLayout

	if err := applicationInstance.initLogger(); err != nil {
		return applicationInstance, fmt.Errorf("initLogger() %w", err)
	}
	applicationInstance.recovery = NewRecoveryInterface(applicationInstance.logger.Logger)

	if err := applicationInstance.initUtm(); err != nil {
		return applicationInstance, fmt.Errorf("initUtm() %w", err)
	}
	if err := applicationInstance.initDb(); err != nil {
		return applicationInstance, fmt.Errorf("initDb() %w", err)
	}
	// if err := appInst.initDb2(); err != nil {
	// 	return appInst, err
	// }
	if err := applicationInstance.config.Set("application.console", isConsole(), true); err != nil {
		applicationInstance.ErrorLog().AnErr("app error", err).Send()
	}
	return applicationInstance, nil
}

func (a *applicationType) initLogger() error {
	cfg := a.configuration
	cc := &logger.ConfigZero{
		// ConsoleLoggingEnabled: cfg.Logging.ConsoleLoggingEnabled,
		ConsoleLoggingEnabled: isConsole(),
		NoColor:               cfg.Logging.NoColor,
		Debug:                 cfg.Debug,
		FileLoggingEnabled:    cfg.Logging.FileLoggingEnabled,
		Directory:             cfg.Logging.Directory,
		Filename:              cfg.Logging.LogFilename,
		// Filename:              "log.txt",
		Rolling:    false,
		MaxSize:    cfg.Logging.MaxSize,    // megabytes
		MaxBackups: cfg.Logging.MaxBackups, // files
		MaxAge:     cfg.Logging.MaxAge,     //days
		Truncate:   cfg.Logging.LogTruncate,
	}

	a.logger = logger.NewLogger(cc)
	return nil
}

func (a *applicationType) GetLogger() logger.Logger {
	return a.logger
}

func (a *applicationType) initConfig() error {
	// Configuration
	config.TomlConfig = entity.TomlConfig
	// tt, err := config.GetInstance("config", a.Configuration, "debug")
	tt, err := config.NewInstance(a, "config", a.configuration)
	if err != nil {
		fmt.Printf("app initConfig() error = %v\n", err.Error())
		os.Exit(1)
	}
	a.config = tt
	return nil
}

func (a *applicationType) GetConfig() entity.ConfigInterface {
	return a.config
}

func (a *applicationType) GetConfiguration() *entity.Configuration {
	return a.configuration
}

func (a *applicationType) initDb() error {
	cfg := a.configuration
	dbs, err := sqlite3.NewDbService(cfg.Database.DbName, sqlite3.RwModeWithCreate, sqlite3.Version)
	if err != nil {
		return fmt.Errorf("initDb %w", err)
	}
	a.db = dbs
	// return fmt.Errorf("test error")
	return nil
}

func (a *applicationType) GetDbService() entity.DbService {
	// тут бы ошибку может какую а так это пустой код
	// if a.Db == nil {
	// 	return nil
	// }
	return a.db
}

func (a *applicationType) initUtm() error {
	cfg := a.configuration
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
	a.repo = repo
}

func (a *applicationType) GetRepo() entity.Repo {
	return a.repo
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
	return a.history
}
func (a *applicationType) SetHistory(h entity.History) {
	a.history = h
}
func (a *applicationType) GetRecovery() entity.RecoverInterface {
	return a.recovery
}
func (a *applicationType) SetGuiService(gi entity.GuiService) {
	a.gui = gi
}
func (a *applicationType) GetGuiService() entity.GuiService {
	return a.gui
}

func (a *applicationType) Shutdown() {
	a.logger.Dispose()
	entity.AppInterrupt <- 0
}

func (a *applicationType) Restart() {
	a.gui.Shutdown()
	a.logger.Dispose()
	entity.AppInterrupt <- 1
}
