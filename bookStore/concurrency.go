package bookStore

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/entities"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/repos"
)

// ReadBookWithWorkerPool reads the CSV file and write the book data into the database
func ReadAndWriteBookWithWorkerPool(path string, bookRepo *repos.BookRepository, authorRepo *repos.AuthorRepository) {
	const workerCount = 3

	jobs := make(chan []string, workerCount)
	results := make(chan entities.Book, workerCount)

	wg := sync.WaitGroup{}

	// initialize workers and start them
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go toStruct(jobs, results, &wg, i)
	}

	// read the CSV file with go routine
	go func() {
		f, _ := os.Open(path)
		defer f.Close()

		lines, _ := csv.NewReader(f).ReadAll()
		for _, line := range lines[1:] {
			jobs <- line
		}
		close(jobs)
	}()

	// wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// insert the books into the database
	for v := range results {
		WriteSampleBookToDB(v, authorRepo, bookRepo)
	}
}

// toStruct converts the CSV line to a Book struct
func toStruct(jobs <-chan []string, results chan<- entities.Book, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	// fmt.Println("worker", i, "started")
	for line := range jobs {
		fmt.Println("worker", i, "working on", line)

		pages, _ := strconv.Atoi(line[1])
		stockCount, _ := strconv.Atoi(line[2])
		price, _ := strconv.ParseFloat(line[3], 64)
		authorID, _ := strconv.Atoi(line[6])

		results <- entities.Book{
			Name:       line[0],
			Pages:      pages,
			StockCount: stockCount,
			Price:      price,
			StockCode:  line[4],
			ISBN:       line[5],
			Author:     entities.Author{ID: uint(authorID), Name: line[7]},
		}
	}
}

// WriteSampleBookToDB inserts sample data into the database
func WriteSampleBookToDB(data entities.Book, a *repos.AuthorRepository, b *repos.BookRepository) {
	newAuthor := a.InsertSampleData(&data.Author)
	data.AuthorID = newAuthor.ID
	b.InsertSampleData(data)
}
