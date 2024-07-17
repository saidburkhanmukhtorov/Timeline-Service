package test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/models"
	"github.com/time_capsule/timeline-service/storage/postgres"
)

func TestCustomEventRepo(t *testing.T) {
	db := CreateDBConnection(t)
	defer db.Close(context.Background())

	customEventRepo := postgres.NewCustomEventRepo(db)

	t.Run("CreateCustomEvent", func(t *testing.T) {
		createCustomEventModel := &models.CreateCustomEventModel{
			UserID:      uuid.New().String(),
			Title:       "Test Custom Event",
			Description: "This is a test custom event.",
			Date:        time.Now(),
			Category:    "Test Category",
		}

		createdID, err := customEventRepo.CreateCustomEvent(context.Background(), createCustomEventModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		// Cleanup
		defer deleteCustomEvent(t, db, createdID)
	})

	t.Run("GetCustomEventByID", func(t *testing.T) {
		createCustomEventModel := &models.CreateCustomEventModel{
			UserID:      uuid.New().String(),
			Title:       "Test Custom Event",
			Description: "This is a test custom event.",
			Date:        time.Now(),
			Category:    "Test Category",
		}

		createdID, err := customEventRepo.CreateCustomEvent(context.Background(), createCustomEventModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		customEvent, err := customEventRepo.GetCustomEventByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.NotNil(t, customEvent)
		assert.Equal(t, createdID, customEvent.Id)

		// Cleanup
		defer deleteCustomEvent(t, db, createdID)
	})

	t.Run("GetAllCustomEvents", func(t *testing.T) {
		// Create some test custom events
		createCustomEventModel1 := &models.CreateCustomEventModel{
			UserID:      uuid.New().String(),
			Title:       "Test Custom Event 1",
			Description: "This is test custom event 1.",
			Date:        time.Now(),
			Category:    "Test Category 1",
		}
		createdID1, err := customEventRepo.CreateCustomEvent(context.Background(), createCustomEventModel1)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID1)

		createCustomEventModel2 := &models.CreateCustomEventModel{
			UserID:      uuid.New().String(),
			Title:       "Test Custom Event 2",
			Description: "This is test custom event 2.",
			Date:        time.Now().Add(time.Hour * 24), // Add a day
			Category:    "Test Category 2",
		}
		createdID2, err := customEventRepo.CreateCustomEvent(context.Background(), createCustomEventModel2)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID2)

		// Test GetAllCustomEvents with no filters
		customEvents, err := customEventRepo.GetAllCustomEvents(context.Background(), &timeline.GetAllCustomEventsRequest{})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(customEvents), 2) // At least 2 custom events should be returned

		// Cleanup
		defer deleteCustomEvent(t, db, createdID1)
		defer deleteCustomEvent(t, db, createdID2)
	})

	t.Run("UpdateCustomEvent", func(t *testing.T) {
		createCustomEventModel := &models.CreateCustomEventModel{
			UserID:      uuid.New().String(),
			Title:       "Test Custom Event",
			Description: "This is a test custom event.",
			Date:        time.Now(),
			Category:    "Test Category",
		}

		createdID, err := customEventRepo.CreateCustomEvent(context.Background(), createCustomEventModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		updateCustomEventModel := &models.UpdateCustomEventModel{
			ID:          createdID,
			UserID:      createCustomEventModel.UserID,
			Title:       "Updated Custom Event Title",
			Description: "Updated custom event description.",
			Date:        time.Now().Add(time.Hour * 24), // Add a day
			Category:    "Updated Category",
		}

		err = customEventRepo.UpdateCustomEvent(context.Background(), updateCustomEventModel)
		assert.NoError(t, err)

		updatedCustomEvent, err := customEventRepo.GetCustomEventByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, updateCustomEventModel.Title, updatedCustomEvent.Title)
		assert.Equal(t, updateCustomEventModel.Description, updatedCustomEvent.Description)
		assert.Equal(t, updateCustomEventModel.Category, updatedCustomEvent.Category)
		// ... (Assert other updated fields)

		// Cleanup
		defer deleteCustomEvent(t, db, createdID)
	})

	t.Run("PatchCustomEvent", func(t *testing.T) {
		createCustomEventModel := &models.CreateCustomEventModel{
			UserID:      uuid.New().String(),
			Title:       "Test Custom Event",
			Description: "This is a test custom event.",
			Date:        time.Now(),
			Category:    "Test Category",
		}

		createdID, err := customEventRepo.CreateCustomEvent(context.Background(), createCustomEventModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		newTitle := "Patched Custom Event Title"
		patchCustomEventModel := &models.PatchCustomEventModel{
			ID:    createdID,
			Title: &newTitle,
		}

		err = customEventRepo.PatchCustomEvent(context.Background(), patchCustomEventModel)
		assert.NoError(t, err)

		updatedCustomEvent, err := customEventRepo.GetCustomEventByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, newTitle, updatedCustomEvent.Title) // Title should be patched

		// Cleanup
		defer deleteCustomEvent(t, db, createdID)
	})

	t.Run("DeleteCustomEvent", func(t *testing.T) {
		createCustomEventModel := &models.CreateCustomEventModel{
			UserID:      uuid.New().String(),
			Title:       "Test Custom Event",
			Description: "This is a test custom event.",
			Date:        time.Now(),
			Category:    "Test Category",
		}

		createdID, err := customEventRepo.CreateCustomEvent(context.Background(), createCustomEventModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		err = customEventRepo.DeleteCustomEvent(context.Background(), createdID)
		assert.NoError(t, err)
		_, err = customEventRepo.GetCustomEventByID(context.Background(), createdID)
		assert.ErrorIs(t, err, pgx.ErrNoRows) // Custom event should not be found
	})
}

// Helper function for cleanup
func deleteCustomEvent(t *testing.T, db *pgx.Conn, id string) {
	_, err := db.Exec(context.Background(), "DELETE FROM custom_events WHERE id = $1", id)
	assert.NoError(t, err)
}
