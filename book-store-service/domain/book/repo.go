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

func (r *BookRepository) InsertSampleData(b Book) {
	r.db.Omit("Author").Where(Book{Name: b.Name}).
		FirstOrCreate(&b)
}

func (r *BookRepository) GetAllBooksWithoutAuthorInformation() ([]Book, error) {
	var books []Book
	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) GetAllBooksWithAuthorInformation() ([]Book, error) {
	var books []Book

	result := r.db.Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) SearchBookNameWithKeyword(keyword string) ([]Book, error) {
	var books []Book

	result := r.db.Preload("Author").Where("name ILIKE ?", "%"+keyword+"%").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) GetBookByIDWithAuthor(id int) (Book, error) {
	var book Book

	result := r.db.Preload("Author").Where("id = ?", id).First(&book)
	if result.Error != nil {
		return Book{}, result.Error
	}
	return book, nil
}
