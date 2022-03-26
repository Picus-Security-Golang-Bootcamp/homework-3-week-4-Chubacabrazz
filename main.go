package main

import (
	"log"

	"github.com/Chubacabrazz/book-db/book_services/author"
	book "github.com/Chubacabrazz/book-db/book_services/book"
	postgres "github.com/Chubacabrazz/book-db/db"
	"github.com/joho/godotenv"
)

func main() {
	//Set environment variables : database infos.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")
	// Repositories
	bookRepo := book.NewBookRepository(db)
	bookRepo.Migrations()
	/* bookRepo.InsertData()
	bookRepo.Buy(<quantity> , <ID>)
	bookRepo.FindByAuthor(<authorname>)
	bookRepo.SoftDeletebyID(<ID>)
	bookRepo.List() */

	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migration()
	/* authorRepo.InsertData()
	authorRepo.List()
	authorRepo.FindByWord(<name>) */

}
