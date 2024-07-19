package consumer_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/time_capsule/timeline-service/config"
	"github.com/time_capsule/timeline-service/kafka/consumer"
	"github.com/time_capsule/timeline-service/models"
	"github.com/time_capsule/timeline-service/storage/test"
)

func TestCustomEventConsumer(t *testing.T) {
	cfg := config.Load()

	// Create a test topic
	topic := "test-custom-event-topic"
	createTopic(t, cfg.KafkaBrokersTest, topic)
	defer deleteTopic(t, cfg.KafkaBrokersTest, topic)

	// Initialize PostgreSQL storage for testing
	storage, err := test.NewPostgresStorageTest(cfg) // Use NewPostgresStorageTest
	if err != nil {
		t.Fatalf("failed to initialize storage: %v", err)
	}

	// Create a test custom event model
	customEventModel := &models.CreateCustomEventModel{
		ID:          uuid.NewString(),
		UserID:      uuid.New().String(),
		Title:       "Test Custom Event",
		Description: "This is a test custom event.",
		Date:        time.Now(),
		Category:    "Test Category",
	}

	// Produce a message to the Kafka topic
	produceMessage(t, cfg.KafkaBrokersTest, topic, "custom_event.create", customEventModel)

	// Create a CustomEventConsumer with the test storage
	consumer := consumer.NewCustomEventConsumer(cfg.KafkaBrokersTest, topic, storage)

	// Consume the message
	go func() {
		if err := consumer.Consume(context.Background()); err != nil {
			t.Errorf("Error consuming message: %v", err)
		}
	}()

	// Wait for the message to be consumed (adjust timeout as needed)
	time.Sleep(time.Second * 2)

	// Retrieve the created custom event from the database
	createdEvent, err := storage.CustomEvent().GetCustomEventByID(context.Background(), customEventModel.ID)
	assert.NoError(t, err)
	assert.NotNil(t, createdEvent)

	// Assertions
	assert.Equal(t, customEventModel.UserID, createdEvent.UserId)
	assert.Equal(t, customEventModel.Title, createdEvent.Title)
	assert.Equal(t, customEventModel.Description, createdEvent.Description)
	assert.Equal(t, customEventModel.Category, createdEvent.Category)
}

// Helper functions to create, delete, and produce messages to a Kafka topic
func createTopic(t *testing.T, brokers []string, topic string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, 0)
	if err != nil {
		t.Fatalf("failed to dial leader: %v", err)
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}
}

func deleteTopic(t *testing.T, brokers []string, topic string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, 0)
	if err != nil {
		t.Fatalf("failed to dial leader: %v", err)
	}
	defer conn.Close()

	err = conn.DeleteTopics(topic)
	if err != nil {
		t.Fatalf("failed to delete topic: %v", err)
	}
}

func produceMessage(t *testing.T, brokers []string, topic string, key string, message interface{}) {
	w := &kafka.Writer{
		Addr:  kafka.TCP(brokers...),
		Topic: topic,
	}

	value, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("failed to marshal message: %v", err)
	}

	err = w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: value,
	})
	if err != nil {
		t.Fatalf("failed to write messages: %v", err)
	}

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close writer: %v", err)
	}
}
