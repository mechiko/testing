package app_state

import (
	"testing/internal/entity"
)

type appstate struct {
	Repo entity.Repo
}

func New(r entity.Repo) entity.AppState {
	return &appstate{
		Repo: r,
	}
}
