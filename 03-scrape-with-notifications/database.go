package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

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

	newSubmissions := []Submission{}

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
