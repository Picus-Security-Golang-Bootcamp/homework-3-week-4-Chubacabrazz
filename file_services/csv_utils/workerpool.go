package csv_utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	models "github.com/Chubacabrazz/book-db/file_services/models"
)

var Booklar []models.Book

func ReadBooksWithWorkerPool(path string) error {
	const numJobs = 5
	jobs := make(chan []string, numJobs)
	results := make(chan models.Book, numJobs)
	wg := sync.WaitGroup{}

	for w := 1; w <= 5; w++ {
		fmt.Println("worker starting", w)
		wg.Add(1)
		go toStruct(jobs, results, &wg)
	}

	go func() {
		fmt.Println("open file running...")
		f, _ := os.Open(path)
		defer f.Close()
		lines, _ := csv.NewReader(f).ReadAll()

		for _, line := range lines[1:] {

			fmt.Println("line", line[0])
			jobs <- line
		}

		close(jobs)
	}()

	go func() {
		fmt.Println("wait")
		wg.Wait()
		close(results)
	}()

	Booklar = make([]models.Book, 5)
	var i int = 0
	for v := range results {
		Booklar[i] = v
		i++
	}
	return nil
}

func toStruct(jobs <-chan []string, results chan<- models.Book, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println()
	for j := range jobs {
		location := models.Book{
			Book_ID:    j[0],
			Book_Name:  j[1],
			Book_Page:  j[2],
			Book_Stock: j[3],
			Book_Price: j[4],
			Book_Scode: j[5],
			Book_ISBN:  j[6],
			Author:     j[7],
		}
		//process...

		results <- location
	}
}
