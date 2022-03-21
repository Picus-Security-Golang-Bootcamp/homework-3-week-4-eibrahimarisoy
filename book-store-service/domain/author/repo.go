package author

import (
	"fmt"

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

func (a *AuthorRepository) InsertSampleData(author Author) {
	fmt.Println("Inserting author: ", author.Name)
	fmt.Println("Inserting author: ", author.ID)
	a.db.Where(Author{Name: author.Name, ID: author.ID}).
		Attrs(Author{ID: author.ID}).
		FirstOrCreate(&author)
}
