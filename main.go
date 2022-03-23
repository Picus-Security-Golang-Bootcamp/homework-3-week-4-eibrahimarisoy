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

	if err := bookStore.NewBookStore().Run(args); err != nil {
		usageAndExit(err.Error())
	}

}

// usageAndExit prints the usage information and exits with the given error message.
func usageAndExit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	fmt.Fprintf(os.Stderr, "\n\n")
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")

	os.Exit(1)
}
