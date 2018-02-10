package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	webhook_url    = flag.String("webhook_url", os.Getenv("SLACK_WEBHOOK_URL"), "Slack Webhook URL")
	fallback       = flag.String("fallback", os.Getenv("SLACK_FALLBACK"), "A plain-text summary of the attachment. This text will be used in clients that don't show formatted text")
	color          = flag.String("color", os.Getenv("SLACK_COLOR"), "An optional value that can either be one of good, warning, danger, or any hex color code (e.g. #439FE0)")
	pretext        = flag.String("pretext", os.Getenv("SLACK_PRETEXT"), "Optional text that appears above the message attachment block")
	author_name    = flag.String("author_name", os.Getenv("SLACK_AUTHOR_NAME"), "Small text used to display the author's name")
	author_link    = flag.String("author_link", os.Getenv("SLACK_AUTHOR_LINK"), "URL that will hyperlink the author_name. Will only work if author_name is present")
	author_icon    = flag.String("author_icon", os.Getenv("SLACK_AUTHOR_ICON"), "URL that displays a small 16x16px image to the left of the author_name text. Will only work if author_name is present")
	title          = flag.String("title", os.Getenv("SLACK_TITLE"), "The title is displayed as larger, bold text near the top of a message attachment")
	title_link     = flag.String("title_link", os.Getenv("SLACK_TITLE_LINK"), "URL for the title text to be hyperlinked")
	text           = flag.String("text", os.Getenv("SLACK_TEXT"), "Main text in a message attachment")
	thumb_url      = flag.String("thumb_url", os.Getenv("SLACK_THUMB_URL"), "URL to an image file that will be displayed as a thumbnail on the right side of a message attachment")
	footer         = flag.String("footer", os.Getenv("SLACK_FOOTER"), "Brief text to help contextualize and identify an attachment")
	footer_icon    = flag.String("footer_icon", os.Getenv("SLACK_FOOTER_ICON"), "A small icon beside the footer text")
	image_url      = flag.String("image_url", os.Getenv("SLACK_IMAGE_URL"), "URL to an image file that will be displayed inside a message attachment")
	environment    = flag.String("environment", os.Getenv("SLACK_ENVIRONMENT"), "Deployment environment")
	namespace      = flag.String("namespace", os.Getenv("SLACK_NAMESPACE"), "Deployment namespace")
	deployment_url = flag.String("deployment_url", os.Getenv("SLACK_DEPLOYMENT_URL"), "Deployment URL")
	build_url      = flag.String("build_url", os.Getenv("SLACK_BUILD_URL"), "Build URL")
	commit_url     = flag.String("commit_url", os.Getenv("SLACK_COMMIT_URL"), "Commit URL")
)

func main() {
	flag.Parse()

	if *webhook_url == "" {
		flag.PrintDefaults()
		log.Fatal("-webhook_url or SLACK_WEBHOOK_URL required")
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

	fields := []Field{}

	if *environment != "" {
		fields = append(fields, Field{
			Title: "Environment",
			Value: *environment,
			Short: true,
		})
	}
	if *namespace != "" {
		fields = append(fields, Field{
			Title: "Namespace",
			Value: *namespace,
			Short: true,
		})
	}
	if *deployment_url != "" {
		fields = append(fields, Field{
			Title: "Deployment URL",
			Value: *deployment_url,
			Short: false,
		})
	}
	if *build_url != "" {
		fields = append(fields, Field{
			Title: "Build URL",
			Value: *build_url,
			Short: false,
		})
	}
	if *commit_url != "" {
		fields = append(fields, Field{
			Title: "Commit URL",
			Value: *commit_url,
			Short: false,
		})
	}

	if len(fields) > 0 {
		attachment.Fields = fields
	}

	payload := Payload{
		Attachments: []Attachment{attachment},
		LinkNames:   true,
		Mrkdwn:      true,
	}

	notifier := NewSlackNotifier(*webhook_url)
	notifier.Notify(payload)
	fmt.Println("slack-notifier: Sent message to Webhook URL")
}
