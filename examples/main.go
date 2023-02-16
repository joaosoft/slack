package main

import "github.com/joaosoft/slack"

func main() {
	config := &slack.Config{
		App: "Service",
		Url: "",
	}

	mySlack := slack.New(config)

	msg := slack.NewMessage().Title("Title").
		Action("Action").
		Environment("Environment").
		Code("Code")

	_ = mySlack.Send(msg)
}
