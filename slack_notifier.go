package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackNotifier struct {
	WebhookUrl string
	DryRun     bool
}

type Payload struct {
	Attachments []Attachment `json:"attachments"`
	LinkNames   bool         `json:"link_names"`
	Mrkdwn      bool         `json:"mrkdwn"`
	IconEmoji   string       `json:"icon_emoji"`
	Username    string       `json:"username"`
}

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

type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

func NewSlackNotifier(webhookURL string) SlackNotifier {
	return SlackNotifier{
		WebhookUrl: webhookURL,
	}
}

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
	request, err := http.NewRequest("POST", sn.WebhookUrl, body)
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
