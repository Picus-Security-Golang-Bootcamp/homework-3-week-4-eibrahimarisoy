package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/bookStore"
	"github.com/joho/godotenv"
)

// define usage information
var usage = `Usage: ./ [commands...] [parameters...]

Commands:
	-list
	-search <bookName>
	-get <bookID>
	-delete <bookID>
	-buy <bookID> <quantity>

Parameters:
	-keyword: string
	-bookID: int
	-quantity: int
`

func main() {
	// Set environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	args := os.Args[1:]

	bs, err := bookStore.NewBookStore()

	if err != nil {
		usageAndExit(err.Error())
	}

	if err := bs.Run(args); err != nil {
		usageAndExit(err.Error())
	}
	// books, err = bookRepo.GetAllBooksWithAuthorInformation()
	// if err != nil {
	// 	panic(err)
	// }

	// for _, book := range books {
	// 	fmt.Println(book.ToString())
	// }
	// fmt.Println("Done")
}

func usageAndExit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	fmt.Fprintf(os.Stderr, "\n\n")
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")

	os.Exit(1)
}
