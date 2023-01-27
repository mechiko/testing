package gui

import (
	"fmt"

	"testing/internal/entity"
	// "testing/pkg/gui/dialogs"

	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

//	type MPMWInterface interface {
//		Update()
//	}
type MultiPageMainWindowConfig struct {
	Name                 string
	Enabled              dcl.Property
	Visible              dcl.Property
	Font                 dcl.Font
	MinSize              dcl.Size
	MaxSize              dcl.Size
	ContextMenuItems     []dcl.MenuItem
	OnKeyDown            walk.KeyEventHandler
	OnKeyPress           walk.KeyEventHandler
	OnKeyUp              walk.KeyEventHandler
	OnMouseDown          walk.MouseEventHandler
	OnMouseMove          walk.MouseEventHandler
	OnMouseUp            walk.MouseEventHandler
	OnSizeChanged        walk.EventHandler
	OnCurrentPageChanged walk.EventHandler
	Title                string
	Size                 dcl.Size
	MenuItems            []dcl.MenuItem
	ToolBar              dcl.ToolBar
}
type MultiPageMainWindow struct {
	*walk.MainWindow

	cfg *MultiPageMainWindowConfig
	gui *guiService

	tv  *walk.TreeView
	tvm *entity.AppmenuTreeModel
	// compos       *walk.Composite
	// callTreeMenu models.CallTreeMenu

	pageCom                     *walk.Composite
	currentPageChangedPublisher walk.EventPublisher
	SbiScan                     *walk.StatusBarItem
	SbiFsrarId                  *walk.StatusBarItem
	SbiUtmState                 *walk.StatusBarItem
	SbiState                    *walk.StatusBarItem
	SbiOptions                  *walk.StatusBarItem
	IconRed                     *walk.Icon
	IconGreen                   *walk.Icon
	IconWatch                   *walk.Icon
}

// func (mpmw *MultiPageMainWindow) initTreeMenu() error {
// 	tvm, err := models.NewAppMenuTreeModel()
// 	if err != nil {
// 		return err
// 	}
// 	mpmw.tvm = tvm
// 	err = mpmw.callTreeMenu(tvm)

// 	return err
// }

func (s *guiService) newMultiPageMainWindow() (*MultiPageMainWindow, error) {
	defer s.App.GetRecovery().RecoverLog("NewMultiPageMainWindow")
	mpmw := &MultiPageMainWindow{
		gui: s,
	}
	if _, err := s.App.GetMonitor().Attach(mpmw, entity.ALL); err != nil {
		return mpmw, fmt.Errorf("newMultiPageMainWindow() %w", err)
	}
	if _, err := s.App.GetMonitor().Attach(mpmw, entity.NEED_UPDATE_STATUS_BAR); err != nil {
		return mpmw, fmt.Errorf("newMultiPageMainWindow() %w", err)
	}
	return mpmw, nil
}

func (mpmw *MultiPageMainWindow) createMainWindow() error {
	defer mpmw.gui.App.GetRecovery().RecoverLog("CreateMainWindow")

	// icon2, err := walk.Resources.Icon("101")
	// if err != nil {
	// 	application.Get().ErrorLog().AnErr("guiService:[gui\\multipagemainwindow.go]", err).Send()
	// }
	if iconUtmRed, err := walk.Resources.Icon("108"); err != nil {
		// application.Get().ErrorLog().AnErr(`walk.Resources.Icon("108")`, err).Send()
		return fmt.Errorf(`walk.Resources.Icon("108") %w`, err)
	} else {
		mpmw.IconRed = iconUtmRed
	}
	iconUtmGreen, err := walk.Resources.Icon("109")
	if err != nil {
		// application.Get().ErrorLog().AnErr("guiService:[gui\\multipagemainwindow.go]", err).Send()
		return fmt.Errorf(`walk.Resources.Icon("109") %w`, err)
	} else {
		mpmw.IconGreen = iconUtmGreen
	}
	iconUtmScan, err := walk.Resources.Icon("106")
	if err != nil {
		// application.Get().ErrorLog().AnErr("guiService:[gui\\multipagemainwindow.go]", err).Send()
		return fmt.Errorf(`walk.Resources.Icon("106") %w`, err)
	}
	mpmw.IconWatch = iconUtmScan

	if err := (dcl.MainWindow{
		AssignTo:         &mpmw.MainWindow,
		Name:             mpmw.cfg.Name,
		Title:            mpmw.cfg.Title,
		Enabled:          mpmw.cfg.Enabled,
		Visible:          mpmw.cfg.Visible,
		Font:             mpmw.cfg.Font,
		MinSize:          mpmw.cfg.MinSize,
		MaxSize:          mpmw.cfg.MaxSize,
		MenuItems:        mpmw.cfg.MenuItems,
		ToolBar:          mpmw.cfg.ToolBar,
		ContextMenuItems: mpmw.cfg.ContextMenuItems,
		OnKeyDown:        mpmw.cfg.OnKeyDown,
		OnKeyPress:       mpmw.cfg.OnKeyPress,
		OnKeyUp:          mpmw.cfg.OnKeyUp,
		OnMouseDown:      mpmw.cfg.OnMouseDown,
		// OnMouseDown: mpmw.MouseDown,
		OnMouseMove: mpmw.cfg.OnMouseMove,
		OnMouseUp:   mpmw.cfg.OnMouseUp,
		// OnMouseUp:     mpmw.MouseUp,
		// OnSizeChanged: mpmw.sizeChanged,
		Layout: dcl.HBox{MarginsZero: true, SpacingZero: true},
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: true},
				// Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
				// Column:     2,
				Children: []dcl.Widget{
					dcl.HSplitter{
						Children: []dcl.Widget{
							dcl.Composite{
								Layout: dcl.VBox{MarginsZero: true, SpacingZero: true},
								// Background:    SolidColorBrush{walk.RGB(125, 50, 0)},
								// Alignment:     Alignment2D(walk.AlignHNearVNear),
								StretchFactor: 1,
								Children: []dcl.Widget{
									// VSpacer{
									// 	MaxSize: Size{10, 10},
									// 	MinSize: Size{10, 10},
									// },
									dcl.TreeView{
										AssignTo: &mpmw.tv,
										Model:    mpmw.tvm,
										OnCurrentItemChanged: func() {
											// core.App().ZeroLogger().Debug().Msgf("change %v", mpmw.сurrentPageTitle())
											mpmw.сhangePage()
										},
									},
								},
							},
							dcl.ScrollView{
								Layout:        dcl.VBox{MarginsZero: true, SpacingZero: true},
								Background:    dcl.SolidColorBrush{Color: walk.RGB(255, 255, 255)},
								StretchFactor: 5,
								Children: []dcl.Widget{
									dcl.Composite{
										AssignTo: &mpmw.pageCom,
										Name:     "placePage",
										Layout:   dcl.VBox{MarginsZero: true, SpacingZero: true},
										// Layout:   VBox{Margins: Margins{Left: 5, Right: 5, Top: 2, Bottom: 2}},
									},
									// CustomWidget{},
								},
							},
						},
					},
				},
			},
		},
		StatusBarItems: []dcl.StatusBarItem{
			{
				AssignTo:    &mpmw.SbiScan,
				Width:       100,
				ToolTipText: "Автоматический прием документов в УТМ",
				OnClicked:   mpmw.sbiScanPress,
			},
			{
				AssignTo: &mpmw.SbiFsrarId,
				Width:    100,
			},
			{
				AssignTo: &mpmw.SbiUtmState,
				Width:    100,
			},
			{
				AssignTo:  &mpmw.SbiState,
				Width:     400,
				Text:      "",
				OnClicked: mpmw.clickHistoryState,
			},
			// {
			// 	AssignTo: &mpmw.SbiOptions,
			// 	Width:    200,
			// 	Text:     "SbiOptions",
			// },
		}}).Create(); err != nil {
		return fmt.Errorf(`MPMW.Create() %w`, err)
	}

	// core.App().ZeroLogger().Debug().Msg("2")
	succeeded := false
	defer func() {
		if !succeeded {
			mpmw.Dispose()
		}
	}()

	mpmw.MainWindow.SetBounds(walk.Rectangle{X: mpmw.gui.X, Y: mpmw.gui.Y, Width: mpmw.gui.Width, Height: mpmw.gui.Height})

	mpmw.tv.SetCurrentItem(mpmw.tvm.DefaultMenu())

	mpmw.сurrentPageChanged().Attach(mpmw.cfg.OnCurrentPageChanged)
	mpmw.UpdateStatusBar()

	succeeded = true

	return nil
}

