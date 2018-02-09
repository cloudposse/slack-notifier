package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SlackNotifier struct {
	Channel    string
	DryRun     bool
	IconEmoji  string
	Username   string
	WebhookURL string
}

// https://api.slack.com/docs/message-attachments
type Payload struct {
	Attachments []Attachment `json:"attachments"`
	Channel     string       `json:"channel"`
	IconEmoji   string       `json:"icon_emoji"`
	LinkNames   bool         `json:"link_names"`
	Mrkdwn      bool         `json:"mrkdwn"`
	Text        string       `json:"text"`
	Username    string       `json:"username"`
}

type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

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

func NewSlackNotifier(channel string, iconEmoji string, username string, webhookURL string) SlackNotifier {
	return SlackNotifier{
		Channel:    channel,
		IconEmoji:  iconEmoji,
		Username:   username,
		WebhookURL: webhookURL,
	}
}

func (sn SlackNotifier) NewPayload() Payload {
	return Payload{
		Channel:   sn.Channel,
		Username:  sn.Username,
		IconEmoji: sn.IconEmoji,
		LinkNames: true,
	}
}

func (payload *Payload) AppendField(field Field, attachmentIndex int) {
	if attachmentIndex >= len(payload.Attachments) {
		return
	}
	payload.Attachments[attachmentIndex].Fields = append(payload.Attachments[attachmentIndex].Fields, field)
}

func (sn SlackNotifier) PostPayload(payload Payload) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if sn.DryRun {
		fmt.Fprintln(os.Stdout, "**dry run**\njson:\n", string(data))
		return nil
	}

	body := bytes.NewBuffer(data)
	request, err := http.NewRequest("POST", sn.WebhookURL, body)
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
