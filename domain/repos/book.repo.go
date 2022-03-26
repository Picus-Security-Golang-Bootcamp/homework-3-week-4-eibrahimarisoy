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
	r.db.Unscoped().Omit("Author").Where(entities.Book{Name: b.Name, StockCode: b.StockCode}).
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

// GetByIDWithAuthor returns books by ID
func (r *BookRepository) GetByID(id int) (entities.Book, error) {
	var book entities.Book

	result := r.db.Unscoped().Where("id = ?", id).First(&book)
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
	instance.StockCount = uint(newStockCount)
	r.db.Model(&instance).Update("stock_count", newStockCount)

	return instance, nil
}

// **************EXTRA QUERIES************** //

// UpdateBookName updates book name
func (r *BookRepository) UpdateBookName(book entities.Book, newName string) (entities.Book, error) {
	book.Name = newName
	r.db.Model(&book).Update("name", newName)

	return book, nil
}

// UpdateBookPrice updates book price
func (r *BookRepository) UpdateBookPrice(book entities.Book, newPrice float64) (entities.Book, error) {
	book.Price = newPrice
	r.db.Model(&book).Update("price", newPrice)

	return book, nil
}

// FilterBookByPriceRange filters book by price range
func (r *BookRepository) FilterBookByPriceRange(min, max float64) ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Unscoped().Where("price BETWEEN ? AND ?", min, max).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// GetBooksWithIDs returns books by IDs
func (r *BookRepository) GetBooksWithIDs(ids []int) ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Where("id IN ?", ids).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// FilterBookByCreatedAtRange filters book by created at range
func (r *BookRepository) FilterBookByCreatedAtRange(min, max string) ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Unscoped().Where("created_at BETWEEN ? AND ?", min, max).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// SearchBookByNameAndStockCode searches book by name and stock code
func (r *BookRepository) SearchBookByNameOrStockCode(keyword string) ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Unscoped().Where("name ILIKE ? ", "%"+keyword+"%").Or("stock_code ILIKE ?", "%"+keyword+"%").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// GetAllBooksOrderByPriceAsc returns all books order by price asc
func (r *BookRepository) GetAllBooksOrderByPriceAsc() ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Unscoped().Order("price asc").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// GetFirstTenBooks returns first ten books
func (r *BookRepository) GetFirstTenBooks() ([]entities.Book, error) {
	var books []entities.Book

	result := r.db.Unscoped().Limit(10).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// GetBooksCount returns books count
func (r *BookRepository) GetCount() (int64, error) {
	var count int64

	result := r.db.Model(&entities.Book{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// GetTotalStockValue returns total stock value
func (r *BookRepository) GetTotalStockValue() (int64, error) {
	var count int64

	err := r.db.Model(&entities.Book{}).Select("sum(stock_count)").Row().Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetAvgPrice returns average price
func (r *BookRepository) GetAvgPrice() (float64, error) {
	var avgPrice float64

	err := r.db.Model(&entities.Book{}).Select("avg(price)").Row().Scan(&avgPrice)
	if err != nil {
		return 0, err
	}
	return avgPrice, nil
}
