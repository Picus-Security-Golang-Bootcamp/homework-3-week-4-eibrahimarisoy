package bookStore

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/entities"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/repos"
)

type BookStore struct {
	BookRepo   *repos.BookRepository
	AuthorRepo *repos.AuthorRepository
}

// NewBookStore creates a new BookStore
func NewBookStore() BookStore {
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
	ReadAndWriteBookWithWorkerPool("data.csv", bookRepo, authorRepo)

	// initialize and return BookStore
	return BookStore{BookRepo: bookRepo, AuthorRepo: authorRepo}

}

// Run runs the bookStore given the command and the arguments
func (bs BookStore) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("No command provided")
	}

	switch args[0] {

	case "list":
		return bs.list()

	case "search":
		return bs.search(args)

	case "get":
		return bs.get(args)

	case "delete":
		return bs.delete(args)

	case "buy":
		return bs.buy(args)

	default:
		return fmt.Errorf("Invalid command")
	}
}

// PrintBooks prints the books given
func printBooks(books []entities.Book) {

	for _, v := range books {
		fmt.Println(v.ToString())
		fmt.Println("-", strings.Repeat("-", 50))
	}
}

// list all books including deleted
func (bs *BookStore) list() error {
	results, err := bs.BookRepo.GetBooksWithAuthor()
	if err != nil {
		return err
	}
	printBooks(results)
	return nil
}

// search checks if given keyword is in book name, and returns the books
func (bs *BookStore) search(args []string) error {
	// if the user has not provided <bookName>
	if len(args) < 2 {
		return fmt.Errorf("No book name provided")
	}

	results, err := bs.BookRepo.FindByName(strings.Join(args[1:], " "))
	if err != nil {
		return err
	}

	if len(results) == 0 {
		return fmt.Errorf("No book found")
	}

	printBooks(results)
	return nil
}

// get returns the book with given id
func (bs *BookStore) get(args []string) error {
	// if the user has not provided <bookID>
	if len(args) < 2 {
		return fmt.Errorf("No book id provided")
	}

	bookId, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	result, err := bs.BookRepo.GetByIDWithAuthor(bookId)
	if err != nil {
		return err
	}
	fmt.Println(result.ToString())
	return nil
}

// delete deletes the book with given id
func (bs *BookStore) delete(args []string) error {
	// if the user has not provided <bookID>
	if len(args) < 2 {
		return fmt.Errorf("No book id provided")
	}

	bookId, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	err = bs.BookRepo.DeleteBookByID(bookId)
	if err != nil {
		return err
	}

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Deleting book id:", bookId)
	fmt.Println(strings.Repeat("-", 50))
	return nil
}

// buy buys the book with given id and quantity and update the stock count
func (bs *BookStore) buy(args []string) error {
	// if the user has not provided <bookID> or <quantity>
	if len(args) < 3 {
		return fmt.Errorf("No book id or quantity provided")
	}

	bookId, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	quantity, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}

	if quantity <= 0 {
		return fmt.Errorf("Quantity must be greater than 0")
	}

	instance, err := bs.BookRepo.GetByIDWithAuthor(bookId)

	if err != nil {
		return err
	}

	if instance.DeletedAt.Valid {
		return fmt.Errorf("Book is not available")
	}

	if instance.StockCount < quantity {
		return fmt.Errorf("Not enough stock")
	}

	newInstance, err := bs.BookRepo.UpdateBookStockCount(&instance, instance.StockCount-quantity)
	if err != nil {
		return err
	}

	fmt.Println(newInstance.ToString())
	return nil
}
