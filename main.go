package main

import (
	"log"

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
	bookRepo.InsertSampleData()
	bookRepo.FindAll()

	/* authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migration()
	authorRepo.InsertSampleData()

	fmt.Println(authorRepo.GetAllAuthorsWithBookInformation()) */

}
