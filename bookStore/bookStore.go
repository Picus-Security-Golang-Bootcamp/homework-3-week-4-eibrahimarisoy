package bookStore

import (
	"fmt"
	"log"
	"strings"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/author"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/book"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/csv_utils"
)

type BookStore struct {
	Books      []*book.Book
	bookRepo   *book.BookRepository
	authorRepo *author.AuthorRepository
}

// NewBookStore
func NewBookStore() (BookStore, error) {
	// connect postgres database
	db, err := db.NewPsqlDB()
	if err != nil {
		log.Fatal(err)
	}

	// Read CSV file
	books, err := csv_utils.ReadCSV("data.csv")
	if err != nil {
		panic(err)
	}

	// Repositories
	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migrations()

	bookRepo := book.NewBookRepository(db)
	bookRepo.Migrations()

	InsertSampleData(books, authorRepo, bookRepo)

	var newBookStore = BookStore{bookRepo: bookRepo, authorRepo: authorRepo}

	return newBookStore, nil
}

func InsertSampleData(books []book.Book, a *author.AuthorRepository, b *book.BookRepository) {

	for _, v := range books {
		newAuthor := a.InsertSampleData(&v.Author)
		v.AuthorID = newAuthor.ID
		b.InsertSampleData(v)
	}
}

// Run runs the bookStore given the command and the arguments
func (bs BookStore) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("No command provided")
	}

	switch args[0] {

	case "list":
		results, err := bs.bookRepo.GetAllBooksWithAuthorInformation()
		if err != nil {
			return err
		}
		PrintBooks(results)

	case "search":
		// if the user has not provided <bookName>
		if len(args) < 2 {
			return fmt.Errorf("No book name provided")
		}

		results, err := bs.bookRepo.SearchBookNameWithKeyword(strings.Join(args[1:], " "))
		if err != nil {
			return err
		}

		if len(results) == 0 {
			return fmt.Errorf("No book found")
		}

		PrintBooks(results)

	// case "get":
	// 	// if the user has not provided <bookID>
	// 	if len(args) < 2 {
	// 		return fmt.Errorf("No book id provided")
	// 	}

	// 	bookId, err := strconv.Atoi(args[1])
	// 	if err != nil {
	// 		return err
	// 	}

	// 	index, err := bs.Get(bookId)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	bs.Books[index].BookInfo()

	// case "delete":
	// 	// if the user has not provided <bookID>
	// 	if len(args) < 2 {
	// 		return fmt.Errorf("No book id provided")
	// 	}

	// 	bookId, err := strconv.Atoi(args[1])
	// 	if err != nil {
	// 		return err
	// 	}

	// 	index, err := bs.Get(bookId)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Println(strings.Repeat("-", 50))
	// 	fmt.Println("Deleting book:")
	// 	bs.Books[index].BookInfo()
	// 	fmt.Println(strings.Repeat("-", 50))

	// 	bs.Delete(index)
	// 	bs.List()

	// case "buy":
	// 	// if the user has not provided <bookID> or <quantity>
	// 	if len(args) < 3 {
	// 		return fmt.Errorf("No book id or quantity provided")
	// 	}

	// 	bookId, err := strconv.Atoi(args[1])
	// 	if err != nil {
	// 		return err
	// 	}

	// 	quantity, err := strconv.Atoi(args[2])
	// 	if err != nil {
	// 		return err
	// 	}

	// 	index, err := bs.Get(bookId)

	// 	if err != nil {
	// 		return err
	// 	}

	// 	instance := bs.Books[index]
	// 	if err := bs.Buy(instance, quantity); err != nil {
	// 		return err
	// 	}

	// 	instance.BookInfo()

	default:
		return fmt.Errorf("Invalid command")
	}
	return nil
}

// List prints all the books in bookStore
func PrintBooks(books []book.Book) {

	for _, v := range books {
		fmt.Println(v.ToString())
		fmt.Println("-", strings.Repeat("-", 50))
	}
}