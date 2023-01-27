package pkg

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed vueapp
var vueapp embed.FS

func GetFileSystem() http.FileSystem {
	fsys, err := fs.Sub(vueapp, "vueapp")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
