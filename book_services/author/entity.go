package author

import (
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Author_ID   string
	Author_Name string
}

func (c *Author) ToString() string {
	return fmt.Sprintf("Book_ID : %s, Book_Name : %s", c.Author_ID, c.Author_Name)
}
