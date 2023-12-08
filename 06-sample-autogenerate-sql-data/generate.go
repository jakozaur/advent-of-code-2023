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

var prompt string = `
I need some sample data for my database. I would provide you with a list of columns and their types.
I would need you to provide CSV (comma separated values) file with sample data for my database. Please escape all strings.
Please answer only with CSV file, not with the output of the statements or any additional explanations.
Please as sample data use realistic values, similar as the examples. Not just 'sample', 'example' or 'test' values.

`

func promptWithDefault(prompt, defaultValue string) string {
	fmt.Printf("%s [%s]:", prompt, defaultValue)
	var input string
	fmt.Scanln(&input)
	if input == "" {
		input = defaultValue
	}
	return input
}

type column struct {
	cid        int
	name       string
	_type      string
	notnull    int
	dflt_value sql.NullString
	pk         int
}

func generateDataForTable(openAiClient *openai.Client, db *sql.DB, table string) {
	fmt.Println("> Generating sample data for table:", table)
	columns := []column{}
	rows, err := db.Query("PRAGMA table_info(" + table + ");")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var column column
		err = rows.Scan(&column.cid, &column.name, &column._type, &column.notnull, &column.dflt_value, &column.pk)
		if err != nil {
			log.Fatal(err)
		}
		if column.name != "id" { // TODO: Avoid primary key
			columns = append(columns, column)
		}
	}

	// TODO: Do select TOP 5 for each column

	// Generate prompt
	columnList := ""
	for _, column := range columns {
		columnList += fmt.Sprintf(" - column with name '%s' and type '%s'\n", column.name, column._type)
	}

	// OpenAI prompt
	fmt.Println("> OpenAI prompt:")
	fmt.Println(prompt + columnList)
	resp, err := openAiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt + columnList,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println("> OpenAI response:")
	openAIResponse := resp.Choices[0].Message.Content
	fmt.Println(openAIResponse)

	linesUnparsed := strings.Split(openAIResponse, "\n")
	linesParsed := []string{}

	for _, lines := range linesUnparsed[1:] {
		linesParsed = append(linesParsed, strings.TrimSpace(lines))
	}

	sqlStatement := "INSERT INTO " + table + "(" + linesUnparsed[0] + ") VALUES (" + strings.Join(linesParsed, "), (") + ");"
	fmt.Println("> SQL statement:")
	fmt.Println(sqlStatement)

	_, err = db.Exec(sqlStatement)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "file:./sample.db"
	}

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tableNames := []string{}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		tableNames = append(tableNames, name)
	}
	fmt.Println("The program will generate sample data to following tables")
	selectedTableNames := promptWithDefault("Please enter table name", strings.Join(tableNames, ","))
	// TODO: validate table names
	selectedTable := strings.Split(selectedTableNames, ",")

	openAiToken := os.Getenv("OPENAI_SECRET_KEY")
	openAiClient := openai.NewClient(openAiToken)

	for _, table := range selectedTable {
		generateDataForTable(openAiClient, db, table)
	}
	fmt.Println("> Done")
}
