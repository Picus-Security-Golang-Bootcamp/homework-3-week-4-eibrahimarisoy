package main

import (
	"fmt"
	"log"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/author"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/book"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/csv_utils"

	"github.com/joho/godotenv"
)

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

	// Read CSV file
	books, err1 := csv_utils.ReadCSV("data.csv")
	if err1 != nil {
		panic(err)
	}

	// Repositories
	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migrations()

	bookRepo := book.NewBookRepository(db)
	bookRepo.Migrations()
	InsertSampleData(books, authorRepo, bookRepo)

	books, err = bookRepo.GetAllBooksWithoutAuthorInformation()
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Println(book.ToString())
	}
	fmt.Println("Done")
}

func InsertSampleData(books []book.Book, a *author.AuthorRepository, b *book.BookRepository) {

	for _, v := range books {
		a.InsertSampleData(v.Author)
		v.Author.ID = 0
		v.Author.Name = ""
		b.InsertSampleData(v)
	}
}