func (mpmw *MultiPageMainWindow) GetTreeViewModel() *entity.AppmenuTreeModel {
	return mpmw.tvm
}

func (mpmw *MultiPageMainWindow) SetTreeViewModel(t *entity.AppmenuTreeModel) {
	mpmw.tvm = t
}

// func (mpmw *MultiPageMainWindow) сurrentPage() models.Page {
// 	return mpmw.tvm.CurrentPage()
// }

// вызывается при смене страницы
func (mpmw *MultiPageMainWindow) сhangePage() error {
	defer mpmw.gui.App.GetRecovery().RecoverLog("сhangePage()")
	menu := mpmw.tv.CurrentItem().(*entity.AppMenu)
	if menu.Action() == nil {
		return nil
	}
	// mpmw.gui.App.DebugLog().Msgf("ChangePage() menu=%+v", menu.Name())
	// models.GetApp().LogMemUsage()
	return mpmw.setCurrentMenu(menu)
}

func (mpmw *MultiPageMainWindow) сurrentPageTitle() string {
	if mpmw.tvm.CurrentPage() == nil {
		// logger.ZeroLog().Debug().Msg("CurrentPageTitle() page=nil")
		return ""
	}
	return mpmw.tvm.CurrentMenu.Name()
}

func (mpmw *MultiPageMainWindow) сurrentPageChanged() *walk.Event {
	return mpmw.currentPageChangedPublisher.Event()
}

