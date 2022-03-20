package book

import (
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Migrations() {
	r.db.AutoMigrate(&Book{})
}

func (r *BookRepository) InsertSampleData(books []Book) {

	for _, book := range books {
		r.db.Where(Book{Name: book.Name}).
			Attrs(Book{Name: book.Name, ID: book.ID}).
			FirstOrCreate(&book)
	}
}
