package main

import (
	"fmt"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/csv_utils"
)

func main() {

	books, err := csv_utils.ReadCSV("data.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Books:", books)
}
