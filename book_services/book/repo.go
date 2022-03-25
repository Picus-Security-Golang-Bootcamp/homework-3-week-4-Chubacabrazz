package book

import (
	"errors"

	"github.com/Chubacabrazz/book-db/file_services/csv_utils"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (c *BookRepository) FindAll() []Book {
	var books []Book
	c.db.Find(&books)

	return books
}

func (c *BookRepository) FindByAuthor(Author string) []Book {
	var books []Book
	c.db.Where(`"Author" = ?`, Author).Order("Id desc,name").Find(&books)
	return books
}

func (c *BookRepository) FindByAuthorCodeOrBookCode(code string) []Book {
	var books []Book
	c.db.Where(`"Author = ?"`, code).Or("code = ?", code).Find(&books)

	return books
}

func (c *BookRepository) FindByName(name string) []Book {
	var books []Book
	c.db.Where("name LIKE ? ", "%"+name+"%").Find(&books)

	return books
}

func (c *BookRepository) FindByNameWithRawSQL(name string) []Book {
	var books []Book
	c.db.Raw("SELECT * FROM Book WHERE name LIKE ?", "%"+name+"%").Scan(&books)

	return books
}

func (c *BookRepository) GetByID(id int) (*Book, error) {
	var book Book
	result := c.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &book, nil
}

func (c *BookRepository) Create(book Book) error {
	result := c.db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *BookRepository) Update(book Book) error {
	result := c.db.Save(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *BookRepository) Delete(book Book) error {
	result := c.db.Delete(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *BookRepository) DeleteById(id int) error {
	result := c.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *BookRepository) Migrations() {
	c.db.AutoMigrate(&Book{})
}

func (c *BookRepository) InsertSampleData() {
	csv_utils.ReadBooksWithWorkerPool("books.csv")
	for _, book := range csv_utils.Booklar {
		/* c.db.Where(Book{Book_Name: book.Book_Name}).
		Attrs(Book{
			Book_Name: book.Book_Name,
			Book_ID: book.Book_ID,
			Book_Price: book.Book_Price,
			Book_Page: book.Book_Page,
			Book_Stock: book.Book_Stock,
			Book_Scode: book.Book_Scode,
			Book_ISBN: book.Book_ISBN,}).
		FirstOrCreate(&book) */
		c.db.Create(&Book{
			Book_ID:    book.Book_ID,
			Book_Name:  book.Book_Name,
			Book_Price: book.Book_Price,
			Book_Page:  book.Book_Page,
			Book_Stock: book.Book_Stock,
			Book_Scode: book.Book_Scode,
			Book_ISBN:  book.Book_ISBN,
			Author:     book.Author})
	}
}
