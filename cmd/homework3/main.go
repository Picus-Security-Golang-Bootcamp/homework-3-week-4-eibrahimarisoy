package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/domain/repos"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/pkg/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/pkg/common/file"
	"github.com/joho/godotenv"
)

type BookStore struct {
	BookRepo   *repos.BookRepository
	AuthorRepo *repos.AuthorRepository
}

func main() {
	// Set environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect postgres database
	db, err := db.NewPsqlDB()
	if err != nil {
		log.Fatal(err)
	}

	// Repositories
	authorRepo := repos.NewAuthorRepository(db)
	authorRepo.Migrations()

	bookRepo := repos.NewBookRepository(db)
	bookRepo.Migrations()

	// Read CSV file and insert data into database with worker pool
	file.ReadAndWriteBookWithWorkerPool(os.Getenv("FILE_PATH"), bookRepo, authorRepo)

	// initialize and return BookStore
	bs := BookStore{BookRepo: bookRepo, AuthorRepo: authorRepo}

	runQueries(bs)

}

// runExtraQuery runs extra queries for homework
func runQueries(bs BookStore) {
	fmt.Println("\n\nExtra Queries:")

	// list
	results, _ := bs.BookRepo.GetBooksWithAuthor()
	for _, book := range results {
		fmt.Println(book.ToString())
	}

	// search
	results, _ = bs.BookRepo.FindByName("keyword")
	for _, book := range results {
		fmt.Println(book.ToString())
	}

	// get
	result, _ := bs.BookRepo.GetByIDWithAuthor(5)
	fmt.Println(result.ToString())

	// delete
	_ = bs.BookRepo.DeleteBookByID(5)

	// buy
	result, _ = bs.BookRepo.UpdateBookStockCountByID(1, 5)
	fmt.Println(result.ToString())

	// get author by id
	author, _ := bs.AuthorRepo.GetByID(1)

	fmt.Println(author.ToString())

	// get book by name
	authors, _ := bs.AuthorRepo.FindByName("author")
	for _, author := range authors {
		fmt.Println(author.ToString())
	}

	// get author by id with books
	author, _ = bs.AuthorRepo.GetByIDWithBooks(1)
	fmt.Println(author.ToString())

	// get authors with books
	authors, _ = bs.AuthorRepo.GetAuthorsWithBooks()
	for _, author := range authors {
		fmt.Println(author.ToString())
	}

	// *******EXTRA QUERIES********

	// delete author by id
	_ = bs.AuthorRepo.DeleteAuthorByID(2)

	// update author
	author.Name = "Eibrahim"
	err := bs.AuthorRepo.UpdateAuthorName(&author)
	if err != nil {
		fmt.Println(err)
	}

	// Update book name
	book, _ := bs.BookRepo.GetByID(15)
	bs.BookRepo.UpdateBookName(book, "New Book Name")

	// Update book price
	book, _ = bs.BookRepo.GetByID(15)
	bs.BookRepo.UpdateBookPrice(book, 333.3333)

	// Filter books by price
	books, _ := bs.BookRepo.FilterBookByPriceRange(100.0, 200.0)
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// Get books with given ids
	books, _ = bs.BookRepo.GetBooksWithIDs([]int{1, 2, 3, 4, 5})
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// Filter books with created_at between dates
	books, _ = bs.BookRepo.FilterBookByCreatedAtRange("2020-01-01", "2020-01-02")
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// Search books name and stock code with given keyword
	books, _ = bs.BookRepo.SearchBookByNameOrStockCode("20")
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// GetAllBooksOrderByPriceAsc returns books ordered by price ascending
	books, _ = bs.BookRepo.GetAllBooksOrderByPriceAsc()
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// GetFirstTenBooks returns first ten books
	books, _ = bs.BookRepo.GetFirstTenBooks()
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// GetBooksCount returns number of books
	count, _ := bs.BookRepo.GetCount()
	fmt.Println(count)

	// GetTotalStockCount returns total stock count
	count, _ = bs.BookRepo.GetTotalStockValue()
	fmt.Println(count)

	// GetAvgPrice returns average price
	avgPrice, _ := bs.BookRepo.GetAvgPrice()
	fmt.Println(avgPrice)

}
