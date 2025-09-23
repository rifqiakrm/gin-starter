package helper

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub/v2"
	"google.golang.org/api/option"

	"gin-starter/config"
)

// initPubSubClient initiates a Pub/Sub v2 client with credentials
func initPubSubClient(ctx context.Context, cfg config.Config) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, cfg.Google.ProjectID, option.WithCredentialsFile(cfg.Google.ServiceAccountFile))
	if err != nil {
		return nil, err
	}
	return client, nil
}

// SendTopic publishes a message to a topic using pubsub/v2
func SendTopic(ctx context.Context, cfg config.Config, topicName string, payload interface{}) error {
	client, err := initPubSubClient(ctx, cfg)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := client.Close(); cerr != nil {
			log.Printf("error closing pubsub client: %v", cerr)
		}
	}()

	publisher := client.Publisher(topicName)

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	result := publisher.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	// Wait to get the server-assigned message ID or error
	_, err = result.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
