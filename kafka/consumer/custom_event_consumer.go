package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/time_capsule/timeline-service/models"
	"github.com/time_capsule/timeline-service/storage"
)

// CustomEventConsumer consumes Kafka messages related to custom events.
type CustomEventConsumer struct {
	reader  *kafka.Reader
	storage storage.StorageIP // Use StorageIP for PostgreSQL
}

// NewCustomEventConsumer creates a new CustomEventConsumer instance.
func NewCustomEventConsumer(kafkaBrokers []string, topic string, storage storage.StorageIP) *CustomEventConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaBrokers,
		Topic:   topic,
		GroupID: "custom-event-group", // Choose a suitable group ID
	})
	return &CustomEventConsumer{reader: reader, storage: storage}
}

// Consume starts consuming messages from the Kafka topic.
func (c *CustomEventConsumer) Consume(ctx context.Context) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			return fmt.Errorf("error fetching message: %w", err)
		}

		// Determine the message type based on the key
		switch string(msg.Key) {
		case "custom_event.create":
			var createModel models.CreateCustomEventModel
			if err := json.Unmarshal(msg.Value, &createModel); err != nil {
				log.Printf("error unmarshalling create custom event message: %v", err)
				continue
			}
			if _, err := c.storage.CustomEvent().CreateCustomEvent(ctx, &createModel); err != nil {
				log.Printf("error creating custom event: %v", err)
			}

		case "custom_event.update":
			var updateModel models.UpdateCustomEventModel
			if err := json.Unmarshal(msg.Value, &updateModel); err != nil {
				log.Printf("error unmarshalling update custom event message: %v", err)
				continue
			}
			if err := c.storage.CustomEvent().UpdateCustomEvent(ctx, &updateModel); err != nil {
				log.Printf("error updating custom event: %v", err)
			}

		case "custom_event.patch":
			var patchModel models.PatchCustomEventModel
			if err := json.Unmarshal(msg.Value, &patchModel); err != nil {
				log.Printf("error unmarshalling patch custom event message: %v", err)
				continue
			}
			if err := c.storage.CustomEvent().PatchCustomEvent(ctx, &patchModel); err != nil {
				log.Printf("error patching custom event: %v", err)
			}

		default:
			log.Printf("unknown message key: %s", msg.Key)
		}

		// Commit the message
		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			return fmt.Errorf("error committing message: %w", err)
		}
	}
}
