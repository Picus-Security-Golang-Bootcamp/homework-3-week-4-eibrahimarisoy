package book

import (
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Migrations() {
	r.db.AutoMigrate(&Book{})
}

func (r *BookRepository) InsertSampleData(books []Book) {

	for _, book := range books {
		// qa := a.db.Where(author.Author{Name: book.Author.Name, ID: book.Author.ID}).
		// 	FirstOrCreate(&book.Author)

		r.db.Where(Book{Name: book.Name}).
			Attrs(Book{Name: book.Name, ID: book.ID}).
			FirstOrCreate(&book)

	}
}

func (r *BookRepository) GetAllBooksWithoutAuthorInformation() ([]Book, error) {
	var books []Book
	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
