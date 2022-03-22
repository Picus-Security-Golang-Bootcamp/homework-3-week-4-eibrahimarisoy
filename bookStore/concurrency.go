package bookStore

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/author"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/book"
)

func ReadBookWithWorkerPool(path string) []book.Book {

	const workerCount = 10
	jobs := make(chan []string, workerCount)
	results := make(chan book.Book, workerCount)

	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go toStruct(jobs, results, &wg, i)
	}

	go func() {
		f, _ := os.Open(path)
		defer f.Close()

		lines, _ := csv.NewReader(f).ReadAll()
		for _, line := range lines[1:] {
			jobs <- line
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var books []book.Book

	for v := range results {
		books = append(books, v)
	}
	return books
}

func toStruct(jobs <-chan []string, results chan<- book.Book, wg *sync.WaitGroup, i int) {
	fmt.Println("worker", i, "started")
	defer wg.Done()
	for line := range jobs {
		fmt.Println("worker", i, "working on")

		pages, _ := strconv.Atoi(line[1])
		stockCount, _ := strconv.Atoi(line[2])
		price, _ := strconv.ParseFloat(line[3], 64)
		IsDeleted, _ := strconv.ParseBool(line[6])
		authorID, _ := strconv.Atoi(line[7])

		results <- book.Book{
			Name:       line[0],
			Pages:      pages,
			StockCount: stockCount,
			Price:      price,
			StockCode:  line[4],
			ISBN:       line[5],
			IsDeleted:  IsDeleted,
			Author:     author.Author{ID: uint(authorID), Name: line[8]},
		}
	}
}
