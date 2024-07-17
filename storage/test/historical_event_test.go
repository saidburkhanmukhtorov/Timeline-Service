package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/time_capsule/timeline-service/config"
	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/models"
	mongodb "github.com/time_capsule/timeline-service/storage/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createMongoDBConnection(t *testing.T) *mongo.Database {
	cfg := config.Load()

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Unable to connect to MongoDB: %v", err)
	}

	// Ping the database to verify the connection
	if err := client.Ping(context.Background(), nil); err != nil {
		t.Fatalf("Unable to ping MongoDB: %v", err)
	}

	return client.Database(cfg.MongoDB)
}

func TestHistoricalEventRepo(t *testing.T) {
	db := createMongoDBConnection(t)
	historicalEventRepo := mongodb.NewHistoricalEventRepo(db)

	t.Run("CreateHistoricalEvent", func(t *testing.T) {
		testEvent := &models.CreateHistoricalEventModel{
			Title:       "Test Event",
			Description: "This is a test event.",
			Date:        time.Now(),
			Category:    "Test Category",
			SourceURL:   "https://example.com",
		}
		response, err := historicalEventRepo.CreateHistoricalEvent(context.Background(), testEvent)

		assert.NoError(t, err, "CreateHistoricalEvent should not return an error")
		assert.NotNil(t, response, "CreateHistoricalEvent response should not be nil")
		assert.NotEmpty(t, response, "Created event should have a valid ID")
	})

	t.Run("GetHistoricalEventByID", func(t *testing.T) {
		// 1. Create an event to retrieve
		testEvent := &models.CreateHistoricalEventModel{
			Title:       "Test Event for Get",
			Description: "This is a test event for Get.",
			Date:        time.Now(),
			Category:    "Test Category",
			SourceURL:   "https://example.com",
		}
		createResponse, err := historicalEventRepo.CreateHistoricalEvent(context.Background(), testEvent)
		assert.NoError(t, err, "Creating event for GetHistoricalEventByID test failed")
		assert.NotEmpty(t, createResponse, "Created event should have a valid ID")

		// 2. Get the event
		getResponse, err := historicalEventRepo.GetHistoricalEventByID(context.Background(), createResponse)

		assert.NoError(t, err, "GetHistoricalEventByID should not return an error")
		assert.NotNil(t, getResponse, "GetHistoricalEventByID response should not be nil")
		assert.Equal(t, testEvent.Title, getResponse.Title)
		assert.Equal(t, testEvent.Description, getResponse.Description)
	})

	t.Run("GetAllHistoricalEvents", func(t *testing.T) {
		// 1. Create some events
		testEvents := []*models.CreateHistoricalEventModel{
			{
				Title:       "Test Event 1 for GetAll",
				Description: "This is test event 1 for GetAll.",
				Date:        time.Now(),
				Category:    "Test Category 1",
				SourceURL:   "https://example1.com",
			},
			{
				Title:       "Test Event 2 for GetAll",
				Description: "This is test event 2 for GetAll.",
				Date:        time.Now().Add(time.Hour * 24), // Add a day
				Category:    "Test Category 2",
				SourceURL:   "https://example2.com",
			},
		}

		for _, event := range testEvents {
			createRes, err := historicalEventRepo.CreateHistoricalEvent(context.Background(), event)
			assert.NoError(t, err, "Creating event for GetAllHistoricalEvents test failed")
			assert.NotEmpty(t, createRes)
		}

		// 2. List all events
		listResponse, err := historicalEventRepo.GetAllHistoricalEvents(context.Background(), &timeline.GetAllHistoricalEventsRequest{})
		assert.NoError(t, err, "GetAllHistoricalEvents should not return an error")
		assert.NotNil(t, listResponse, "GetAllHistoricalEvents response should not be nil")
		assert.GreaterOrEqual(t, len(listResponse), 2, "Should have at least two events")
	})

	t.Run("UpdateHistoricalEvent", func(t *testing.T) {
		// 1. Create an event
		testEvent := &models.CreateHistoricalEventModel{
			Title:       "Test Event for Update",
			Description: "This is a test event for Update.",
			Date:        time.Now(),
			Category:    "Test Category",
			SourceURL:   "https://example.com",
		}
		createResponse, err := historicalEventRepo.CreateHistoricalEvent(context.Background(), testEvent)
		assert.NoError(t, err, "Creating event for UpdateHistoricalEvent test failed")
		assert.NotEmpty(t, createResponse)

		// 2. Update the event
		updateReq := &models.UpdateHistoricalEventModel{
			ID:          createResponse,
			Title:       "Updated Event Title",
			Description: "Updated event description.",
			Date:        time.Now().Add(time.Hour * 24),
			Category:    "Updated Category",
			SourceURL:   "https://updated-example.com",
		}
		err = historicalEventRepo.UpdateHistoricalEvent(context.Background(), updateReq)
		assert.NoError(t, err, "UpdateHistoricalEvent should not return an error")

		// 3. Retrieve the event and verify the update
		getReq := &timeline.GetHistoricalEventByIdRequest{Id: createResponse}
		getRes, err := historicalEventRepo.GetHistoricalEventByID(context.Background(), getReq.Id)
		assert.NoError(t, err, "GetHistoricalEventByID after update should not return an error")
		assert.Equal(t, updateReq.Title, getRes.Title, "Event title should be updated")
		assert.Equal(t, updateReq.Description, getRes.Description, "Event description should be updated")
	})

	t.Run("PatchHistoricalEvent", func(t *testing.T) {
		// 1. Create an event
		testEvent := &models.CreateHistoricalEventModel{
			Title:       "Test Event for Patch",
			Description: "This is a test event for Patch.",
			Date:        time.Now(),
			Category:    "Test Category",
			SourceURL:   "https://example.com",
		}
		createResponse, err := historicalEventRepo.CreateHistoricalEvent(context.Background(), testEvent)
		assert.NoError(t, err, "Creating event for PatchHistoricalEvent test failed")
		assert.NotEmpty(t, createResponse)

		// 2. Patch the event
		newTitle := "Patched Event Title"
		patchModel := &models.PatchHistoricalEventModel{
			ID:    createResponse,
			Title: &newTitle,
		}
		err = historicalEventRepo.PatchHistoricalEvent(context.Background(), patchModel)
		assert.NoError(t, err, "PatchHistoricalEvent should not return an error")

		// 3. Retrieve the event and verify the patch
		getRes, err := historicalEventRepo.GetHistoricalEventByID(context.Background(), createResponse)
		assert.NoError(t, err, "GetHistoricalEventByID after patch should not return an error")
		assert.Equal(t, newTitle, getRes.Title, "Event title should be patched")
	})

	t.Run("DeleteHistoricalEvent", func(t *testing.T) {
		// 1. Create an event
		testEvent := &models.CreateHistoricalEventModel{
			Title:       "Test Event for Delete",
			Description: "This is a test event for Delete.",
			Date:        time.Now(),
			Category:    "Test Category",
			SourceURL:   "https://example.com",
		}
		createRes, err := historicalEventRepo.CreateHistoricalEvent(context.Background(), testEvent)
		assert.NoError(t, err, "Creating event for DeleteHistoricalEvent test failed")
		assert.NotEmpty(t, createRes)

		// 2. Delete the event
		err = historicalEventRepo.DeleteHistoricalEvent(context.Background(), createRes)
		assert.NoError(t, err, "DeleteHistoricalEvent should not return an error")

		// 3. Attempt to retrieve the deleted event (should fail)
		getRes, err := historicalEventRepo.GetHistoricalEventByID(context.Background(), createRes)
		assert.Error(t, err, "GetHistoricalEventByID after delete should return an error")
		assert.Nil(t, getRes, "GetHistoricalEventByID response should be nil after delete")
	})
}
