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

// MilestoneConsumer consumes Kafka messages related to milestones.
type MilestoneConsumer struct {
	reader  *kafka.Reader
	storage storage.StorageIP // Use StorageIP for PostgreSQL
}

// NewMilestoneConsumer creates a new MilestoneConsumer instance.
func NewMilestoneConsumer(kafkaBrokers []string, topic string, storage storage.StorageIP) *MilestoneConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaBrokers,
		Topic:   topic,
		GroupID: "milestone-group", // Choose a suitable group ID
	})
	return &MilestoneConsumer{reader: reader, storage: storage}
}

// Consume starts consuming messages from the Kafka topic.
func (c *MilestoneConsumer) Consume(ctx context.Context) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			return fmt.Errorf("error fetching message: %w", err)
		}

		// Determine the message type based on the key
		switch string(msg.Key) {
		case "milestone.create":
			var createModel models.CreateMilestoneModel
			if err := json.Unmarshal(msg.Value, &createModel); err != nil {
				log.Printf("error unmarshalling create milestone message: %v", err)
				continue
			}
			if _, err := c.storage.Milestone().CreateMilestone(ctx, &createModel); err != nil {
				log.Printf("error creating milestone: %v", err)
			}

		case "milestone.update":
			var updateModel models.UpdateMilestoneModel
			if err := json.Unmarshal(msg.Value, &updateModel); err != nil {
				log.Printf("error unmarshalling update milestone message: %v", err)
				continue
			}
			if err := c.storage.Milestone().UpdateMilestone(ctx, &updateModel); err != nil {
				log.Printf("error updating milestone: %v", err)
			}

		case "milestone.patch":
			var patchModel models.PatchMilestoneModel
			if err := json.Unmarshal(msg.Value, &patchModel); err != nil {
				log.Printf("error unmarshalling patch milestone message: %v", err)
				continue
			}
			if err := c.storage.Milestone().PatchMilestone(ctx, &patchModel); err != nil {
				log.Printf("error patching milestone: %v", err)
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
