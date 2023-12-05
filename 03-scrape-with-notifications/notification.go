package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type NotificationTrigger struct {
	email    string
	keywords []string
}

func parseNotificationTriggers() []NotificationTrigger {
	var triggers []NotificationTrigger
	triggersText := os.Getenv("TRIGGERS")
	fmt.Println("Parsng notification triggers", triggersText)

	for _, text := range strings.Split(triggersText, ";") {
		parts := strings.Split(text, ":")
		if len(parts) != 2 {
			log.Fatal("Invalid trigger: ", text, "Right format email@example.com=keyword1,keyword2")
		}
		email := parts[0]
		keywords := strings.Split(parts[1], ",")
		triggers = append(triggers, NotificationTrigger{email, keywords})
	}
	return triggers
}

func sendEmail(recepient string, subject string, body string) {
	// send email logic...
	fmt.Println("Sending email to", recepient, "with title", subject, "and body", body)

	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))
	sender := os.Getenv("MAILGUN_SENDER")
	message := mg.NewMessage(sender, subject, body, recepient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

func SendNotificatios(newSubmissions []Submission) {
	triggers := parseNotificationTriggers()

	fmt.Println("Checking whether to send notifications for", len(triggers), "triggers")
	for _, trigger := range triggers {
		newSubmissionsForTrigger := []Submission{}
		for _, submission := range newSubmissions {
			sumbissionMatchKeyword := false
			for _, keyword := range trigger.keywords {
				if strings.Contains(strings.ToLower(submission.title), strings.ToLower(keyword)) {
					sumbissionMatchKeyword = true
				}
			}
			if sumbissionMatchKeyword {
				newSubmissionsForTrigger = append(newSubmissionsForTrigger, submission)
			}
		}

		if len(newSubmissionsForTrigger) > 0 {
			fmt.Println("Sending email to", trigger.email, "with", len(newSubmissionsForTrigger), "new submissions")
			// format email body
			title := "New hacker news submissions: "
			body := "New hacker news submissions matching your keywords: " + strings.Join(trigger.keywords, ", ") + "\n\n"
			body += "https://news.ycombinator.com\n\n"

			for i, submission := range newSubmissionsForTrigger {
				if i > 0 {
					title += ", "
				}
				title += submission.title
				body += fmt.Sprintf("%s: https://news.ycombinator.com/item?id=%s\n\n", submission.title, submission.id)
			}

			body += "Thanks for using our service!\n\n"

			sendEmail(trigger.email, title, body)
		}
	}
}