func (mpmw *MultiPageMainWindow) setCurrentMenu(pageMenu *entity.AppMenu) error {
	defer func() {
		// core.App().DebugLog().Int("Детей", mpmw.pageCom.Children().Len()).Bool("приемник.IsDisposed()", mpmw.pageCom.IsDisposed()).Msg("При выходе из смены страницы")
		if !mpmw.pageCom.IsDisposed() {
			mpmw.pageCom.RestoreState()
		}
	}()

	prevPage := mpmw.tvm.CurrentPage()

	if prevPage != nil {
		mpmw.pageCom.SaveState()
		prevPage.SetVisible(false)
		prevPage.(walk.Widget).SetParent(nil)
		prevPage.Disposing().Attach(func() {
		})
		prevPage.Dispose()
		prevPage.Clear()
	}

	newPage := mpmw.tvm.Menu2NewPage[pageMenu]

	if mpmw.pageCom.Children().Len() > 0 {
		mpmw.DisposeChildren(mpmw.pageCom)
	}

	page, err := newPage(mpmw.pageCom, mpmw.gui.App)
	if err != nil {
		return fmt.Errorf("newPage(mpmw.pageCom) %w", err)
	}

	mpmw.tvm.SetCurrentPage(page)
	mpmw.tvm.CurrentMenu = pageMenu

	// создаем событие смены страницы
	mpmw.currentPageChangedPublisher.Publish()

	mpmw.SetFocus()

	return nil
}

