package author

import (
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (Author) TableName() string {
	return "author"
}

func (a *Author) ToString() string {
	return fmt.Sprintf("ID: %v Name: %s", a.ID, a.Name)
}
