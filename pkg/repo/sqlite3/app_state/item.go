package app_state

import (
	"fmt"
)

func (r *appstate) Set(module string, key string, value string) error {
	query := `INSERT OR REPLACE INTO app_state (module, key, value) values (?, ?, ?);`
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, module, key, value); err == nil {
		return fmt.Errorf("result.LastInsertId() %w", err)
	}
	return nil
}
