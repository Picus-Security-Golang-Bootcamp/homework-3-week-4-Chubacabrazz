package author

import (
	book "github.com/Chubacabrazz/book-db/book_services/book"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Book_ID   string /* `gorm:"unique"` */
	Book_Name string
	Books     []book.Book `gorm:"foreignKey:Author;references:Book_Name"`
}
