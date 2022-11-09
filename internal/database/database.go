package database

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
	Logger *zerolog.Logger
}

func New(conn *gorm.DB, logger *zerolog.Logger) *Database {
	return &Database{
		Conn: conn,
		Logger: logger,
	}
}