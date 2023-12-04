package messaging

import (
	"context"
	"github.com/slack-go/slack"
)

var _ Messenger = (*Slack)(nil)

type Slack struct {
	WebhookURL *string
	Token      *string
}

func (r *Slack) Post(m Message) error {
	ctx := context.Background()
	return r.PostWithContext(ctx, m)
}

func (r *Slack) PostWithContext(ctx context.Context, m Message) error {
	if r.WebhookURL != nil {
		slack.PostWebhook(*r.WebhookURL, &slack.WebhookMessage{
			Channel: m.Channel,
			Text:    m.Text,
		})
		slack.PostMessage
		return r.PostSlackWebhook(m)
	} else {
		slack.post
	}
	return nil
}
