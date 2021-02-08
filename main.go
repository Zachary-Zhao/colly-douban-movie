package main

import (
	"fmt"
	"regexp"
	"strings"

	"strconv"

	"colly/models"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func main() {
	url := "https://movie.douban.com/top250"
	c := colly.NewCollector()
	rd := regexp.MustCompile(`导演:(.+)`)
	ry := regexp.MustCompile(`\b\d{4}\b`)

	movies := []*models.Movie{}

	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		1, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 100}, // Use default queue storage
	)

	c.OnHTML("li div[class=info]", func(e *colly.HTMLElement) {
		movie := &models.Movie{}

		s := e.DOM.Find("div[class=hd] a span").Text()
		movie.Name = s
		fmt.Println(s)

		s = e.DOM.Find("div[class=bd] p").Text()
		sd := rd.FindStringSubmatch(s)
		sy := ry.FindStringSubmatch(s)
		s = strings.TrimSpace(sd[1])
		movie.Director = s
		fmt.Println(s)
		s = sy[0]
		movie.Year, _ = strconv.Atoi(s)
		fmt.Println(s)

		s = e.DOM.Find("div[class=bd] div span:nth-child(2)").Text()
		movie.Score, _ = strconv.ParseFloat(s, 64)
		fmt.Println(s)

		movies = append(movies, movie)
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

	models.BulkCreate(movies)
}
