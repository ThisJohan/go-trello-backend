package database

import (
	"github.com/ThisJohan/go-trello-clone/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.DB), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
