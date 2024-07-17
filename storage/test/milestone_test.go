package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/models"
	"github.com/time_capsule/timeline-service/storage/postgres"
)

func CreateDBConnection(t *testing.T) *pgx.Conn {
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"postgres",
	)

	// Connecting to postgres
	db, err := pgx.Connect(context.Background(), dbCon)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	return db
}

func TestMilestoneRepo(t *testing.T) {
	db := CreateDBConnection(t)
	defer db.Close(context.Background())

	milestoneRepo := postgres.NewMilestoneRepo(db)

	t.Run("CreateMilestone", func(t *testing.T) {
		createMilestoneModel := &models.CreateMilestoneModel{
			UserID:   uuid.New().String(),
			Title:    "Test Milestone",
			Date:     time.Now(),
			Category: "Test Category",
		}

		createdID, err := milestoneRepo.CreateMilestone(context.Background(), createMilestoneModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		// Cleanup
		defer deleteMilestone(t, db, createdID)
	})

	t.Run("GetMilestoneByID", func(t *testing.T) {
		createMilestoneModel := &models.CreateMilestoneModel{
			UserID:   uuid.New().String(),
			Title:    "Test Milestone",
			Date:     time.Now(),
			Category: "Test Category",
		}

		createdID, err := milestoneRepo.CreateMilestone(context.Background(), createMilestoneModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		milestone, err := milestoneRepo.GetMilestoneByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.NotNil(t, milestone)
		assert.Equal(t, createdID, milestone.Id)

		// Cleanup
		defer deleteMilestone(t, db, createdID)
	})

	t.Run("GetAllMilestones", func(t *testing.T) {
		// Create some test milestones
		createMilestoneModel1 := &models.CreateMilestoneModel{
			UserID:   uuid.New().String(),
			Title:    "Test Milestone 1",
			Date:     time.Now(),
			Category: "Test Category 1",
		}
		createdID1, err := milestoneRepo.CreateMilestone(context.Background(), createMilestoneModel1)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID1)

		createMilestoneModel2 := &models.CreateMilestoneModel{
			UserID:   uuid.New().String(),
			Title:    "Test Milestone 2",
			Date:     time.Now().Add(time.Hour * 24), // Add a day
			Category: "Test Category 2",
		}
		createdID2, err := milestoneRepo.CreateMilestone(context.Background(), createMilestoneModel2)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID2)

		// Test GetAllMilestones with no filters
		milestones, err := milestoneRepo.GetAllMilestones(context.Background(), &timeline.GetAllMilestonesRequest{})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(milestones), 2) // At least 2 milestones should be returned

		// Cleanup
		defer deleteMilestone(t, db, createdID1)
		defer deleteMilestone(t, db, createdID2)
	})

	t.Run("UpdateMilestone", func(t *testing.T) {
		createMilestoneModel := &models.CreateMilestoneModel{
			UserID:   uuid.New().String(),
			Title:    "Test Milestone",
			Date:     time.Now(),
			Category: "Test Category",
		}

		createdID, err := milestoneRepo.CreateMilestone(context.Background(), createMilestoneModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		updateMilestoneModel := &models.UpdateMilestoneModel{
			ID:       createdID,
			UserID:   createMilestoneModel.UserID,
			Title:    "Updated Milestone Title",
			Date:     time.Now().Add(time.Hour * 24), // Add a day
			Category: "Updated Category",
		}

		err = milestoneRepo.UpdateMilestone(context.Background(), updateMilestoneModel)
		assert.NoError(t, err)

		updatedMilestone, err := milestoneRepo.GetMilestoneByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, updateMilestoneModel.Title, updatedMilestone.Title)
		assert.Equal(t, updateMilestoneModel.Category, updatedMilestone.Category)

		// Cleanup
		defer deleteMilestone(t, db, createdID)
	})

	t.Run("PatchMilestone", func(t *testing.T) {
		createMilestoneModel := &models.CreateMilestoneModel{
			UserID:   uuid.New().String(),
			Title:    "Test Milestone",
			Date:     time.Now(),
			Category: "Test Category",
		}

		createdID, err := milestoneRepo.CreateMilestone(context.Background(), createMilestoneModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		newTitle := "Patched Milestone Title"
		patchMilestoneModel := &models.PatchMilestoneModel{
			ID:    createdID,
			Title: &newTitle,
		}

		err = milestoneRepo.PatchMilestone(context.Background(), patchMilestoneModel)
		assert.NoError(t, err)

		updatedMilestone, err := milestoneRepo.GetMilestoneByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, newTitle, updatedMilestone.Title)

		// Cleanup
		defer deleteMilestone(t, db, createdID)
	})

	t.Run("DeleteMilestone", func(t *testing.T) {
		createMilestoneModel := &models.CreateMilestoneModel{
			UserID:   uuid.New().String(),
			Title:    "Test Milestone",
			Date:     time.Now(),
			Category: "Test Category",
		}

		createdID, err := milestoneRepo.CreateMilestone(context.Background(), createMilestoneModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		err = milestoneRepo.DeleteMilestone(context.Background(), createdID)
		assert.NoError(t, err)

		_, err = milestoneRepo.GetMilestoneByID(context.Background(), createdID)
		assert.ErrorIs(t, err, pgx.ErrNoRows) // Milestone should not be found
	})
}

// Helper function for cleanup
func deleteMilestone(t *testing.T, db *pgx.Conn, id string) {
	_, err := db.Exec(context.Background(), "DELETE FROM milestones WHERE id = $1", id)
	assert.NoError(t, err)
}
