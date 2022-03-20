package book

import (
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
