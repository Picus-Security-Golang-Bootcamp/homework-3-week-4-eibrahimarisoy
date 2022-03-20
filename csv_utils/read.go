package csv_utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book"
)

func ReadCSV(filename string) ([]book.Book, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var locations []book.Book
	for _, line := range records[1:] {

		id, err := strconv.Atoi(line[0])
		if err != nil {
			id = 0
		}

		pages, err := strconv.Atoi(line[2])
		if err != nil {
			pages = 0
		}

		stockCount, err := strconv.Atoi(line[3])
		if err != nil {
			stockCount = 0
		}

		price, err := strconv.ParseFloat(line[4], 64)
		if err != nil {
			price = 0
		}
		IsDeleted, err := strconv.ParseBool(line[5])
		if err != nil {
			IsDeleted = false
		}
		authorID, err := strconv.Atoi(line[6])
		if err != nil {
			authorID = 0
		}

		locations = append(locations, book.Book{
			ID:         id,
			Name:       line[1],
			Pages:      pages,
			StockCount: stockCount,
			Price:      price,
			StockCode:  line[5],
			ISBN:       line[6],
			IsDeleted:  IsDeleted,
			Author:     book.NewAuthor(authorID, line[9]),
		})
	}

	return locations, nil
}
