package sqlite3

import (
	_ "embed"
	"fmt"

	"github.com/hashicorp/go-version"

	// этот драйвер не зависит от CGO поэтому не проблема для 64 бит
	_ "modernc.org/sqlite"
)

//goland:noinspection ALL
const createVersion string = `
CREATE TABLE if not exists dboptions (
	name TEXT NOT NULL DEFAULT '',
	value TEXT NOT NULL DEFAULT '',
	PRIMARY KEY (name)
);`

const setVersion string = `INSERT OR REPLACE INTO dboptions (name, value) VALUES ('version',?)`

const getVersion string = `select value from dboptions where name = 'version' limit 1;`

const ExeVersion string = "1.0.1"

//go:embed createDB.sql
var CreateDB string

//go:embed updateDB.sql
var UpdateDB string

//go:embed index.sql
var IndexDB string

func (r *Repository) CheckVersionDb() (result error) {
	var ver string
	db, err := r.App.GetDbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		r.App.GetDbService().Close()
		if err := db.Close(); err != nil {
			result = fmt.Errorf("db.Close() %w", err)
		}
	}()
	// проверяем есть ли вообще версия в БД
	if err := db.Get(&ver, getVersion); err != nil {
		return fmt.Errorf("db.Get() %w", err)
	}
	if err := r.updateVersionDb(ver); err != nil {
		return fmt.Errorf("updateVersionDb(ver) %w", err)
	}
	return result
}

// create() создаем БД сначала таблицу для опций, номер версии туда, и скрипт s.CreateDb
func (r *Repository) create() error {
	var result error
	db, err := r.App.GetDbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		r.App.GetDbService().Close()
		if err := db.Close(); err != nil {
			result = fmt.Errorf("db.Close() %w", err)
		}
	}()
	if _, err = db.Exec(createVersion); err != nil {
		return fmt.Errorf("db.Exec(createVersion) %w", err)
	}
	if _, err = db.Exec(setVersion, ExeVersion); err != nil {
		return fmt.Errorf("db.Exec(setVersion, ExeVersion) %w", err)
	}
	if _, err = db.Exec(CreateDB, ExeVersion); err != nil {
		return fmt.Errorf("db.Exec(CreateDB, ExeVersion) %w", err)
	}
	if _, err = db.Exec(IndexDB, ExeVersion); err != nil {
		return fmt.Errorf("db.Exec(IndexDB, ExeVersion) %w", err)
	}
	return result
}

// upgradeDb() обновляем БД и номер версии туда скрипт s.UpdateDb
func (r *Repository) upgradeDb() error {
	var result error
	db, err := r.App.GetDbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		r.App.GetDbService().Close()
		if err := db.Close(); err != nil {
			result = fmt.Errorf("db.Close() %w", err)
		}
	}()
	// сначала пытаемся обновить БД
	if UpdateDB != "" {
		if _, err = db.Exec(UpdateDB); err != nil {
			return fmt.Errorf("db.Exec(UpdateDB) %w", err)
		}
	}
	if _, err = db.Exec(setVersion, ExeVersion); err != nil {
		return fmt.Errorf("db.Exec(setVersion, ExeVersion) %w", err)
	}
	return result
}

// updateVersionDb(dbVersion string) если надо то обновляем
func (r *Repository) updateVersionDb(dbVersion string) error {
	v1, _ := version.NewVersion(ExeVersion)
	v2, _ := version.NewVersion(dbVersion)
	if v2.LessThan(v1) {
		if err := r.upgradeDb(); err != nil {
			return fmt.Errorf("r.upgradeDb() %w", err)
		}
	}
	return nil
}
