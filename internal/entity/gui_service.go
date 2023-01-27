package entity

import (
	"github.com/lxn/walk"
)

type GuiService interface {
	NewMainWindow() error
	GetMainWindow() walk.MainWindow
	// GetMPMW() MPMWInterface
	SetShutdown(f func())
	SetDefer(f func())
	SetTitle(t string)
	SetTreeMenu(t *AppmenuTreeModel)
	AddAction(a *walk.Action)
	Actions() []*walk.Action
	Run() int
	// GetGlobalInterface() models.GlobalGuiService
	// newMultiPageMainWindow() (*MultiPageMainWindow, error)
	Shutdown()
}
