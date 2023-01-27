// Package postgres implements postgres connection.
package sqlite3

import (
	"context"
	"fmt"
	"os"

	// "golang.org/x/sync/semaphore"

	// _ "github.com/mattn/go-sqlite3"
	// "github.com/Masterminds/squirrel"
	"testing/internal/entity"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

// type DbService interface {
// 	Ping() (*sqlx.DB, error)
// 	CheckDb() error
// 	Db() (*sqlx.DB, error)
// 	Close()
// 	GetName() string
// 	IsCreated() bool
// 	SetCreated()
// 	// GetRequestRules() (*entity.RequestRuleList, error)
// }

type dbService struct {
	ctx           context.Context
	sem           Semaphore
	ConnectionUri string
	Driver        string
	DbName        string
	Mode          UriType
	Version       CheckMode
	Created       bool
	// timeout       int
}

//	type dbService struct {
//		dbStruct
//	}
type Semaphore interface {
	Acquire()
	Release()
}
type semaphore struct {
	semC chan struct{}
}

func New(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}
func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}
func (s *semaphore) Release() {
	<-s.semC
}

// глубина канала сколько одновременно может работать с БД всегда с sqlite один!
const maxWorkers = 1

var _ entity.DbService = &dbService{}

// dbs, err := sqlite3.NewDbService(cfg.Database.DbName, sqlite3.RwModeWithCreate, sqlite3.Version)
// sqlite3.RwModeWithCreate константа в enum_types.go (sqlite3.RwModeWithCreate sqlite3.RoMode sqlite3.RwMode)
// versioning CheckMode сюда передаем sqlite3.Version если версионность поддерживается,
//
//	и sqlite3.NoMatter если не имеет значения
//
// имя БД файла определяется в (a *applicationType) initDb() и задается по фсрар ид если УТМ запущен берем там, если нет, из конфига
// a.GetConfig().Set("database.dbname", cfg.Application.Fsrarid+".db", true)
func NewDbService(dbname string, mode UriType, versioning CheckMode) (entity.DbService, error) {
	s := &dbService{}
	s.DbName = dbname
	s.Driver = "sqlite"
	s.Mode = mode
	s.ConnectionUri = mode.String()
	s.Version = versioning
	s.Created = false
	s.ctx = context.TODO()
	s.sem = New(maxWorkers)
	if err := s.CheckDb(); err != nil {
		return s, fmt.Errorf("s.CheckDb() %w", err)
	}
	return s, nil
}

func (s *dbService) GetName() string {
	return s.DbName
}

func (s *dbService) Ping() (*sqlx.DB, error) {
	var result error
	if err := s.CheckDb(); err != nil {
		return nil, fmt.Errorf("s.CheckDb() %w", err)
	}
	db, err := sqlx.Open(s.Driver, s.DbName+s.ConnectionUri)
	if err != nil {
		return nil, fmt.Errorf("файл %s не существует %w", s.DbName, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			result = fmt.Errorf("db.Close() %w", err)
		}
	}()

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping() %w", err)
	}

	return db, result
}

func (s *dbService) CheckDb() error {
	// fmt.Printf("CheckDb() для %s режим создания:'%s' режим контроля версии:'%s'\n\r", s.DbName, s.Mode.String(), s.Version.String())
	_, err := os.Stat(s.DbName)
	if os.IsNotExist(err) {
		if s.Mode == RwModeWithCreate {
			if err := s.create(); err != nil {
				return fmt.Errorf("s.create() %w", err)
			}
			// создали без ошибки
			s.Created = false
			return nil
		} else {
			// открытие без создания и файла нет
			return fmt.Errorf("os.IsNotExist(err) %w", err)
		}
	} else if err != nil {
		// какая то другая ошибка кроме наличия файла
		return fmt.Errorf("файл %s ошибка %s %w", s.DbName, err.Error(), err)
	}
	s.Created = true
	return nil
}

func (s *dbService) Db() (*sqlx.DB, error) {
	s.sem.Acquire()
	err := s.CheckDb()
	if err != nil {
		return nil, fmt.Errorf("s.CheckDb() %w", err)
	}
	dbConn, err := sqlx.Open(s.Driver, s.DbName+s.ConnectionUri)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open() %w", err)
	}
	dbConn.SetMaxOpenConns(0)
	return dbConn, nil
}

func (s *dbService) Close() {
	// fmt.Printf("Close dbService DB file:%s\n\r", s.DbName)
	s.sem.Release()
}

// func (s *dbService)	GetRequestRules() (*entity.RequestRuleList, error) {
// }
func (s *dbService) create() error {
	var result error
	db, err := sqlx.Open(s.Driver, s.DbName+s.ConnectionUri)
	if err != nil {
		return fmt.Errorf("sqlx.Open() %w", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			result = fmt.Errorf("db.Close() %w", err)
		}
	}()
	return result
}

func (s *dbService) IsCreated() bool {
	return s.Created
}

func (s *dbService) SetCreated() {
	s.Created = true
}
