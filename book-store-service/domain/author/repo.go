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

func (a *AuthorRepository) InsertSampleData(author *Author) Author {
	result := a.db.FirstOrCreate(author)
	if result.Error != nil {
		panic(result.Error)
	}
	return *author
}
