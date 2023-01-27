package gui

import (
	"time"

	"github.com/lxn/walk"
)

var trayMouseDownHit bool = false
var timerTrayMouseDown *time.Timer

func (s *guiService) TrayWalk() {
	defer s.App.GetRecovery().RecoverLog("guiService TrayWalk()")
	// We need either a walk.MainWindow or a walk.Dialog for their message loop.
	// We will not make it visible in this example, though.

	icon, err := walk.Resources.Icon("7")
	if err != nil {
		s.App.ErrorLog().Msgf("guiservice Notify error %v", err.Error())
	}

	// Create the notify icon and make sure we clean it up on exit.
	ni, err := walk.NewNotifyIcon(s.MainWindow.Form())
	if err != nil {
		s.App.ErrorLog().Msgf("guiservice Notify error %v", err.Error())
	}
	s.NotifyIcon = ni
	// Set the icon and a tool tip text.
	if err := ni.SetIcon(icon); err != nil {
		s.App.ErrorLog().Msgf("guiservice Notify error %v", err.Error())
	}

	if err := ni.SetToolTip("Click for info or use the context menu to exit."); err != nil {
		s.App.ErrorLog().Msgf("guiservice Notify error %v", err.Error())
	}

	// When the left mouse button is pressed, bring up our balloon.
	ni.MouseDown().Attach(s.TrayMouseDown)
	ni.MouseUp().Attach(s.TrayMouseUp)

	for _, a := range s.Actions() {
		s.App.DebugLog().Msgf("add action to menu notify %v", a.Text())
		if err := ni.ContextMenu().Actions().Add(a); err != nil {
			s.App.ErrorLog().Msgf("guiservice Notify error %v", err.Error())
		}
	}

	// The notify icon is hidden initially, so we have to make it visible.
	if err := ni.SetVisible(true); err != nil {
		s.App.ErrorLog().Msgf("guiservice Notify error %v", err.Error())
	}

	// Now that the icon is visible, we can bring up an info balloon.
	// if err := ni.ShowInfo("Walk NotifyIcon Example", "Click the icon to show again."); err != nil {
	// 	log.Fatal().Msgf("guiservice Notify error %v", err.Error())
	// }
}

func (s *guiService) TrayMouseDown(x, y int, button walk.MouseButton) {
	defer s.App.GetRecovery().RecoverLog("guiService TrayMouseDown()")

	s.App.DebugLog().Msgf("trayMouseDown hit=%v", trayMouseDownHit)
	if button != walk.LeftButton {
		return
	}
	if trayMouseDownHit {
		// если за время для двойного клика успеваем, то двойной клик
		s.App.DebugLog().Msg("TrayMouseDown Double Click")
	}
}

// на отпускание мыши сделаем задержку для определения двойного нажатия
func (s *guiService) TrayMouseUp(x, y int, button walk.MouseButton) {
	defer s.App.GetRecovery().RecoverLog("guiService TrayMouseUp()")

	if button != walk.LeftButton {
		return
	}
	if trayMouseDownHit {
		// если за время для двойного клика успеваем, то двойной клик
		vis := s.MainWindow.Visible()
		s.MainWindow.SetVisible(!vis)
	} else {
		trayMouseDownHit = true
		timerTrayMouseDown = time.NewTimer(time.Millisecond * 500)
	}
	go func() {
		<-timerTrayMouseDown.C
		trayMouseDownHit = false
	}()
}
