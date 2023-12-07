package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"

	openai "github.com/sashabaranov/go-openai"
)

var createTableSql string = `
CREATE TABLE IF NOT EXISTS orders (
	id TEXT PRIMARY KEY,
	name TEXT,
	date TEXT,
	price INTEGER,
	url TEXT,
	status TEXT)`

var prompt string = `
Can you create a few INSERT SQL statements with sample data in Sqlite3 for the table I would paste at end of this message?
Please answer only with SQL statements, not with the output of the statements or any additional explanations.
Please as sample data use realistic English values, not just 'sample', 'example' or 'test' values.
`

func main() {
	openAiToken := os.Getenv("OPENAI_SECRET_KEY")
	client := openai.NewClient(openAiToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt + createTableSql,
				},
			},
		},
	)

	fmt.Println("> OpenAI prompt:")
	fmt.Println(prompt + createTableSql)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println("> OpenAI response:")
	fmt.Println(resp.Choices[0].Message.Content)
	statements := strings.Split(resp.Choices[0].Message.Content, "\n")

	// update database logic...
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "file:./sample.db"
	}

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(createTableSql)
	if err != nil {
		log.Fatal(err)
	}

	for _, statement := range statements {
		if statement == "" {
			continue
		}
		fmt.Println("> Executing SQL statement:", statement)
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Error executing SQL statement:", statement, "error", err)
		}
	}

	fmt.Println("> Done")
}
