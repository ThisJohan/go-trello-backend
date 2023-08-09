package dal

import (
	"github.com/ThisJohan/go-trello-clone/config/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func CreateUser(user *User) *gorm.DB {
	return database.DB.Create(user)
}
