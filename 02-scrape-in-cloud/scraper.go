package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
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
	var submissions []Submission

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

func UpdateDatabase(submissions []Submission) []Submission {
	// update database logic...
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "file:./hackernews.db"
	}

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS submissions (
		id TEXT PRIMARY KEY,
		title TEXT,
		url TEXT,
		author TEXT,
		time TEXT,
		createdAt TEXT,
		votes INTEGER,
		comments INTEGER)`)
	if err != nil {
		log.Fatal(err)
	}

	var newSubmissions []Submission

	for _, submission := range submissions {
		// fmt.Println("submission.id", submission.id)
		var id string
		err := db.QueryRow(`SELECT id FROM submissions WHERE id = ?`, submission.id).Scan(&id)

		if err == sql.ErrNoRows {
			// Add news
			// fmt.Println("Adding new submission: ", submission)
			_, err = db.Exec(`INSERT INTO submissions (
				id, title, url, author, time, createdAt, votes, comments)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
				submission.id, submission.title, submission.url,
				submission.author, submission.time, submission.createdAt,
				submission.votes, submission.comments)
			if err != nil {
				log.Fatal(err)
			}

			newSubmissions = append(newSubmissions, submission)
		} else {
			// Update news
			// fmt.Println("Updating submission: ", submission)
			_, err = db.Exec(`UPDATE submissions
				SET title = ?, url = ?, author = ?, time = ?, createdAt = ?, votes = ?, comments = ?
				WHERE id = ?`,
				submission.title, submission.url,
				submission.author, submission.time, submission.createdAt,
				submission.votes, submission.comments, submission.id)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return newSubmissions
}

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintln("event", event)

	submissions := ScrapeHackerNews()
	message += fmt.Sprintln("submissions", submissions)

	newSubmissions := UpdateDatabase(submissions)
	message += fmt.Sprintln("newSubmissions", newSubmissions)

	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
