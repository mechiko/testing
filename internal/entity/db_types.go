package entity

import (
	"github.com/jmoiron/sqlx"
)

type DbService interface {
	Ping() (*sqlx.DB, error)
	CheckDb() error
	Db() (*sqlx.DB, error)
	Close()
	GetName() string
	IsCreated() bool
	SetCreated()
}
