package repos

import (
	entities "github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-eibrahimarisoy/domain/entities"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

// NewAuthorRepository returns a new AuthorRepository
func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}

}

// Migrations runs the database migrations
func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&entities.Author{})
}

// InsertSampleData inserts sample data into the database
func (a *AuthorRepository) InsertSampleData(author *entities.Author) entities.Author {
	result := a.db.Unscoped().Where("name = ?", author.Name).FirstOrCreate(author)

	if result.Error != nil {
		panic(result.Error) // TODO: handle error
	}
	return *author
}

// GetByID returns an author by id
func (a *AuthorRepository) GetByID(id int) (entities.Author, error) {
	var author entities.Author

	result := a.db.Where("id = ?", id).First(&author)
	if result.Error != nil {
		return entities.Author{}, result.Error
	}
	return author, nil
}

// FindByName returns an author by name
func (a *AuthorRepository) FindByName(name string) ([]entities.Author, error) {
	var authors []entities.Author

	result := a.db.Where("name ILIKE ?", "%"+name+"%").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

// GetAuthorsWithBooks returns author with books
func (a *AuthorRepository) GetByIDWithBooks(id int) (entities.Author, error) {
	var author entities.Author

	result := a.db.Preload("Books").Where("id = ?", id).First(&author)
	if result.Error != nil {
		return entities.Author{}, result.Error
	}
	return author, nil
}

// GetAuthorsWithBooks returns authors with books
func (a *AuthorRepository) GetAuthorsWithBooks() ([]entities.Author, error) {
	var authors []entities.Author
	result := a.db.Preload("Books").Find(&authors)
	if result.Error != nil {
		return []entities.Author{}, result.Error
	}
	return authors, nil
}

// **************EXTRA QUERIES************** //

// DeleteAuthorByID deletes author by id
func (a *AuthorRepository) DeleteAuthorByID(id int) error {
	result := a.db.Where("id = ?", id).Delete(&entities.Author{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateAuthorName updates author name
func (a *AuthorRepository) UpdateAuthorName(author *entities.Author) error {
	result := a.db.Model(&author).Where("id = ?", author.ID).Update("name", author.Name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
