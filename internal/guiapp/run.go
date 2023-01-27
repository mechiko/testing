// Package app configures and runs application.
package guiapp

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"

	utm "testing/internal/controller/utm/v1"
	"testing/internal/entity"
	"testing/pkg/application"
	"testing/pkg/events"
	"testing/pkg/gui"
	"testing/pkg/gui/views"
	"testing/pkg/httpserver"
	"testing/pkg/repo/sqlite3"
	"testing/pkg/tasks"

	"github.com/lxn/walk"
)

// меняем каталог на каталог запуска это не помню от куда велосипед
func init() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)
}

// Run creates objects via constructors.
func Run() int {
	var err error
	var app entity.App

	if app, err = application.NewApplication(); err != nil {
		fmt.Printf("Ошибка application.NewApplication() инициализации приложения %s\n", err.Error())
		os.Exit(1)
	}
	// инициализируем если забыли
	// if !application.IsInit() {
	// 	fmt.Printf("Ошибка Run() инициализации приложения %s\n", err.Error())
	// 	os.Exit(1)
	// }
	defer app.GetRecovery().RecoverLog("cmdapp.Run()")

	app.InfoLog().Str("cmdapp.Run() Database.DbName", app.GetConfiguration().Database.DbName).Send()

	// инициализируем REPO тут как реализацию всего доступа к частям системы через таблицы и драйвер
	repo := sqlite3.NewRepository(app)
	app.SetRepo(repo)
	if err := repo.Start(); err != nil {
		app.ErrorLog().AnErr("repo.Start()", err).Send()
		walk.MsgBox(nil, "Ошбика БД", err.Error(), walk.MsgBoxOK)
	}

	handler := echo.New()

	// v1.NewRouter(handler, *exampleUseCase)
	// assetHandler := http.FileServer(pkg.GetFileSystem())
	// handler.GET("/*", echo.WrapHandler(assetHandler))

	httpServer := httpserver.New(handler, httpserver.Port(app.GetConfiguration().HostPort))

	// тикер для считывания документов из УТМ каждую минуту
	conf := app.GetConfiguration()
	stopScan := make(chan bool)
	scanTimer := conf.Application.ScanTimer
	if scanTimer <= 0 {
		scanTimer = 30
	}
	utmStateTimer := conf.Application.UtmStateTimer
	if utmStateTimer <= 0 {
		utmStateTimer = 10
	}
	requestSendTimer := conf.Application.RequestSendTimer
	if requestSendTimer <= 0 {
		requestSendTimer = 60
	}
	requestSendCount := conf.Application.RequestSendCount
	if requestSendCount <= 0 {
		requestSendCount = 20
	}
	scanTicker := time.NewTicker(time.Duration(scanTimer) * time.Second)
	utmStateTicker := time.NewTicker(time.Duration(utmStateTimer) * time.Second)
	oneSecondTicker := time.NewTicker(time.Second)
	requestTicker := time.NewTicker(time.Duration(requestSendTimer) * time.Second)
	// Go function
	go func() {
		// Using for loop
		for {
			// Select statement
			select {
			case <-oneSecondTicker.C:
				// app.GetMonitor().Notify(entity.NEED_TICKS_EVERY_SECOND, "app.run() <-oneSecondTicker.C")
			case <-stopScan:
				app.DebugLog().Msg("on exit stop all ticker")
				scanTicker.Stop()
				utmStateTicker.Stop()
				return
			// case utmState
			case <-utmStateTicker.C:
				if utmSvc, err := utm.NewUtmService(app); err != nil {
					app.ErrorLog().AnErr("utm.NewUtmService()", err).Send()
				} else {
					utmSvc.Ping()
					// mpmwi.Update()
					app.GetMonitor().Notify(entity.NEED_UPDATE_STATUS_BAR, "app.run() <-utmStateTicker.C")
				}
			// Case to scan
			case <-scanTicker.C:
				scanTicker.Stop()
				startTime := time.Now()
				fmt.Printf("Получение документов заняло  %s секунд\n", time.Until(startTime).Abs())
				scanTicker.Reset(time.Duration(scanTimer) * time.Second)
			// Case to request send to utm
			case <-requestTicker.C:
				requestTicker.Stop()
				startTimeSend := time.Now()
				fmt.Printf("requestTicker.Stop() %s \n", time.Now())
				secs := time.Until(startTimeSend).Abs()
				dur := time.Duration(requestSendTimer)*time.Second - secs
				tDur := time.Duration(requestSendTimer) * time.Second
				fmt.Printf("End time %s  secs:%s dur:%s tdur:%s\n", time.Now(), secs, dur, tDur)
				fmt.Printf("Отправка документов заняла  %s \n", secs)
				// fmt.Printf("requestTicker timerdur %s \n", tDur)
				// fmt.Printf("requestTicker dur %s \n", dur)
				if secs > tDur {
					requestTicker.Reset(tDur)
				} else {
					requestTicker.Reset(dur)
				}
			}
		}
	}()

	// тут запуск оконного интерфейса
	go func() {
		defer app.GetRecovery().RecoverFmt("запуск оконного интерфейса")
		fmt.Println("start go GUI!")
		// GUI
		gui := gui.NewGuiService("", app)
		app.SetGuiService(gui)
		// список действий тоже до создания окна наверное
		app.DebugLog().Int("gui.Actions.len()", len(gui.Actions())).Send()

		exitAction := walk.NewAction()
		if err := exitAction.SetText("E&xit"); err != nil {
			app.ErrorLog().AnErr(`exitAction.SetText("E&xit")`, err).Send()
		}
		exitAction.Triggered().Attach(Shutdown)
		gui.AddAction(exitAction)
		openAction := walk.NewAction()
		if err := openAction.SetText("Открыть браузер"); err != nil {
			app.ErrorLog().AnErr(`openAction.SetText("Открыть браузер")`, err).Send()
		}
		openAction.Triggered().Attach(Open)
		gui.AddAction(openAction)
		app.DebugLog().Int("gui.Actions.len()", len(gui.Actions())).Send()

		// дерево меню инициализируем до создания главного окна
		tm := views.CreateTreeMenu()
		gui.SetTreeMenu(tm)
		gui.NewMainWindow()
		// вот теперь между созданием главного окна и запуском можем что то менять
		mw := gui.GetMainWindow()
		// mpmwi := gui.GetMPMW()
		mw.SetVisible(true)
		// Interrupt <- gui.Run()
		entity.AppInterrupt <- gui.Run()
	}()

	// инициализируем монитор событий тут внешне по отношнеию к application
	// где либо далее будет пробовать регистрировать слушателя с методами Name() и Update()
	// через вызов монитора Notify(evt EventInt) оповещать слушателей о нужном событии
	// History array
	history := entity.NewHistory(100)
	app.SetHistory(history)

	monitor := events.NewMonitor(app)
	app.SetMonitor(monitor)

	// TasksList
	task := tasks.NewTaskScheduler(app)
	app.SetTaskScheduler(task)
	if err := task.Load(); err != nil {
		app.ErrorLog().AnErr("Run", err).Send()
	}

	exitCode := 0
	// for n := 1; n > 0; {
	for n := 1; n > 0; {
		select {
		case api := <-entity.AppInterrupt:
			app.InfoLog().Msgf("[internal|app][run] app - Shutdown - int: %v channel - %v", api, len(entity.AppInterrupt))
			if len(entity.AppInterrupt) > 0 {
				<-entity.AppInterrupt
			}
			exitCode = api
			n--
		// case s := <-Interrupt:
		// 	app.InfoLog().Msgf("[internal|app][run] app - Run - signal: %v", s.String())
		// 	n--
		case err = <-httpServer.Notify():
			app.ErrorLog().AnErr("[internal|app][run] app - Run - httpServer.Notify", err).Send()
			n--
		}
	}

	// Shutdown
	stopScan <- true

	err = httpServer.Shutdown()
	if err != nil {
		app.ErrorLog().AnErr("[internal|app][run] app - Run - httpServer.Shutdown", err).Send()
	}

	return exitCode
}

// автоматическое открытие приложения в браузере
func Open() {
	// app.Logging().ZeroLogger().Info().Msg("main:[main.go] Open()")
	// uri := "http://" + core.App().Config().Hostname + ":" + core.App().Config().HostPort
	// core.App().Open(uri, core.App().Config().Browser)
}

// вызывается по выбору в трее выход
func Shutdown() {
	// app.Logging().ZeroLogger().Info().Msg("main:[main.go] Shutdown()")
	// walk.App().Exit(0)
}
