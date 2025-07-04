package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const ON = "1"

type DB struct {
	Conn *gorm.DB
}

func Init(dsn, sqlDebug string) (DB, error) {
	logLevel := logger.Silent
	if sqlDebug == ON {
		logLevel = logger.Info
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		panic(err)
	}

	return DB{Conn: db}, nil
}
