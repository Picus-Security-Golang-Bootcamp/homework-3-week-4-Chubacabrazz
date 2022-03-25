package csv_utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	models "github.com/Chubacabrazz/book-db/file_services/models"
)

var BookList []models.Book

func ReadBooksWithWorkerPool(path string) error {
	const numJobs = 5
	jobs := make(chan []string, numJobs)
	results := make(chan models.Book, numJobs)
	wg := sync.WaitGroup{}

	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go toStruct(jobs, results, &wg)
	}

	go func() {
		fmt.Println("open file running...")
		f, _ := os.Open(path)
		defer f.Close()
		lines, _ := csv.NewReader(f).ReadAll()

		for _, line := range lines[1:] {

			jobs <- line
		}

		close(jobs)
	}()

	go func() {
		fmt.Println("wait")
		wg.Wait()
		close(results)
	}()

	BookList = make([]models.Book, 5)
	var i int = 0
	for v := range results {
		BookList[i] = v
		i++
	}
	return nil
}

//func toStruct reads given csv file by line and pushes them to result channel.
func toStruct(jobs <-chan []string, results chan<- models.Book, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println()
	for j := range jobs {
		books := models.Book{
			Book_ID:    j[0],
			Book_Name:  j[1],
			Book_Page:  j[2],
			Book_Stock: j[3],
			Book_Price: j[4],
			Book_Scode: j[5],
			Book_ISBN:  j[6],
			Author:     j[7],
		}

		results <- books
	}
}
