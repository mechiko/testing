package gui

import (
	"bytes"
	"fmt"
	"os"

	"testing/internal/entity"

	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

// SettingsService allows querying, updating, and saving settings.
type GuiService interface {
	NewMainWindow() error
	GetMainWindow() walk.MainWindow
	// GetMPMW() MPMWInterface
	SetShutdown(f func())
	SetDefer(f func())
	SetTitle(t string)
	SetTreeMenu(t *entity.AppmenuTreeModel)
	AddAction(a *walk.Action)
	Actions() []*walk.Action
	Run() int
	// GetGlobalInterface() models.GlobalGuiService
	newMultiPageMainWindow() (*MultiPageMainWindow, error)
	Shutdown()
}

const (
	maxWidth  int = 600
	maxHeight int = 400
	// коэффициент уменьшения ширины и высоты, расчитывается как максимальная ширина реального экрана деленная на этот коэф
	k float32 = float32(3) / float32(5)
)

// type Gui struct {
// 	shutdown    func()
// 	ondefer     func()
// 	resourceDir string
// 	title       string
// 	MainWindow  *MultiPageMainWindow
// 	NotifyIcon  *walk.NotifyIcon
// 	tvm         *entity.AppmenuTreeModel
// 	niActions   []*walk.Action
// 	ScrWidth    int
// 	ScrHeight   int
// 	Width       int
// 	Height      int
// 	X           int
// 	Y           int
// }

type guiService struct {
	App entity.App
	// Db       entity.DbService
	// Logger   *zerolog.Logger
	// Recovery entity.RecoverInterface
	// Config   *config.Config
	shutdown    func()
	ondefer     func()
	resourceDir string
	title       string
	MainWindow  *MultiPageMainWindow
	NotifyIcon  *walk.NotifyIcon
	tvm         *entity.AppmenuTreeModel
	niActions   []*walk.Action
	ScrWidth    int
	ScrHeight   int
	Width       int
	Height      int
	X           int
	Y           int
}

var _ GuiService = &guiService{}

func NewGuiService(resourceDir string, a entity.App) GuiService {
	defer a.GetRecovery().RecoverLog("NewGuiService()")
	s := &guiService{
		App: a,
		// Db:       a.GetDbService(),
		// Logger:   a.GetLogger().Logger,
		// Recovery: a.GetRecovery(),
		// Config:   a.GetConfig(),
	}
	// s.Configuration = application.Get().GetConfiguration()

	s.niActions = make([]*walk.Action, 0)
	s.resourceDir = resourceDir
	s.Width = maxWidth
	s.Height = maxHeight
	hDC := win.GetDC(0)
	defer win.ReleaseDC(0, hDC)

	s.ScrWidth = int(win.GetDeviceCaps(hDC, win.HORZRES))
	s.ScrHeight = int(win.GetDeviceCaps(hDC, win.VERTRES))
	if s.ScrWidth < maxWidth {
		s.Width = s.ScrWidth - 100
	} else {
		s.Width = int(float32(s.ScrWidth) * k)
	}
	if s.ScrHeight < maxHeight {
		s.Height = s.ScrHeight - 50
	} else {
		s.Height = int(float32(s.ScrHeight) * k)
	}
	if (s.ScrWidth - s.Width) > 0 {
		s.X = (s.ScrWidth - s.Width) / 2
	} else {
		s.X = 0
	}
	if (s.ScrHeight - s.Height) > 0 {
		s.Y = (s.ScrHeight - s.Height) / 2
	} else {
		s.Y = 0
	}

	return s
}

func (s *guiService) AddAction(a *walk.Action) {
	s.niActions = append(s.niActions, a)
}

func (s *guiService) Actions() []*walk.Action {
	return s.niActions
}

func (s *guiService) SetDefer(f func()) {
	s.ondefer = f
}

func (s *guiService) SetShutdown(f func()) {
	s.shutdown = f
}

func (s *guiService) SetTreeMenu(t *entity.AppmenuTreeModel) {
	s.tvm = t
}

func (s *guiService) NewMainWindow() error {
	defer s.App.GetRecovery().RecoverLog("NewMainWindow()")
	walk.Resources.SetRootDirPath(s.resourceDir)
	cfg := &MultiPageMainWindowConfig{
		Name:    "mainWindow",
		MinSize: dcl.Size{Width: 300, Height: 200},
		MaxSize: dcl.Size{Width: s.Width, Height: s.Height},
		OnCurrentPageChanged: func() {
			s.updateTitle(s.MainWindow.сurrentPageTitle())
		},
		Visible: true,
	}
	mpmw, err := s.newMultiPageMainWindow()
	if err != nil {
		s.App.ErrorLog().AnErr("s.newMultiPageMainWindow()", err).Send()
		os.Exit(1)
	}
	if s.tvm == nil {
		s.App.ErrorLog().Msg("TreeViewModel не установлена")
		os.Exit(1)
	}
	mpmw.cfg = cfg
	mpmw.SetTreeViewModel(s.tvm)
	if err = mpmw.createMainWindow(); err != nil {
		s.App.ErrorLog().AnErr("mpmw.createMainWindow()", err).Send()
		os.Exit(1)
	}
	s.MainWindow = mpmw

	s.updateTitle(s.MainWindow.сurrentPageTitle())

	mpmw.MainWindow.Closing().Attach(s.Closing)

	return nil
}

func (s *guiService) Closing(canceled *bool, reason walk.CloseReason) {
	fmt.Printf("Closing() canceled=%v reason=%+v", *canceled, reason)
}

func (s *guiService) Run() int {
	// b := s.Config.Viper.GetBool("gui.tray")
	b := s.App.GetConfiguration().Gui.Tray
	if b {
		s.TrayWalk()
	}

	// тут тестируем пока сервис УТМ
	// svc, _ := utm.NewUtmService()
	// application.MessageBox("Адрес АПИ УТМ", svc.GetApiUri())

	ii := s.MainWindow.Run()
	fmt.Printf("END EXIT FROM MainWindow.Run() CODE=%v\n", ii)
	s.Defer()
	s.MainWindow.Dispose()
	return 0
}

func (s *guiService) Defer() {
	if s.NotifyIcon != nil {
		s.NotifyIcon.Dispose()
	}
	if s.ondefer != nil {
		s.ondefer()
	}
}

func (s *guiService) SetTitle(t string) {
	s.title = t
}

func (s *guiService) updateTitle(prefix string) {
	var buf bytes.Buffer

	if prefix != "" {
		buf.WriteString(prefix)
		buf.WriteString(" - ")
	}

	buf.WriteString(s.title)

	s.MainWindow.SetTitle(buf.String())
}

// func (s *guiService) GetGlobalInterface() models.GlobalGuiService {
// 	return s
// }

// models.GlobalGuiService

func (s *guiService) GetCurrentPage() (*entity.Page, error) {
	currentPage := s.tvm.CurrentPage()
	return &currentPage, nil
}

func (s *guiService) GetMainWindow() walk.MainWindow {
	return *s.MainWindow.MainWindow
}

func (s *guiService) Shutdown() {
	s.MainWindow.Close()
}
