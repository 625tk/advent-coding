package messaging

import "context"

type Message struct {
	Channel string
	Text    string
	Level   string
}

type Messenger interface {
	Post(Message) error
	PostWithContext(context.Context, Message) error
}
