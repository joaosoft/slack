package main

import (
	"github.com/joaosoft/slack"
)

func main() {
	slack := slack.New()

	config := &slack.Config{
		Service: "Service",
		Url:     app.Config().Slack.SlackHookUrl,
		Enabled: app.Config().Slack.Enabled,
	}
	return ds.Slack.Init(config)

	// write to slack
	msg := &slack.Message{}
	msg.Title("Title").
		Action("Action").
		Environment("Environment").
		Code("Code")

	_ = slack().Send(msg)
}
