package repos

import (
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/domain/entities"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

// NewBookRepository returns a new BookRepository
func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

// Migrations runs the database migrations
func (r *BookRepository) Migrations() {
	r.db.AutoMigrate(&entities.Book{})
}

// InsertSampleData inserts sample data into the database
func (r *BookRepository) InsertSampleData(b entities.Book) {
	r.db.Unscoped().Omit("Author").Where(entities.Book{Name: b.Name}).
		FirstOrCreate(&b)
}

// GetAuthorWithoutAuthorInformation returns only books
func (r *BookRepository) GetAllBooksWithoutAuthorInformation() ([]entities.Book, error) {
	var books []entities.Book
	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// GetBooksWithAuthor returns books with author
func (r *BookRepository) GetBooksWithAuthor() ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Unscoped().Preload("Author").Order("id").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// FindByName returns books by name
func (r *BookRepository) FindByName(keyword string) ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Preload("Author").Where("name ILIKE ?", "%"+keyword+"%").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// GetByIDWithAuthor returns books by ID with author
func (r *BookRepository) GetByIDWithAuthor(id int) (entities.Book, error) {
	var book entities.Book

	result := r.db.Unscoped().Preload("Author").Where("id = ?", id).First(&book)
	if result.Error != nil {
		return entities.Book{}, result.Error
	}
	return book, nil
}

// DeleteBookByID deletes book by ID
func (r *BookRepository) DeleteBookByID(id int) error {
	result := r.db.Where("id = ?", id).Delete(&entities.Book{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateBookStockCountByID updates book stock count by ID
func (r *BookRepository) UpdateBookStockCountByID(id, newStockCount int) (entities.Book, error) {
	instance, _ := r.GetByIDWithAuthor(id)
	instance.StockCount = newStockCount
	r.db.Model(&instance).Update("stock_count", newStockCount)

	return instance, nil
}
