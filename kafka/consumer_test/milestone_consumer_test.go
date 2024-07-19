package consumer_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/time_capsule/timeline-service/config"
	"github.com/time_capsule/timeline-service/kafka/consumer"
	"github.com/time_capsule/timeline-service/models"
	"github.com/time_capsule/timeline-service/storage/test"
)

func TestMilestoneConsumer(t *testing.T) {
	cfg := config.Load()

	// Create a test topic
	topic := "test-milestone-topic"
	createTopic(t, cfg.KafkaBrokersTest, topic)
	defer deleteTopic(t, cfg.KafkaBrokersTest, topic)

	// Initialize PostgreSQL storage for testing
	storage, err := test.NewPostgresStorageTest(cfg) // Use NewPostgresStorageTest
	if err != nil {
		t.Fatalf("failed to initialize storage: %v", err)
	}

	// Create a test milestone model
	milestoneModel := &models.CreateMilestoneModel{
		ID:       uuid.NewString(),
		UserID:   uuid.New().String(),
		Title:    "Test Milestone",
		Date:     time.Now(),
		Category: "Test Category",
	}

	// Produce a message to the Kafka topic
	produceMessage(t, cfg.KafkaBrokersTest, topic, "milestone.create", milestoneModel)

	// Create a MilestoneConsumer with the test storage
	consumer := consumer.NewMilestoneConsumer(cfg.KafkaBrokersTest, topic, storage)

	// Consume the message
	go func() {
		if err := consumer.Consume(context.Background()); err != nil {
			t.Errorf("Error consuming message: %v", err)
		}
	}()

	// Wait for the message to be consumed (adjust timeout as needed)
	time.Sleep(time.Second * 2)

	// Retrieve the created milestone from the database
	createdMilestone, err := storage.Milestone().GetMilestoneByID(context.Background(), milestoneModel.ID)
	assert.NoError(t, err)
	assert.NotNil(t, createdMilestone)

	// Assertions
	assert.Equal(t, milestoneModel.UserID, createdMilestone.UserId)
	assert.Equal(t, milestoneModel.Title, createdMilestone.Title)
	assert.Equal(t, milestoneModel.Category, createdMilestone.Category)
}
