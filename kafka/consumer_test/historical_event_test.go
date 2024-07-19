package consumer_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/time_capsule/timeline-service/config"
	"github.com/time_capsule/timeline-service/kafka/consumer"
	"github.com/time_capsule/timeline-service/models"
	"github.com/time_capsule/timeline-service/storage/test"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHistoricalEventConsumer(t *testing.T) {
	cfg := config.Load()

	// Create a test topic
	topic := "test-historical-event-topic"
	createTopic(t, cfg.KafkaBrokersTest, topic)
	defer deleteTopic(t, cfg.KafkaBrokersTest, topic)

	// Initialize MongoDB storage for testing
	storage, err := test.NewMongoStorageTest(cfg) // Use NewMongoStorageTest
	if err != nil {
		t.Fatalf("failed to initialize storage: %v", err)
	}

	// Create a test historical event model
	historicalEventModel := &models.CreateHistoricalEventModel{
		ID:          primitive.NewObjectID().Hex(),
		Title:       "Test Historical Event",
		Description: "This is a test historical event.",
		Date:        time.Now(),
		Category:    "Test Category",
		SourceURL:   "https://example.com",
	}

	// Produce a message to the Kafka topic
	produceMessage(t, cfg.KafkaBrokersTest, topic, "historical_event.create", historicalEventModel)

	// Create a HistoricalEventConsumer with the test storage
	consumer := consumer.NewHistoricalEventConsumer(cfg.KafkaBrokersTest, topic, storage)

	// Consume the message
	go func() {
		if err := consumer.Consume(context.Background()); err != nil {
			t.Errorf("Error consuming message: %v", err)
		}
	}()

	// Wait for the message to be consumed (adjust timeout as needed)
	time.Sleep(time.Second * 2)

	// Retrieve the created historical event from the database
	createdEvent, err := storage.HistoricalEvent().GetHistoricalEventByID(context.Background(), historicalEventModel.ID)
	assert.NoError(t, err)
	assert.NotNil(t, createdEvent)

	// Assertions
	assert.Equal(t, historicalEventModel.Title, createdEvent.Title)
	assert.Equal(t, historicalEventModel.Description, createdEvent.Description)
	assert.Equal(t, historicalEventModel.Category, createdEvent.Category)
	assert.Equal(t, historicalEventModel.SourceURL, createdEvent.SourceUrl)
}
