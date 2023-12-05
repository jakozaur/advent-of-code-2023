package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Submission struct {
	id, time           string
	createdAt          string
	url, title, author string
	votes, comments    int
}

func ScrapeHackerNews() []Submission {
	// scraping logic...
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// slice of submissions
	submissions := []Submission{}

	c.OnHTML("#hnmain table tr", func(e *colly.HTMLElement) {
		if e.Attr("class") == "athing" {
			submission := Submission{
				id:    e.Attr("id"),
				url:   e.ChildAttr("td.title a", "href"),
				title: e.ChildText("td.title a"),
			}
			//fmt.Printf("%v\n", e.ChildText("td.title a"))
			submissions = append(submissions, submission)
		} else {
			scoreText := e.ChildText(".score")
			if scoreText != "" {
				time := e.ChildText("span.age a")
				commentsText := e.ChildText("td.subtext a:last-child")
				author := e.ChildText("td.subtext a.hnuser")
				createdAt := e.ChildAttr("td.subtext span.age", "title")

				score, scoreErr := strconv.Atoi(strings.Fields(scoreText)[0])
				comments, commentsErr := strconv.Atoi(strings.Fields(commentsText)[0])

				if scoreErr != nil {
					log.Println("Could not parse score: ", scoreErr)
				}

				if commentsErr != nil {
					log.Println("Could not parse comments: ", commentsErr)
				}

				submissions[len(submissions)-1].votes = score
				submissions[len(submissions)-1].time = time
				submissions[len(submissions)-1].comments = comments
				submissions[len(submissions)-1].author = author
				submissions[len(submissions)-1].createdAt = createdAt
			}
		}

	})

	// c.OnScraped(func(r *colly.Response) {
	// 	for _, submission := range submissions {
	// 		fmt.Printf("%+v\n", submission)
	// 	}
	// })

	c.Visit("https://news.ycombinator.com")

	return submissions
}
