package main

import (
	"context"
	"github.com/625tk/advent-coding/messaging"
	"os"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	m, err := messaging.NewDiscord(&token)
	if err != nil {
		panic(err)
	}

	err = m.PostWithContext(context.Background(), messaging.Message{Text: "test"})
	if err != nil {
		panic(err)
	}
}
