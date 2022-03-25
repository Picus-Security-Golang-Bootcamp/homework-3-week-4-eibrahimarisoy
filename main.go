package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/common/file"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/repos"
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
	author, _ = bs.AuthorRepo.GetAuthorWithBooks(1)
	fmt.Println(author.ToString())

	// get authors with books
	authors, _ = bs.AuthorRepo.GetAuthorsWithBooks()
	for _, author := range authors {
		fmt.Println(author.ToString())
	}

}
