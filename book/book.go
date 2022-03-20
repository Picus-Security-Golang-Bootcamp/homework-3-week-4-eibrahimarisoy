package book

import (
	"fmt"
)

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	IsDeleted  bool    `json:"is_deleted"`
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Pages      int     `json:"pages"`
	StockCount int     `json:"stockCount"`
	Price      float64 `json:"price"`
	StockCode  string  `json:"stockCode"`
	ISBN       string
	Author     `json:"author"`
}

// NewAuthor creates a new Author instance
func NewAuthor(id int, name string) Author {
	return Author{id, name}
}

// NewBook returns a new Book instance
func NewBook(v map[string]interface{}) *Book {
	return &Book{
		ID:         int(v["id"].(float64)),
		Name:       v["name"].(string),
		Pages:      int(v["pages"].(float64)),
		StockCount: int(v["stockCount"].(float64)),
		Price:      v["price"].(float64),
		StockCode:  v["stockCode"].(string),
		ISBN:       v["ISBN"].(string),
		IsDeleted:  v["isDeleted"].(bool),
		Author: NewAuthor(
			int(v["author"].(map[string]interface{})["id"].(float64)),
			v["author"].(map[string]interface{})["name"].(string),
		),
	}
}

// BookInfo prints the book information
func (b *Book) BookInfo() {
	fmt.Printf("ID : %v \n", b.ID)
	fmt.Printf("Name : %s \n", b.Name)
	fmt.Printf("Pages : %v \n", b.Pages)
	fmt.Printf("Stock Count : %d \n", b.StockCount)
	fmt.Printf("Price : %.2f \n", b.Price)
	fmt.Printf("Stock Code : %s \n", b.StockCode)
	fmt.Printf("ISBN : %s \n", b.ISBN)
	fmt.Printf("Is Deleted : %t \n", b.IsDeleted)
	fmt.Printf("Author ID : %v \n", b.Author.ID)
	fmt.Printf("Author Name : %s \n", b.Author.Name)
}

// setBookStock sets the stock count
func (b *Book) SetBookStock(stock int) {
	b.StockCount = stock
}

// setIsDeleted sets the stock count
func (b *Book) SetIsDeleted(stat bool) {
	b.IsDeleted = stat
}
