package file_services

import (
	"log"

	csv_utils "github.com/Chubacabrazz/book-db/file_services/csv_utils"
)

func File() {

	err := csv_utils.ReadBooksWithWorkerPool("books.csv")
	if err != nil {
		log.Fatal(err)
	}
}
