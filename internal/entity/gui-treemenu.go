package entity

import (
	"log"

	"github.com/lxn/walk"
)

type CallTreeMenu func(tvm *AppmenuTreeModel) error

type AppMenuAction func(name string)

type AppMenu struct {
	name     string
	action   PageFactoryFunc
	image    string
	parent   *AppMenu
	children []*AppMenu
}

func (a *AppmenuTreeModel) NewRootAppMenu(name string, action PageFactoryFunc, image string) *AppMenu {
	am := &AppMenu{name: name, parent: nil, action: action, image: image}
	a.roots = append(a.roots, am)
	return am
}

var _ walk.TreeItem = new(AppMenu)

func (a *AppMenu) AddChild(name string, action PageFactoryFunc, image string) *AppMenu {
	am := &AppMenu{name: name, parent: a, action: action, image: image}
	a.children = append(a.children, am)
	return am
}

func (a *AppMenu) Text() string {
	return a.name
}
func (a *AppMenu) Name() string {
	return a.name
}
func (a *AppMenu) Action() PageFactoryFunc {
	return a.action
}

func (a *AppMenu) Parent() walk.TreeItem {
	if a.parent == nil {
		// We can't simply return d.parent in this case, because the interface
		// value then would not be nil.
		return nil
	}

	return a.parent
}

func (a *AppMenu) ChildCount() int {
	if a.children == nil {
		// It seems this is the first time our child count is checked, so we
		// use the opportunity to populate our direct children.
		if err := a.ResetChildren(); err != nil {
			log.Print(err)
		}
	}

	return len(a.children)
}

func (a *AppMenu) ChildAt(index int) walk.TreeItem {
	return a.children[index]
}

func (a *AppMenu) Image() interface{} {
	// log.Printf("func (a *AppMenu) Image() menu=%+v", a)
	// if a.tv != nil {
	// 	if a.tv.Expanded(a) {
	// 		return "minus.png"
	// 	}
	// }
	if a.image == "" {
		if a.action != nil {
			return "103"
		}
		if a.ChildCount() > 0 {
			return ""
		}
		return "103"
	}
	return a.image
}

func (a *AppMenu) ResetChildren() error {
	a.children = nil
	return nil
}

func (a *AppMenu) Path() string {
	return a.Image().(string)
	// if a.image == "" {
	// 	if a.action != nil {
	// 		return "document-properties.png"
	// 	}
	// 	if a.ChildCount() > 0 {
	// 		return "plus.png"
	// 	}
	// 	return "stop.ico"
	// }
	// return a.image
}

type AppmenuTreeModel struct {
	walk.TreeModelBase
	defaultMenu  *AppMenu
	CurrentMenu  *AppMenu
	currentPage  Page
	Menu2NewPage map[*AppMenu]PageFactoryFunc
	roots        []*AppMenu
}

var _ walk.TreeModel = new(AppmenuTreeModel)

func NewAppMenuTreeModel() *AppmenuTreeModel {
	model := new(AppmenuTreeModel)
	model.Menu2NewPage = make(map[*AppMenu]PageFactoryFunc)
	return model
}

func (*AppmenuTreeModel) LazyPopulation() bool {
	// We don't want to eagerly populate our tree view with the whole file system.
	return false
}

func (m *AppmenuTreeModel) RootCount() int {
	return len(m.roots)
}

func (m *AppmenuTreeModel) RootAt(index int) walk.TreeItem {
	return m.roots[index]
}

func (m *AppmenuTreeModel) SetDefaultMenu(a *AppMenu) {
	m.defaultMenu = a
}

func (m *AppmenuTreeModel) DefaultMenu() *AppMenu {
	return m.defaultMenu
}

func (m *AppmenuTreeModel) SetCurrentPage(p Page) {
	m.currentPage = p
}

func (m *AppmenuTreeModel) CurrentPage() Page {
	return m.currentPage
}
