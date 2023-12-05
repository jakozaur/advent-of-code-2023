package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

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

	// Log message
	fmt.Println(message)

	SendNotificatios(newSubmissions)

	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
	//HandleRequest(context.Background(), &MyEvent{Name: "test"})
}
