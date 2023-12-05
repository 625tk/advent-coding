package messaging

import (
	"context"
	"errors"
	"github.com/bwmarrin/discordgo"
)

var _ Messenger = (*Discord)(nil)

type Discord struct {
	sess *discordgo.Session
}

func NewDiscord(token *string) (*Discord, error) {
	if token == nil {
		return nil, errors.New("invalid parameter")
	}

	v, err := discordgo.New(*token)
	if err != nil {
		return nil, err
	}

	return &Discord{
		sess: v,
	}, nil
}
func (r *Discord) Post(m Message) error {
	return r.PostWithContext(context.Background(), m)
}

func (r *Discord) PostWithContext(ctx context.Context, m Message) error {
	_, err := r.sess.ChannelMessageSend(m.Channel, m.Text, discordgo.WithContext(ctx))
	return err
}
