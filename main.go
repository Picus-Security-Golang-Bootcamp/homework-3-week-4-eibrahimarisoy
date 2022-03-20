package main

import (
	"fmt"
	"log"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/csv_utils"
	"github.com/joho/godotenv"
)

func main() {
	// Set environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect postgres database
	_, err := db.NewPsqlDB()
	if err != nil {
		log.Fatal(err)
	}

	// Read CSV file
	books, err := csv_utils.ReadCSV("data.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Books:", books)

}
