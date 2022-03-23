package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID    uint   `gorm:"primarykey"`
	Name  string `json:"name" gorm:"unique"`
	Books []Book
}

func (Author) TableName() string {
	return "author"
}

func (a *Author) ToString() string {
	if a.Books != nil {
		return fmt.Sprintf("ID: %v Name: %s Books: %v", a.ID, a.Name, a.Books)
	}
	return fmt.Sprintf("ID: %v Name: %s", a.ID, a.Name)
}
