package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	webhook_url  = flag.String("webhook_url", os.Getenv("SLACK_WEBHOOK_URL"), "Slack Webhook URL")
	user_name    = flag.String("user_name", os.Getenv("SLACK_USER_NAME"), "Slack user name")
	icon_emoji   = flag.String("icon_emoji", os.Getenv("SLACK_ICON_EMOJI"), "Slack icon emoji for the user's avatar. https://www.webpagefx.com/tools/emoji-cheat-sheet")
	fallback     = flag.String("fallback", os.Getenv("SLACK_FALLBACK"), "A plain-text summary of the attachment. This text will be used in clients that don't show formatted text")
	color        = flag.String("color", os.Getenv("SLACK_COLOR"), "An optional value that can either be one of good, warning, danger, or any hex color code (e.g. #439FE0)")
	pretext      = flag.String("pretext", os.Getenv("SLACK_PRETEXT"), "Optional text that appears above the message attachment block")
	author_name  = flag.String("author_name", os.Getenv("SLACK_AUTHOR_NAME"), "Small text used to display the author's name")
	author_link  = flag.String("author_link", os.Getenv("SLACK_AUTHOR_LINK"), "URL that will hyperlink the author_name. Will only work if author_name is present")
	author_icon  = flag.String("author_icon", os.Getenv("SLACK_AUTHOR_ICON"), "URL that displays a small 16x16px image to the left of the author_name text. Will only work if author_name is present")
	title        = flag.String("title", os.Getenv("SLACK_TITLE"), "The title is displayed as larger, bold text near the top of a message attachment")
	title_link   = flag.String("title_link", os.Getenv("SLACK_TITLE_LINK"), "URL for the title text to be hyperlinked")
	text         = flag.String("text", os.Getenv("SLACK_TEXT"), "Main text in a message attachment")
	thumb_url    = flag.String("thumb_url", os.Getenv("SLACK_THUMB_URL"), "URL to an image file that will be displayed as a thumbnail on the right side of a message attachment")
	footer       = flag.String("footer", os.Getenv("SLACK_FOOTER"), "Brief text to help contextualize and identify an attachment")
	footer_icon  = flag.String("footer_icon", os.Getenv("SLACK_FOOTER_ICON"), "A small icon beside the footer text")
	image_url    = flag.String("image_url", os.Getenv("SLACK_IMAGE_URL"), "URL to an image file that will be displayed inside a message attachment")
	field1_title = flag.String("field1_title", os.Getenv("SLACK_FIELD1_TITLE"), "Field1 title")
	field1_value = flag.String("field1_value", os.Getenv("SLACK_FIELD1_VALUE"), "Field1 value")
	field1_short = flag.String("field1_short", os.Getenv("SLACK_FIELD1_SHORT"), "Field1 short")
	field2_title = flag.String("field2_title", os.Getenv("SLACK_FIELD2_TITLE"), "Field2 title")
	field2_value = flag.String("field2_value", os.Getenv("SLACK_FIELD2_VALUE"), "Field2 value")
	field2_short = flag.String("field2_short", os.Getenv("SLACK_FIELD2_SHORT"), "Field2 short")
	field3_title = flag.String("field3_title", os.Getenv("SLACK_FIELD3_TITLE"), "Field3 title")
	field3_value = flag.String("field3_value", os.Getenv("SLACK_FIELD3_VALUE"), "Field3 value")
	field3_short = flag.String("field3_short", os.Getenv("SLACK_FIELD3_SHORT"), "Field3 short")
	field4_title = flag.String("field4_title", os.Getenv("SLACK_FIELD4_TITLE"), "Field4 title")
	field4_value = flag.String("field4_value", os.Getenv("SLACK_FIELD4_VALUE"), "Field4 value")
	field4_short = flag.String("field4_short", os.Getenv("SLACK_FIELD4_SHORT"), "Field4 short")
	field5_title = flag.String("field5_title", os.Getenv("SLACK_FIELD5_TITLE"), "Field5 title")
	field5_value = flag.String("field5_value", os.Getenv("SLACK_FIELD5_VALUE"), "Field5 value")
	field5_short = flag.String("field5_short", os.Getenv("SLACK_FIELD5_SHORT"), "Field5 short")
)

func addField(fields []Field, fieldTitle string, fieldValue string, fieldShort string) []Field {
	if fieldTitle != "" && fieldValue != "" {
		var short = false
		var err error

		if fieldShort != "" {
			if short, err = strconv.ParseBool(fieldShort); err != nil {
				fmt.Println("slack-notifier: Error: ", err.Error())
				short = false
			}
		}

		fields = append(fields, Field{
			Title: fieldTitle,
			Value: fieldValue,
			Short: short,
		})
	}

	return fields
}

func main() {
	flag.Parse()

	if *webhook_url == "" {
		flag.PrintDefaults()
		log.Fatal("-webhook_url or SLACK_WEBHOOK_URL required")
	}
	if *user_name == "" {
		flag.PrintDefaults()
		log.Fatal("-user_name or SLACK_USER_NAME required")
	}
	if *icon_emoji == "" {
		flag.PrintDefaults()
		log.Fatal("-icon_emoji or SLACK_ICON_EMOJI required")
	}

	attachment := Attachment{
		MrkdwnIn:   []string{"text", "pretext"},
		AuthorIcon: *author_icon,
		AuthorLink: *author_link,
		AuthorName: *author_name,
		Color:      *color,
		Fallback:   *fallback,
		FooterIcon: *footer_icon,
		Footer:     *footer,
		ImageURL:   *image_url,
		Pretext:    *pretext,
		Text:       *text,
		ThumbURL:   *thumb_url,
		TitleLink:  *title_link,
		Title:      *title,
	}

	fields := addField([]Field{}, *field1_title, *field1_value, *field1_short)
	fields = addField(fields, *field2_title, *field2_value, *field2_short)
	fields = addField(fields, *field3_title, *field3_value, *field3_short)
	fields = addField(fields, *field4_title, *field4_value, *field4_short)
	fields = addField(fields, *field5_title, *field5_value, *field5_short)

	if len(fields) > 0 {
		attachment.Fields = fields
	}

	payload := Payload{
		Attachments: []Attachment{attachment},
		LinkNames:   true,
		Mrkdwn:      true,
		Username:    *user_name,
		IconEmoji:   *icon_emoji,
	}

	notifier := NewSlackNotifier(*webhook_url)
	err := notifier.Notify(payload)
	if err != nil {
		fmt.Println("slack-notifier: Failed to sent message to Webhook URL. Error: ", err.Error())
	} else {
		fmt.Println("slack-notifier: Sent message to Webhook URL")
	}
}
