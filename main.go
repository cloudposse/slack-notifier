package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	action = flag.String("action", os.Getenv("GITHUB_ACTION"), "Action to perform: 'update_state' or 'update_branch_protection'")
	token  = flag.String("token", os.Getenv("GITHUB_TOKEN"), "Github access token")
	owner  = flag.String("owner", os.Getenv("GITHUB_OWNER"), "Github repository owner")
	repo   = flag.String("repo", os.Getenv("GITHUB_REPO"), "Github repository name")
)

func main() {
	flag.Parse()

	if *action == "" {
		flag.PrintDefaults()
		log.Fatal("-action or GITHUB_ACTION required")
	}
	if *action != "update_state" && *action != "update_branch_protection" {
		flag.PrintDefaults()
		log.Fatal("-action or GITHUB_ACTION must be 'update_state' or 'update_branch_protection'")
	}
	if *token == "" {
		flag.PrintDefaults()
		log.Fatal("-token or GITHUB_TOKEN required")
	}
	if *owner == "" {
		flag.PrintDefaults()
		log.Fatal("-owner or GITHUB_OWNER required")
	}
	if *repo == "" {
		flag.PrintDefaults()
		log.Fatal("-repo or GITHUB_REPO required")
	}

	notifier := NewSlackNotifier("#demo", ":octocat:", "", "")

	payload := notifier.NewPayload()
	payload.Mrkdwn = true

	statusField := Field{
		Title: "Status",
		Value: "Good",
		Short: true,
	}

	dateString := time.Now().Format("2006-01-02 15:04")
	dateField := Field{
		Title: "Date",
		Value: dateString,
		Short: true,
	}

	attachment := Attachment{
		Fallback: "GitHub Status: Good - https://status.github.com",
		Text:     "<https://status.github.com/|GitHub Status> : *Good*",
		Color:    "good",
		Fields:   []Field{statusField, dateField},
		MrkdwnIn: []string{"text"},
	}

	payload.Attachments = []Attachment{attachment}

	notifier.PostPayload(payload)
	fmt.Println("slack-notifier: Sent message to the channel")
}
