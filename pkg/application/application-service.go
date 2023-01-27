package application

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/skratchdot/open-golang/open"
)

const DateFormat = "2006.01.02"

// это для повторной инициализации конфигурации повторное чтение данных
func (a *applicationType) InitConfiguration() error {
	defer a.GetRecovery().RecoverFmtFunc("(a *application) InitConfiguration()")
	var err error
	var cfg = a.configuration

	if a.pwd, err = os.Getwd(); err != nil {
		return fmt.Errorf("InitConfiguration() %w", err)
	}
	if err = a.GetConfig().Unmarshal(cfg); err != nil {
		return fmt.Errorf("InitConfiguration() %w", err)
	}
	if a.configuration.Gui.Output == "" {
		a.SetOutput("output")
	} else {
		a.SetOutput(a.configuration.Gui.Output)
	}
	a.export = a.configuration.Gui.Export
	a.browser = a.configuration.Browser
	return nil
}

func (a *applicationType) SaveConfig() error {
	// vp, err := config.GetViper()
	// if err != nil {
	// 	return fmt.Errorf("GetViper() %w", err)
	// }
	if err := a.config.SaveAs("cccc.ccc"); err != nil {
		return fmt.Errorf("application:SaveConfig() %w", err)
	}
	return nil
}

func (a *applicationType) GetExport() string {
	return a.export
}

func (a *applicationType) GetBrowser() string {
	return a.browser
}

func (a *applicationType) SetBrowser(e string) error {
	if a.browser != e {
		a.browser = e
		if err := a.GetConfig().Set("browser", a.browser, true); err != nil {
			return fmt.Errorf("SetBrowser() %w", err)

		}
	}
	return nil
}

func (a *applicationType) SetExport(e string) error {
	if a.export != e {
		a.export = e
		// s.Config().Gui.Export = s.export
		// if err := s.Config().Set("gui.export", s.export); err != nil {
		// 	return errors.Errorf("core Settings SetBrowser() %s", err.Error())
		// }
		// if err := s.Config().Save(); err != nil {
		// 	return errors.Errorf("core Settings SetExport() %s", err.Error())
		// }
	}
	return nil
}

func (a *applicationType) SetOutput(e string) {
	pathOut := filepath.Join(a.pwd, e)
	if _, err := os.Stat(pathOut); os.IsNotExist(err) {
		// pathOut does exist
		os.Mkdir(pathOut, 0755)
	}
	a.output = e
}

func (a *applicationType) GetPwd() string {
	return a.pwd
}

func (a *applicationType) GetBaseUrl() string {
	// var cfg = &entity.Configuration{}
	// config.GetConfig().Viper.Unmarshal(cfg)

	uri := "http://" + a.configuration.Hostname + ":" + a.configuration.HostPort
	return uri
}

// пусть вывода экспорта и других файлов выгрузки
func (a *applicationType) GetOutput() string {
	pathOut := filepath.Join(a.pwd, a.output)
	if _, err := os.Stat(pathOut); os.IsNotExist(err) {
		// pathOut does not exist
		os.Mkdir(pathOut, 0755)
	}
	return pathOut
}

func (a *applicationType) DebugLog() *zerolog.Event {
	return a.logger.Logger.Debug()
}

func (a *applicationType) InfoLog() *zerolog.Event {
	return a.logger.Logger.Info()
}

// ErrorLog() *zerolog.Event
func (a *applicationType) ErrorLog() *zerolog.Event {
	return a.logger.Logger.Error()
}

func (a *applicationType) OpenDir() {
	defer a.GetRecovery().RecoverFmtFunc("application OpenDir()")

	if a.output == "" {
		return
	}
	if err := open.Run(a.output); err != nil {
		a.ErrorLog().Str("Dir", a.output).AnErr("OpenDir()", err).Send()
	}
}

func (a *applicationType) Open(url string) {
	defer a.GetRecovery().RecoverFmtFunc("application Open()")

	if url == "" {
		return
	}

	if a.browser != "" {
		if err := open.RunWith(url, a.browser); err != nil {
			a.ErrorLog().Str("URL", url).AnErr("open.RunWith()", err).Send()
		}

	} else {
		if err := open.Run(url); err != nil {
			a.ErrorLog().Str("URL", url).AnErr("open.Run()", err).Send()
		}
	}

}
