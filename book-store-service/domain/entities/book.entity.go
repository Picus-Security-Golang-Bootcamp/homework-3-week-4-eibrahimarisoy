package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name       string  `json:"name"`
	Pages      int     `json:"pages"`
	StockCount int     `json:"stock_count"`
	Price      float64 `json:"price"`
	StockCode  string  `json:"stock_code"`
	ISBN       string
	AuthorID   uint `gorm:"foreignKey:ID"`
	Author     Author
}

func (Book) TableName() string {
	return "book"
}

func (b *Book) ToString() string {
	if b.Author.ID != 0 {
		return fmt.Sprintf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s AuthorID: %v AuthorName: %s DeletedAt: %v", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.Author.ID, b.Author.Name, b.DeletedAt)
	}
	return fmt.Sprintf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s DeletedAt: %v", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.DeletedAt)
}
