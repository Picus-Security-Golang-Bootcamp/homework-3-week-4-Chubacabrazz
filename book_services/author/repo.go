package author

import "gorm.io/gorm"

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (c *AuthorRepository) GetAllAuthorsWithBookInformation() ([]Author, error) {
	var authors []Author
	result := c.db.Preload("Books").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

func (c *AuthorRepository) GetAuthorWithName(name string) (*Author, error) {
	var author *Author
	result := c.db.
		Where(Author{Book_ID: name}).
		Attrs(Author{Book_Name: "NULL", Book_ID: "NULL"}).
		FirstOrInit(&author) // Eğer sorgu sonucunda veri bulunursa Attrs kısmında yazılanlar ignore edilir.

	if result.Error != nil {
		return nil, result.Error
	}

	return author, nil
}

func (c *AuthorRepository) GetAllAuthorsWithoutBookInformation() ([]Author, error) {
	var authors []Author
	result := c.db.Find(&authors)

	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}

func (c *AuthorRepository) Migration() {
	c.db.AutoMigrate(&Author{})
}

func (c *AuthorRepository) InsertSampleData() {
	books := []Author{
		{Book_ID: "Türkiye", Book_Name: "TR"},
		{Book_ID: "Amerika", Book_Name: "US"},
	}

	for _, book := range books {
		c.db.Create(&book)
	}
}
