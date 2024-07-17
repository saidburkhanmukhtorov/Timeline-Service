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

// HistoricalEventConsumer consumes Kafka messages related to historical events.
type HistoricalEventConsumer struct {
	reader  *kafka.Reader
	storage storage.StorageIM // Use StorageIM for MongoDB
}

// NewHistoricalEventConsumer creates a new HistoricalEventConsumer instance.
func NewHistoricalEventConsumer(kafkaBrokers []string, topic string, storage storage.StorageIM) *HistoricalEventConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaBrokers,
		Topic:   topic,
		GroupID: "historical-event-group", // Choose a suitable group ID
	})
	return &HistoricalEventConsumer{reader: reader, storage: storage}
}

// Consume starts consuming messages from the Kafka topic.
func (c *HistoricalEventConsumer) Consume(ctx context.Context) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			return fmt.Errorf("error fetching message: %w", err)
		}

		// Determine the message type based on the key
		switch string(msg.Key) {
		case "historical_event.create":
			var createModel models.CreateHistoricalEventModel
			if err := json.Unmarshal(msg.Value, &createModel); err != nil {
				log.Printf("error unmarshalling create historical event message: %v", err)
				continue
			}
			if _, err := c.storage.HistoricalEvent().CreateHistoricalEvent(ctx, &createModel); err != nil {
				log.Printf("error creating historical event: %v", err)
			}

		case "historical_event.update":
			var updateModel models.UpdateHistoricalEventModel
			if err := json.Unmarshal(msg.Value, &updateModel); err != nil {
				log.Printf("error unmarshalling update historical event message: %v", err)
				continue
			}
			if err := c.storage.HistoricalEvent().UpdateHistoricalEvent(ctx, &updateModel); err != nil {
				log.Printf("error updating historical event: %v", err)
			}

		case "historical_event.patch":
			var patchModel models.PatchHistoricalEventModel
			if err := json.Unmarshal(msg.Value, &patchModel); err != nil {
				log.Printf("error unmarshalling patch historical event message: %v", err)
				continue
			}
			if err := c.storage.HistoricalEvent().PatchHistoricalEvent(ctx, &patchModel); err != nil {
				log.Printf("error patching historical event: %v", err)
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
