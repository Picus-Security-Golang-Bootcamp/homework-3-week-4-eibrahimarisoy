package book

import (
	"fmt"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/author"
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
	IsDeleted  bool `json:"isDeleted"`
	AuthorID   uint `gorm:"foreignKey:ID"`
	Author     author.Author
}

func (Book) TableName() string {
	return "book"
}

func (b *Book) ToString() string {
	if b.Author.ID != 0 {
		return fmt.Sprintf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s IsDeleted: %t AuthorID: %v AuthorName: %s", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.IsDeleted, b.Author.ID, b.Author.Name)
	}
	return fmt.Sprintf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s IsDeleted: %t", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.IsDeleted)
}
