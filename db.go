package database

import (
	"fmt"
	"github.com/tarasov-krk/env"
	"github.com/tarasov-krk/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	var (
		dbHost     = env.Get("DB_HOST")
		dbPort     = env.Get("DB_PORT")
		dbUserName = env.Get("DB_USERNAME")
		dbUserPass = env.Get("DB_PASSWORD")
		dbName     = env.Get("DB_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUserName, dbUserPass, dbHost, dbPort, dbName)
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	db, err := gormDb.DB()
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		logger.Error(err)
		panic(err)
	}

	DB = gormDb

	return DB
}

func Instance() *gorm.DB {
	return DB
}