func (mpmw *MultiPageMainWindow) sbiScanPress() {
	// disconnected := mpmw.gui.Config.Viper.GetBool("application.disconnected")
	disconnected := mpmw.gui.App.GetConfiguration().Application.Disconnected
	// val := mpmw.gui.Config.Viper.GetBool("application.scanutm")
	val := mpmw.gui.App.GetConfiguration().Application.ScanUtm
	if !disconnected {
		_ = mpmw.gui.App.GetConfig().Set("application.scanutm", !val, true)
	} else {
		_ = mpmw.gui.App.GetConfig().Set("application.scanutm", false, true)
	}

	mpmw.UpdateStatusBar()
}

func (mpmw *MultiPageMainWindow) CheckChildren(wtest walk.Container) error {
	// ws := wtest.Children().Len()
	// for i := 0; i < ws; i++ {
	// 	w := wtest.Children().At(i)
	// 	mpmw.gui.Logger.Debug().Int("#widget", i).Bool("Disposed", w.IsDisposed()).Str("Name", w.Name()).Send()
	// }
	return nil
}

func (mpmw *MultiPageMainWindow) DisposeChildren(wtest walk.Container) error {
	defer mpmw.gui.App.GetRecovery().RecoverLog("DisposeChildren")
	// mpmw.gui.App.DebugLog().Msg("DisposeChildren")
	ws := wtest.Children().Len()
	// mpmw.gui.App.DebugLog().Int("Детей", ws).Msg("DisposeChildren")
	for i := 0; i < ws; i++ {
		w := wtest.Children().At(i)
		w.Dispose()
		// mpmw.gui.App.DebugLog().Int("#widget", i).Bool("Disposed", w.IsDisposed()).Str("Name", w.Name()).Send()
	}
	return nil
}

func (mpmw *MultiPageMainWindow) UpdateStatusBar() {
	defer mpmw.gui.App.GetRecovery().RecoverLog("MultiPageMainWindow:UpdateStatusBar")
	conf := mpmw.gui.App.GetConfiguration()
	if conf.Application.Disconnected {
		mpmw.SbiUtmState.SetText("УТМ Выкл")
		mpmw.SbiUtmState.SetIcon(mpmw.IconRed)
	} else {
		mpmw.SbiUtmState.SetText("УТМ Вкл")
		mpmw.SbiUtmState.SetIcon(mpmw.IconGreen)
	}
	if conf.Application.ScanUtm {
		mpmw.SbiScan.SetText("Прием Вкл")
		mpmw.SbiScan.SetIcon(mpmw.IconGreen)
	} else {
		mpmw.SbiScan.SetText("Прием Выкл")
		mpmw.SbiScan.SetIcon(mpmw.IconRed)
	}
	mpmw.SbiFsrarId.SetText(conf.Application.Fsrarid)
	s := mpmw.gui.App.GetHistory().Last()
	mpmw.SbiState.SetText(s)
}

func (mpmw *MultiPageMainWindow) UpdateStatusBarEvent(evt entity.EventInt, val string) {
	defer mpmw.gui.App.GetRecovery().RecoverLog("MultiPageMainWindow:UpdateStatusBarEvent")
	if mpmw.SbiState == nil {
		return
	}
	mpmw.SbiState.SetText(mpmw.gui.App.GetHistory().Last())
}

// as Subcriber on evt NEED_UPDATE_STATUS_BAR
func (mpmw *MultiPageMainWindow) NameSubcriber() string {
	return "MultiPageMainWindow"
}

func (mpmw *MultiPageMainWindow) UpdateSubcriber(val string, evt entity.EventInt) {
	defer mpmw.gui.App.GetRecovery().RecoverLog("MultiPageMainWindow:UpdateSubcriber")
	switch evt {
	case entity.NEED_TICKS_EVERY_SECOND:
	case entity.NEED_UPDATE_STATUS_BAR:
		mpmw.UpdateStatusBar()
	default:
		mpmw.UpdateStatusBarEvent(evt, val)
	}
}

func (mpmw *MultiPageMainWindow) clickHistoryState() {
	// dialogs.ViewHistory(mpmw)
}
