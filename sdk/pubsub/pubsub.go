package pubsub

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub/v2"
	"google.golang.org/api/option"
)

// Subscriber is an interface that defines the methods a pubsub receiver must implement.
type Subscriber interface {
	SubscriptionName() string
	ProcessMessage(context.Context, *pubsub.Message)
}

// PubSub is a pubsub engine.
type PubSub struct {
	*pubsub.Client
}

// NewPubSub creates an instance of PubSub v2 client.
func NewPubSub(projectID string, credentialFile *string) *PubSub {
	var client *pubsub.Client
	var err error

	ctx := context.Background()
	if credentialFile != nil {
		client, err = pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(*credentialFile))
	} else {
		client, err = pubsub.NewClient(ctx, projectID)
	}
	if err != nil {
		panic(err)
	}

	return &PubSub{
		Client: client,
	}
}

// StartSubscriptions loops subscriptions and receive it from google
func (ps *PubSub) StartSubscriptions(subscribers ...Subscriber) error {
	for _, s := range subscribers {
		go func(sub Subscriber) {
			subscriber := ps.Subscriber(sub.SubscriptionName())
			err := subscriber.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
				sub.ProcessMessage(ctx, m)
				m.Ack() // acknowledge after processing
			})
			if err != nil {
				log.Printf("Error subscribing %s: %v\n", sub.SubscriptionName(), err)
				panic(err)
			}
		}(s)
	}
	return nil
}
