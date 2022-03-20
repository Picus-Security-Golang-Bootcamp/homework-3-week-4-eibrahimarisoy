package author

import (
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}

}

func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&Author{})
}

func (a *AuthorRepository) InsertSampleData(authors []Author) {

	for _, author := range authors {
		a.db.Where(Author{Name: author.Name}).
			Attrs(Author{Name: author.Name, ID: author.ID}).
			FirstOrCreate(&author)
	}
}
