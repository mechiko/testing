package entity

import (
	"github.com/lxn/walk"
)

// var pages = []PageConfig{
// 	{"Foo", "document-new.png", newFooPage},
// 	{"Bar", "document-properties.png", newBarPage},
// 	{"Baz", "system-shutdown.png", newBazPage},
// }
type PageFactoryFunc func(parent walk.Container, a App) (Page, error)

type Page interface {
	walk.Container
	// Root() walk.Container
	Parent() walk.Container
	SetParent(parent walk.Container) error
	Clear()
	Update()
}

type PageConfig struct {
	Title   string
	Image   string
	NewPage PageFactoryFunc
}
