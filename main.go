package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	webhookURL  = flag.String("webhook_url", os.Getenv("SLACK_WEBHOOK_URL"), "Slack Webhook URL")
	userName    = flag.String("user_name", os.Getenv("SLACK_USER_NAME"), "Slack user name (the username from which the messages will be sent)")
	iconEmoji   = flag.String("icon_emoji", os.Getenv("SLACK_ICON_EMOJI"), "Slack icon emoji for the user's avatar. https://www.webpagefx.com/tools/emoji-cheat-sheet")
	fallback    = flag.String("fallback", os.Getenv("SLACK_FALLBACK"), "A plain-text summary of the attachment. This text will be used in clients that don't show formatted text")
	color       = flag.String("color", os.Getenv("SLACK_COLOR"), "An optional value that can either be one of good, warning, danger, or any hex color code (e.g. #439FE0)")
	pretext     = flag.String("pretext", os.Getenv("SLACK_PRETEXT"), "Optional text that appears above the message attachment block")
	authorName  = flag.String("author_name", os.Getenv("SLACK_AUTHOR_NAME"), "Small text to display the attachment author's name")
	authorLink  = flag.String("author_link", os.Getenv("SLACK_AUTHOR_LINK"), "URL that will hyperlink the author's name. Will only work if author_name is present")
	authorIcon  = flag.String("author_icon", os.Getenv("SLACK_AUTHOR_ICON"), "URL of a small 16x16px image to the left of the author's name. Will only work if `author_name` is present")
	title       = flag.String("title", os.Getenv("SLACK_TITLE"), "The title is displayed as larger, bold text near the top of a message attachment")
	titleLink   = flag.String("title_link", os.Getenv("SLACK_TITLE_LINK"), "URL for the title text to be hyperlinked")
	text        = flag.String("text", os.Getenv("SLACK_TEXT"), "Main text in a message attachment")
	thumbURL    = flag.String("thumb_url", os.Getenv("SLACK_THUMB_URL"), "URL to an image file that will be displayed as a thumbnail on the right side of a message attachment")
	footer      = flag.String("footer", os.Getenv("SLACK_FOOTER"), "Brief text to help contextualize and identify an attachment")
	footerIcon  = flag.String("footer_icon", os.Getenv("SLACK_FOOTER_ICON"), "URL of a small icon beside the footer text")
	imageURL    = flag.String("image_url", os.Getenv("SLACK_IMAGE_URL"), "URL to an image file that will be displayed inside a message attachment")
	field1Title = flag.String("field1_title", os.Getenv("SLACK_FIELD1_TITLE"), "Field1 title")
	field1Value = flag.String("field1_value", os.Getenv("SLACK_FIELD1_VALUE"), "Field1 value")
	field1Short = flag.String("field1_short", os.Getenv("SLACK_FIELD1_SHORT"), "An optional boolean indicating whether the 'value' is short enough to be displayed side-by-side with other values (default 'false')")
	field2Title = flag.String("field2_title", os.Getenv("SLACK_FIELD2_TITLE"), "Field2 title")
	field2Value = flag.String("field2_value", os.Getenv("SLACK_FIELD2_VALUE"), "Field2 value")
	field2Short = flag.String("field2_short", os.Getenv("SLACK_FIELD2_SHORT"), "An optional boolean indicating whether the 'value' is short enough to be displayed side-by-side with other values (default 'false')")
	field3Title = flag.String("field3_title", os.Getenv("SLACK_FIELD3_TITLE"), "Field3 title")
	field3Value = flag.String("field3_value", os.Getenv("SLACK_FIELD3_VALUE"), "Field3 value")
	field3Short = flag.String("field3_short", os.Getenv("SLACK_FIELD3_SHORT"), "An optional boolean indicating whether the 'value' is short enough to be displayed side-by-side with other values (default 'false')")
	field4Title = flag.String("field4_title", os.Getenv("SLACK_FIELD4_TITLE"), "Field4 title")
	field4Value = flag.String("field4_value", os.Getenv("SLACK_FIELD4_VALUE"), "Field4 value")
	field4Short = flag.String("field4_short", os.Getenv("SLACK_FIELD4_SHORT"), "An optional boolean indicating whether the 'value' is short enough to be displayed side-by-side with other values (default 'false')")
	field5Title = flag.String("field5_title", os.Getenv("SLACK_FIELD5_TITLE"), "Field5 title")
	field5Value = flag.String("field5_value", os.Getenv("SLACK_FIELD5_VALUE"), "Field5 value")
	field5Short = flag.String("field5_short", os.Getenv("SLACK_FIELD5_SHORT"), "An optional boolean indicating whether the 'value' is short enough to be displayed side-by-side with other values (default 'false')")
	field6Title = flag.String("field6_title", os.Getenv("SLACK_FIELD6_TITLE"), "Field6 title")
	field6Value = flag.String("field6_value", os.Getenv("SLACK_FIELD6_VALUE"), "Field6 value")
	field6Short = flag.String("field6_short", os.Getenv("SLACK_FIELD6_SHORT"), "An optional boolean indicating whether the 'value' is short enough to be displayed side-by-side with other values (default 'false')")
	field7Title = flag.String("field7_title", os.Getenv("SLACK_FIELD7_TITLE"), "Field7 title")
	field7Value = flag.String("field7_value", os.Getenv("SLACK_FIELD7_VALUE"), "Field7 value")
	field7Short = flag.String("field7_short", os.Getenv("SLACK_FIELD7_SHORT"), "An optional boolean indicating whether the 'value' is short enough to be displayed side-by-side with other values (default 'false')")
	field8Title = flag.String("field8_title", os.Getenv("SLACK_FIELD8_TITLE"), "Field8 title")
	field8Value = flag.String("field8_value", os.Getenv("SLACK_FIELD8_VALUE"), "Field8 value")
	field8Short = flag.String("field8_short", os.Getenv("SLACK_FIELD8_SHORT"), "An optional boolean indicating whether the 'value' is short enough to be displayed side-by-side with other values (default 'false')")
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

	if *webhookURL == "" {
		flag.PrintDefaults()
		log.Fatal("-webhook_url or SLACK_WEBHOOK_URL required")
	}
	if *userName == "" {
		flag.PrintDefaults()
		log.Fatal("-user_name or SLACK_USER_NAME required")
	}
	if *iconEmoji == "" {
		flag.PrintDefaults()
		log.Fatal("-icon_emoji or SLACK_ICON_EMOJI required")
	}

	attachment := Attachment{
		MrkdwnIn:   []string{"text", "pretext"},
		AuthorIcon: *authorIcon,
		AuthorLink: *authorLink,
		AuthorName: *authorName,
		Color:      *color,
		Fallback:   *fallback,
		FooterIcon: *footerIcon,
		Footer:     *footer,
		ImageURL:   *imageURL,
		Pretext:    *pretext,
		Text:       *text,
		ThumbURL:   *thumbURL,
		TitleLink:  *titleLink,
		Title:      *title,
	}

	fields := addField([]Field{}, *field1Title, *field1Value, *field1Short)
	fields = addField(fields, *field2Title, *field2Value, *field2Short)
	fields = addField(fields, *field3Title, *field3Value, *field3Short)
	fields = addField(fields, *field4Title, *field4Value, *field4Short)
	fields = addField(fields, *field5Title, *field5Value, *field5Short)
	fields = addField(fields, *field6Title, *field6Value, *field6Short)
	fields = addField(fields, *field7Title, *field7Value, *field7Short)
	fields = addField(fields, *field8Title, *field8Value, *field8Short)

	if len(fields) > 0 {
		attachment.Fields = fields
	}

	payload := Payload{
		Attachments: []Attachment{attachment},
		LinkNames:   true,
		Mrkdwn:      true,
		Username:    *userName,
		IconEmoji:   *iconEmoji,
	}

	notifier := NewSlackNotifier(*webhookURL)
	err := notifier.Notify(payload)
	if err != nil {
		fmt.Println("slack-notifier: Failed to sent message to Webhook URL. Error: ", err.Error())
	} else {
		fmt.Println("slack-notifier: Sent message to Webhook URL")
	}
}
