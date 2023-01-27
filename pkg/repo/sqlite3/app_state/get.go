package app_state

import (
	"fmt"
)

//	Set(module string, key string, value string) error

func (r *appstate) Get(module string, key string) (val string, result error) {
	value := ""
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return val, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		r.Repo.DbService().Close()
		if err := db.Close(); err != nil {
			val = ""
			result = err
		}
	}()

	query := `select value from app_state where module = ? and key = ?;`

	err = db.Get(&value, query, module, key)
	if err != nil {
		return value, fmt.Errorf("db.Select() %s", err)
	}
	return value, nil
}
