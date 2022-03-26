package book

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Chubacabrazz/book-db/file_services/csv_utils"
	"gorm.io/gorm"
)

var csvfile string = "books.csv"

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) List() {
	var books []Book
	b.db.Find(&books)

	for _, thebook := range books {
		fmt.Println(thebook.ToString())
	}
}

//Func SoftDeletebyID applies a soft delete to a book
func (b *BookRepository) SoftDeletebyID(id int) error {
	var book Book
	result := b.db.First(&book, id)
	if result.Error != nil {
		return result.Error
	} else {
		fmt.Println("Success, book soft-deleted:", id)
	}
	result = b.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Func Buy does purchase for a book with given ID and quantity
func (b *BookRepository) Buy(quantity, id int) error {
	var book Book
	stock, _ := strconv.Atoi(book.Book_Stock)
	result := b.db.First(&book, id)
	if result.Error != nil {
		return result.Error
	} else if stock < quantity {
		return fmt.Errorf("we don't have that much. we have: %d", stock)
	} else {
		fmt.Println("Shopping successfull.")
	}

	result = b.db.Model(&book).Where("id = ? AND num_of_books_in_stock >= ?", id, quantity).
		Update("num_of_books_in_stock", gorm.Expr("num_of_books_in_stock - ?", quantity))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Func FindByAuthor finds books with Author Name
func (c *BookRepository) FindByAuthor(Author string) []Book {
	var books []Book
	c.db.Where(`"Author" = ?`, Author).Order("Id desc,name").Find(&books)
	return books
}

//Func SearchWord lists the books with the given word. (case insensitive)
func (b *BookRepository) SearchWord(name string) {
	var books []Book
	b.db.Where("name ILIKE ? ", "%"+name+"%").Find(&books)

	for _, thebook := range books {
		fmt.Println(thebook.ToString())
	}
}

//Func GetByID prints book details of given ID
func (c *BookRepository) GetByID(id int) (*Book, error) {
	var book Book
	result := c.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &book, nil
}

//Func GetByID hard deletes book details of given ID
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

// Func InsertData starts a concurrent csv reading operation and write them to database.
func (c *BookRepository) InsertData() {
	csv_utils.ReadBooksWithWorkerPool(csvfile)
	for _, book := range csv_utils.BookList {
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
