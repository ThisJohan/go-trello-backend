package dal

import (
	"github.com/ThisJohan/go-trello-clone/config/database"
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	ListId      uint `gorm:"foreignKey:list_id"`
	AssignedTo  uint
	DueDate     string
}

func CreateCard(c *Card) *gorm.DB {
	return database.DB.Create(c)
}

func FindCardsByList(dest interface{}, listId interface{}) *gorm.DB {
	return database.DB.Model(&Card{}).Find(dest, "list_id = ?", listId)
}

func FindCardById(dest interface{}, id interface{}) *gorm.DB {
	return database.DB.Model(&Card{}).First(dest, "id = ?", id)
}

func DeleteCard(id interface{}) *gorm.DB {
	return database.DB.Delete(&Card{}, id)
}

func UpdateCard(id interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&Card{}).Where("id = ?", id).Updates(data)
}
