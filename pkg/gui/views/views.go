package views

import (
	// "testing/pkg/gui/views/requests"
	// "testing/pkg/gui/views/rests"
	// "testing/pkg/gui/views/tasks"

	"testing/internal/entity"
)

var pages = map[string]entity.PageConfig{
	"Home": {Title: "Старт", Image: "104", NewPage: newHomePage},
	// "Requests": {Title: "Запросы", Image: "107", NewPage: requests.NewPage},
	// "Rests":    {Title: "Остатки", Image: "107", NewPage: rests.NewPage},
	// "Tasks":    {Title: "Задания", Image: "107", NewPage: tasks.NewPage},
}

func CreateTreeMenu() *entity.AppmenuTreeModel {
	tvm := entity.NewAppMenuTreeModel()
	page1 := pages["Home"]
	home := tvm.NewRootAppMenu(page1.Title, page1.NewPage, page1.Image)
	tvm.SetDefaultMenu(home)
	tvm.Menu2NewPage[home] = page1.NewPage

	// page2 := pages["Requests"]
	// menu2 := tvm.NewRootAppMenu(page2.Title, page2.NewPage, page2.Image)
	// tvm.Menu2NewPage[menu2] = page2.NewPage

	// page3 := pages["Rests"]
	// menu3 := tvm.NewRootAppMenu(page3.Title, page3.NewPage, page3.Image)
	// tvm.Menu2NewPage[menu3] = page3.NewPage

	// page4 := pages["Tasks"]
	// menu4 := tvm.NewRootAppMenu(page4.Title, page4.NewPage, page4.Image)
	// tvm.Menu2NewPage[menu4] = page4.NewPage

	// page5 := pages["AutoRequests"]
	// menu5 := tvm.NewRootAppMenu(page5.Title, page5.NewPage, page5.Image)
	// tvm.Menu2NewPage[menu5] = page5.NewPage
	return tvm
}
