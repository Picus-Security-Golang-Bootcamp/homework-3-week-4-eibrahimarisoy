package repos

import (
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/entities"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Migrations() {
	r.db.AutoMigrate(&entities.Book{})
}

func (r *BookRepository) InsertSampleData(b entities.Book) {
	r.db.Unscoped().Omit("Author").Where(entities.Book{Name: b.Name}).
		FirstOrCreate(&b)
}

func (r *BookRepository) GetAllBooksWithoutAuthorInformation() ([]entities.Book, error) {
	var books []entities.Book
	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) GetBooksWithAuthor() ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Unscoped().Preload("Author").Order("id").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) FindByName(keyword string) ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Preload("Author").Where("name ILIKE ?", "%"+keyword+"%").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) GetByIDWithAuthor(id int) (entities.Book, error) {
	var book entities.Book

	result := r.db.Unscoped().Preload("Author").Where("id = ?", id).First(&book)
	if result.Error != nil {
		return entities.Book{}, result.Error
	}
	return book, nil
}

func (r *BookRepository) DeleteBookByID(id int) error {
	result := r.db.Where("id = ?", id).Delete(&entities.Book{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *BookRepository) UpdateBookStockCount(b *entities.Book, newStockCount int) (*entities.Book, error) {
	tx := r.db.Model(&b).Update("stock_count", newStockCount)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return b, nil
}
