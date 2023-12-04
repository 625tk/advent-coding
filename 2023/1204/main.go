package main

import (
	"github.com/625tk/advent-coding/messaging"
	"os"
)

func main() {
	url := os.Getenv("SLACK_WEBHOOK_URL")
	mCli, err := messaging.NewSlack(&url, nil)
	if err != nil {
		panic(err)
	}

	_ = mCli.Post(messaging.Message{
		Channel: "#general",
		Text:    "<!channel> test message",
	})
}
