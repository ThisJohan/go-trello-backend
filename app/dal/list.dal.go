package dal

import (
	"github.com/ThisJohan/go-trello-clone/config/database"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	BoardId     uint `gorm:"foreignKey:board_id"`
	Lists       []List
}

func CreateList(l *List) *gorm.DB {
	return database.DB.Create(l)
}

func FindListsByBoard(dest interface{}, boardId interface{}) *gorm.DB {
	return database.DB.Model(&List{}).Find(dest, "board_id = ?", boardId)
}

func FindListById(dest interface{}, id interface{}) *gorm.DB {
	return database.DB.Model(&List{}).First(dest, "id = ?", id)
}

func DeleteList(id interface{}) *gorm.DB {
	return database.DB.Delete(&List{}, id)
}

func UpdateList(id interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&List{}).Where("id = ?", id).Updates(data)
}
