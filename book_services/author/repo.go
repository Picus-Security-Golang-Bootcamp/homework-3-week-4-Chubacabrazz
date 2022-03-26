package author

import (
	"fmt"

	"github.com/Chubacabrazz/book-db/file_services/csv_utils"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

//FindByWord lists the authors with the given word case-insensitively
func (a *AuthorRepository) FindByWord(name string) {
	var authors []Author
	a.db.Where("name ILIKE ? ", "%"+name+"%").Find(&authors)

	for _, author := range authors {
		fmt.Println(author.ToString())
	}
}

//Func List prints all authors from db
func (a *AuthorRepository) List() {
	var authors []Author
	a.db.Find(&authors)

	for _, author := range authors {
		fmt.Println(author.ToString())
	}
}

//DeleteByID does a soft delete to an author with given ID
func (a *AuthorRepository) DeleteById(id int) error {
	var author Author
	result := a.db.First(&author, id)
	if result.Error != nil {
		return result.Error
	} else {
		fmt.Println("Success, soft deleted:", id)
	}
	result = a.db.Delete(&Author{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *AuthorRepository) Migration() {
	c.db.AutoMigrate(&Author{})
}

// Func InsertData starts a concurrent csv reading operation and write them to database.
func (c *AuthorRepository) InsertData() {
	for _, book := range csv_utils.BookList {
		c.db.FirstOrCreate(Author{
			Author_ID:   book.Author_ID,
			Author_Name: book.Author})
	}

}
