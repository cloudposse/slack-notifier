package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SlackNotifier type
// Set DryRun to 'true' to print the message to the console for testing (not send it to the Slack channel)
type SlackNotifier struct {
	WebhookURL string
	DryRun     bool
}

// Payload is a Slack message with attachments
type Payload struct {
	Attachments []Attachment `json:"attachments"`
	LinkNames   bool         `json:"link_names"`
	Mrkdwn      bool         `json:"mrkdwn"`
	IconEmoji   string       `json:"icon_emoji"`
	Username    string       `json:"username"`
	Channel     string       `json:"channel"`
	Thread      string       `json:"thread_ts,omitempty"`
}

// Attachment for a Slack message
// https://api.slack.com/docs/message-attachments
type Attachment struct {
	AuthorIcon string   `json:"author_icon"`
	AuthorLink string   `json:"author_link"`
	AuthorName string   `json:"author_name"`
	Color      string   `json:"color"`
	Fallback   string   `json:"fallback"`
	Fields     []Field  `json:"fields"`
	FooterIcon string   `json:"footer_icon"`
	Footer     string   `json:"footer"`
	ImageURL   string   `json:"image_url"`
	MrkdwnIn   []string `json:"mrkdwn_in"`
	Pretext    string   `json:"pretext"`
	Text       string   `json:"text"`
	ThumbURL   string   `json:"thumb_url"`
	TitleLink  string   `json:"title_link"`
	Title      string   `json:"title"`
	Ts         int64    `json:"ts"`
}

// Field of an attachment
type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

// NewSlackNotifier creates a new SlackNotifier
func NewSlackNotifier(webhookURL string) SlackNotifier {
	return SlackNotifier{
		WebhookURL: webhookURL,
	}
}

// Notify sends a message to the Slack channel
func (sn SlackNotifier) Notify(message Payload) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if sn.DryRun {
		fmt.Println(string(data))
		return nil
	}

	body := bytes.NewBuffer(data)
	request, err := http.NewRequest("POST", sn.WebhookURL, body)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
