package book

import (
	"fmt"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/author"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Pages      string `json:"pages"`
	StockCount int    `json:"stock_count"`
	Price      string `json:"price"`
	StockCode  string `json:"stock_code"`
	ISBN       string
	IsDeleted  bool          `json:"isDeleted"`
	Author     author.Author `gorm:"foreignkey:AuthorID;references:id"`
}

func (Book) TableName() string {
	return "authors"
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID: %v Name: %s Pages: %s StockCount: %d Price: %s StockCode: %s ISBN: %s IsDeleted: %t AuthorID: %v AuthorName: %s", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.IsDeleted, b.Author.ID, b.Author.Name)
}
