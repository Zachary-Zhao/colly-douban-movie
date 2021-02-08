package main

import (
	"fmt"
	"io/ioutil"

	_ "colly/models"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func main() {
	url := "https://movie.douban.com/top250"
	c := colly.NewCollector()
	var contents []byte

	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		1, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 1000}, // Use default queue storage
	)

	c.OnHTML("li div[class=info] a", func(e *colly.HTMLElement) {
		// content := []byte("")
		e.ForEach("span", func(i int, ee *colly.HTMLElement) {
			fmt.Print(ee.Text)
			contents = append(contents, []byte(ee.Text)...)
		})
		fmt.Println("\n")
		content := []byte("\n")
		contents = append(contents, content...)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// c.Visit("")
	for i := 0; i < 10; i++ {
		// Add URLs to the queue
		q.AddURL(fmt.Sprintf("%s?start=%d&filter=", url, i*25))
	}
	// Consume URLs
	q.Run(c)

	ioutil.WriteFile("豆瓣电影top250.txt", contents, 0644)
	// models.BulkCreate([]*models.Movie{&models.Movie{Name: "电影1"}, &models.Movie{Name: "电影2"}})
}
