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

// ToString returns a string representation of the book
func (b *Book) ToString() string {
	if b.Author.ID != 0 {
		return fmt.Sprintf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s AuthorID: %v AuthorName: %s DeletedAt: %v", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.Author.ID, b.Author.Name, b.DeletedAt)
	}
	return fmt.Sprintf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s DeletedAt: %v", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.DeletedAt)
}

// PrintBooks prints books
func (bs BookSlice) PrintBooks() {
	for _, b := range bs {
		fmt.Printf("ID: %v Name: %s Pages: %v StockCount: %v Price: %v StockCode: %s ISBN: %s DeletedAt: %v", b.ID, b.Name, b.Pages, b.StockCount, b.Price, b.StockCode, b.ISBN, b.DeletedAt)
	}
}

// BeforeCreate is a callback that gets called before creating
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println(b.ID, "is created")
	return
}

// AfterDelete is a callback that gets called after deleting
func (b *Book) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println(b.ID, "is deleted")
	return
}
