package repos

import (
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/book-store-service/domain/entities"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}

}

func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&entities.Author{})
}

func (a *AuthorRepository) InsertSampleData(author *entities.Author) entities.Author {
	result := a.db.FirstOrCreate(author)
	if result.Error != nil {
		panic(result.Error)
	}
	return *author
}

func (a *AuthorRepository) GetByID(id int) (entities.Author, error) {
	var author entities.Author
	result := a.db.Where("id = ?", id).First(&author)
	if result.Error != nil {
		return entities.Author{}, result.Error
	}
	return author, nil
}

func (a *AuthorRepository) FindByName(name string) ([]entities.Author, error) {
	var authors []entities.Author

	result := a.db.Where("name ILIKE ?", "%"+name+"%").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

func (a *AuthorRepository) GetAuthorWithBooks(id int) (entities.Author, error) {
	var author entities.Author
	result := a.db.Preload("Books").Where("id = ?", id).First(&author)
	if result.Error != nil {
		return entities.Author{}, result.Error
	}
	return author, nil
}

/*	var author Author
	var books []Book

	result := a.db.Where("id = ?", id).First(&author)
	if result.Error != nil {
		return Author{}, nil, result.Error
	}
	result = a.db.Model(&author).Related(&books)
	if result.Error != nil {
		return Author{}, nil, result.Error
	}
	return author, books, nil*/
