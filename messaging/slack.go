package messaging

import (
	"context"
	"errors"
	"github.com/slack-go/slack"
)

var _ Messenger = (*Slack)(nil)

type Slack struct {
	webhookURL *string
	token      *string
}

func NewSlack(webhookURL, token *string) (*Slack, error) {
	if webhookURL == nil && token == nil {
		return nil, errors.New("invalid parameter")
	}

	return &Slack{
		webhookURL: webhookURL,
		token:      token,
	}, nil
}

func (r *Slack) Post(m Message) error {
	ctx := context.Background()
	return r.PostWithContext(ctx, m)
}

func (r *Slack) PostWithContext(ctx context.Context, m Message) error {
	if r.webhookURL != nil {
		return slack.PostWebhookContext(ctx, *r.webhookURL, &slack.WebhookMessage{
			Channel: m.Channel,
			Text:    m.Text,
		})
	} else if r.token != nil {
		_, _, err := slack.New(*r.token).PostMessageContext(ctx, m.Channel, slack.MsgOptionText(m.Text, false))
		return err
	}
	return errors.New("invalid parameter")
}
