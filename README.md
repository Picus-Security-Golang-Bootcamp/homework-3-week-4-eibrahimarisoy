This repository contains Book Store application written by Go.
When application started it will read book/author data from csv file and store it in database.
- Reading from file is done using concurrency and worker pool.

Application includes many database queries made using the gorm tool.
- list, search, get, delete, buy
- other queries are made using the gorm tool.

## Clone the project
```
$ git clone https://github.com/Picus-Security-Golang-Bootcamp/homework-1-week-2-eibrahimarisoy.git
$ cd homework-1-week-2-eibrahimarisoy
$ go mod tidy
$ go run main.go
```
