package dal

import (
	"github.com/ThisJohan/go-trello-clone/config/database"
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	UserID      uint `gorm:"foreignKey:user_id"`
}

func CreateBoard(b *Board) *gorm.DB {
	return database.DB.Create(b)
}

func FindBoardsByUser(dest interface{}, userId interface{}) *gorm.DB {
	return database.DB.Model(&Board{}).Find(dest, "user_id = ?", userId)
}

func FindBoardById(dest interface{}, id interface{}) *gorm.DB {
	return database.DB.Model(&Board{}).First(dest, "id = ?", id)
}

func DeleteBoard(id interface{}, userId interface{}) *gorm.DB {
	b := new(Board)
	database.DB.Model(&Board{}).First(b, "id = ? AND user_id = ?", id, userId)
	return database.DB.Delete(b)
}

func UpdateBoard(id interface{}, userId interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&Board{}).Where("id = ? AND user_id = ?", id, userId).Updates(data)
}
