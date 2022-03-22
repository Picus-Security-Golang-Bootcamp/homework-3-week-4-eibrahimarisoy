package csv_utils

// import (
// 	"encoding/csv"
// 	"os"
// 	"strconv"

// 	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/author"
// 	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/book"
// )

// func ReadCSV(filename string) ([]book.Book, error) {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()

// 	csvReader := csv.NewReader(f)
// 	records, err := csvReader.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var books []book.Book
// 	for _, line := range records[1:] {

// 		pages, err := strconv.Atoi(line[1])
// 		if err != nil {
// 			pages = 0
// 		}

// 		stockCount, err := strconv.Atoi(line[2])
// 		if err != nil {
// 			stockCount = 0
// 		}

// 		price, err := strconv.ParseFloat(line[3], 64)
// 		if err != nil {
// 			price = 0
// 		}
// 		IsDeleted, err := strconv.ParseBool(line[6])
// 		if err != nil {
// 			IsDeleted = false
// 		}
// 		authorID, err := strconv.Atoi(line[7])
// 		if err != nil {
// 			authorID = 0
// 		}
// 		books = append(books, book.Book{
// 			Name:       line[0],
// 			Pages:      pages,
// 			StockCount: stockCount,
// 			Price:      price,
// 			StockCode:  line[4],
// 			ISBN:       line[5],
// 			IsDeleted:  IsDeleted,
// 			Author:     author.Author{ID: uint(authorID), Name: line[8]},
// 		})

// 	}

// 	return books, nil
// }
