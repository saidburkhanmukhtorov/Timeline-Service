package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/helper"
	"github.com/time_capsule/timeline-service/models"
)

// MilestoneRepo implements the storage.MilestoneRepoI interface for PostgreSQL.
type MilestoneRepo struct {
	db *pgx.Conn
}

// NewMilestoneRepo creates a new MilestoneRepo instance.
func NewMilestoneRepo(db *pgx.Conn) *MilestoneRepo {
	return &MilestoneRepo{
		db: db,
	}
}

// CreateMilestone creates a new milestone in the database.
func (r *MilestoneRepo) CreateMilestone(ctx context.Context, milestone *models.CreateMilestoneModel) (string, error) {
	if milestone.ID == "" {
		milestone.ID = uuid.NewString()
	}
	query := `
		INSERT INTO milestones (
			id,
			user_id,
			title,
			date,
			category,
			created_at
		) VALUES (
			$1, $2, $3, $4, $5, NOW()
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		milestone.ID,
		milestone.UserID,
		milestone.Title,
		milestone.Date,
		milestone.Category,
	).Scan(&milestone.ID)

	if err != nil {
		return "", fmt.Errorf("failed to create milestone: %w", err)
	}

	return milestone.ID, nil
}

// GetMilestoneByID retrieves a milestone by its ID.
func (r *MilestoneRepo) GetMilestoneByID(ctx context.Context, id string) (*timeline.Milestone, error) {
	var (
		milestoneModel timeline.Milestone
		date           sql.NullTime
		createdAt      sql.NullTime
	)
	query := `
		SELECT 
			id,
			user_id,
			title,
			date,
			category,
			created_at
		FROM milestones
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&milestoneModel.Id,
		&milestoneModel.UserId,
		&milestoneModel.Title,
		&date,
		&milestoneModel.Category,
		&createdAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, fmt.Errorf("failed to get milestone by ID: %w", err)
	}

	milestoneModel.Date = helper.DateToString(date)
	milestoneModel.CreatedAt = helper.DateToString(createdAt)
	return &milestoneModel, nil
}

// GetAllMilestones retrieves all milestones, optionally filtered by parameters in the request.
func (r *MilestoneRepo) GetAllMilestones(ctx context.Context, req *timeline.GetAllMilestonesRequest) ([]*timeline.Milestone, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT
			id,
			user_id,
			title,
			date,
			category,
			created_at
		FROM 
			milestones
		WHERE 1=1 
	`

	filter := ""

	if req.UserId != "" {
		filter += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	if req.Title != "" {
		filter += fmt.Sprintf(" AND title ILIKE $%d", count)
		args = append(args, "%"+req.Title+"%")
		count++
	}

	if req.Category != "" {
		filter += fmt.Sprintf(" AND category ILIKE $%d", count)
		args = append(args, "%"+req.Category+"%")
		count++
	}

	if req.StartDate != "" {
		startTime, err := time.Parse(time.RFC3339, req.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start time format: %w", err)
		}
		filter += fmt.Sprintf(" AND date >= $%d", count)
		args = append(args, startTime)
		count++
	}

	if req.EndDate != "" {
		endTime, err := time.Parse(time.RFC3339, req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid end time format: %w", err)
		}
		filter += fmt.Sprintf(" AND date <= $%d", count)
		args = append(args, endTime)
		count++
	}

	query += filter

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get all milestones: %w", err)
	}
	defer rows.Close()

	var milestones []*timeline.Milestone

	for rows.Next() {
		var (
			milestoneModel timeline.Milestone
			date           sql.NullTime
			createdAt      sql.NullTime
		)
		err = rows.Scan(
			&milestoneModel.Id,
			&milestoneModel.UserId,
			&milestoneModel.Title,
			&date,
			&milestoneModel.Category,
			&createdAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan milestone: %w", err)
		}
		milestoneModel.Date = helper.DateToString(date)
		milestoneModel.CreatedAt = helper.DateToString(createdAt)
		milestones = append(milestones, &milestoneModel)
	}

	return milestones, nil
}

// UpdateMilestone updates an existing milestone in the database.
func (r *MilestoneRepo) UpdateMilestone(ctx context.Context, milestone *models.UpdateMilestoneModel) error {
	query := `
		UPDATE milestones
		SET 
			user_id = $1,
			title = $2,
			date = $3,
			category = $4
		WHERE id = $5
	`

	result, err := r.db.Exec(ctx, query,
		milestone.UserID,
		milestone.Title,
		milestone.Date,
		milestone.Category,
		milestone.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update milestone: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// PatchMilestone partially updates an existing milestone in the database.
func (r *MilestoneRepo) PatchMilestone(ctx context.Context, milestone *models.PatchMilestoneModel) error {
	var args []interface{}
	count := 1
	query := `
		UPDATE milestones
		SET 
	`

	filter := ""

	if milestone.Title != nil {
		filter += fmt.Sprintf(" title = $%d, ", count)
		args = append(args, *milestone.Title)
		count++
	}

	if milestone.Date != nil {
		filter += fmt.Sprintf(" date = $%d, ", count)
		args = append(args, *milestone.Date)
		count++
	}

	if milestone.Category != nil {
		filter += fmt.Sprintf(" category = $%d, ", count)
		args = append(args, *milestone.Category)
		count++
	}

	if filter == "" {
		return fmt.Errorf("at least one field to update is required")
	}

	filter = filter[:len(filter)-2] // Remove the trailing comma and space
	query += filter + fmt.Sprintf(" WHERE id = $%d", count)
	args = append(args, milestone.ID)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to patch milestone: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// DeleteMilestone deletes a milestone from the database.
func (r *MilestoneRepo) DeleteMilestone(ctx context.Context, id string) error {
	query := `
		DELETE FROM milestones
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete milestone: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
