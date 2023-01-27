package sqlite3

import (
	"testing/internal/entity"
	"testing/pkg/repo/sqlite3/app_state"

	"github.com/rs/zerolog"
	// "testing/internal/entity/clients"
	// "testing/internal/entity/form1"
	// "testing/internal/entity/form2"
	// "testing/internal/entity/nodesout"
	// "testing/internal/entity/request_rests"
	// "testing/internal/entity/requests"
	// "testing/internal/entity/rests"
	// "testing/internal/entity/rules"
)

type Repository struct {
	App      entity.App
	Db       entity.DbService
	Logger   zerolog.Logger
	Recovery entity.RecoverInterface
	Config   entity.ConfigInterface
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewRepository(a entity.App) entity.Repo {
	return &Repository{
		App:      a,
		Db:       a.GetDbService(),
		Logger:   a.GetLogger().Logger,
		Recovery: a.GetRecovery(),
		Config:   a.GetConfig(),
	}
}

func (r *Repository) DbService() entity.DbService {
	return r.Db
}

func (r *Repository) GetAppState() entity.AppState {
	return app_state.New(r)
}
