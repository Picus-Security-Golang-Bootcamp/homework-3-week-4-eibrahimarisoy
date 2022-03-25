package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type BookSlice []Book

type Book struct {
	gorm.Model
	Name       string  `json:"name"`
	Pages      uint    `json:"pages"`
	StockCount uint    `json:"stock_count"`
	Price      float64 `json:"price"`
	StockCode  string  `json:"stock_code" gorm:"unique"`
	ISBN       string  `gorm:"unique"`
	AuthorID   uint
	Author     Author `gorm:"OnDelete:SET NULL"`
}

func (Book) TableName() string {
	return "books"
}

func (b *Book) ToString() string {
	if b.Author.ID != 0 {
		return fmt.Sprintf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s AuthorID: %v AuthorName: %s DeletedAt: %v", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.Author.ID, b.Author.Name, b.DeletedAt)
	}
	return fmt.Sprintf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s DeletedAt: %v", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.DeletedAt)
}

func (bs BookSlice) PrintBooks() {
	for _, b := range bs {
		fmt.Printf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s DeletedAt: %v", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.DeletedAt)
	}
}
