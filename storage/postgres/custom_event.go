package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/time_capsule/timeline-service/genproto/timeline" // Assuming your proto package is timeline
	"github.com/time_capsule/timeline-service/helper"            // Assuming you have a helper package
	"github.com/time_capsule/timeline-service/models"
)

// CustomEventRepo implements the storage.CustomEventRepoI interface for PostgreSQL.
type CustomEventRepo struct {
	db *pgx.Conn
}

// NewCustomEventRepo creates a new CustomEventRepo instance.
func NewCustomEventRepo(db *pgx.Conn) *CustomEventRepo {
	return &CustomEventRepo{
		db: db,
	}
}

// CreateCustomEvent creates a new custom event in the database.
func (r *CustomEventRepo) CreateCustomEvent(ctx context.Context, event *models.CreateCustomEventModel) (string, error) {
	if event.ID == "" {
		event.ID = uuid.NewString()
	}
	query := `
		INSERT INTO custom_events (
			id,
			user_id,
			title,
			description,
			date,
			category,
			created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, NOW()
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		event.ID,
		event.UserID,
		event.Title,
		event.Description,
		event.Date,
		event.Category,
	).Scan(&event.ID)

	if err != nil {
		return "", fmt.Errorf("failed to create custom event: %w", err)
	}

	return event.ID, nil
}

// GetCustomEventByID retrieves a custom event by its ID.
func (r *CustomEventRepo) GetCustomEventByID(ctx context.Context, id string) (*timeline.CustomEvent, error) {
	var (
		eventModel timeline.CustomEvent
		date       sql.NullTime
		createdAt  sql.NullTime
	)
	query := `
		SELECT 
			id,
			user_id,
			title,
			description,
			date,
			category,
			created_at
		FROM custom_events
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&eventModel.Id,
		&eventModel.UserId,
		&eventModel.Title,
		&eventModel.Description,
		&date,
		&eventModel.Category,
		&createdAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, fmt.Errorf("failed to get custom event by ID: %w", err)
	}

	eventModel.Date = helper.DateToString(date)
	eventModel.CreatedAt = helper.DateToString(createdAt)
	return &eventModel, nil
}

// GetAllCustomEvents retrieves all custom events, optionally filtered by parameters in the request.
func (r *CustomEventRepo) GetAllCustomEvents(ctx context.Context, req *timeline.GetAllCustomEventsRequest) ([]*timeline.CustomEvent, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT
			id,
			user_id,
			title,
			description,
			date,
			category,
			created_at
		FROM 
			custom_events
		WHERE 1=1 
	`

	filter := ""

	if req.UserId != "" {
		filter += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	if req.Title != "" {
		filter += fmt.Sprintf(" AND title LIKE $%d", count)
		args = append(args, "%"+req.Title+"%")
		count++
	}

	if req.Description != "" {
		filter += fmt.Sprintf(" AND description LIKE $%d", count)
		args = append(args, "%"+req.Description+"%")
		count++
	}

	if req.Category != "" {
		filter += fmt.Sprintf(" AND category = $%d", count)
		args = append(args, req.Category)
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
		return nil, fmt.Errorf("failed to get all custom events: %w", err)
	}
	defer rows.Close()

	var customEvents []*timeline.CustomEvent

	for rows.Next() {
		var (
			eventModel timeline.CustomEvent
			date       sql.NullTime
			createdAt  sql.NullTime
		)
		err = rows.Scan(
			&eventModel.Id,
			&eventModel.UserId,
			&eventModel.Title,
			&eventModel.Description,
			&date,
			&eventModel.Category,
			&createdAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan custom event: %w", err)
		}
		eventModel.Date = helper.DateToString(date)
		eventModel.CreatedAt = helper.DateToString(createdAt)
		customEvents = append(customEvents, &eventModel)
	}

	return customEvents, nil
}

// UpdateCustomEvent updates an existing custom event in the database.
func (r *CustomEventRepo) UpdateCustomEvent(ctx context.Context, event *models.UpdateCustomEventModel) error {
	query := `
		UPDATE custom_events
		SET 
			user_id = $1,
			title = $2,
			description = $3,
			date = $4,
			category = $5
		WHERE id = $6
	`

	result, err := r.db.Exec(ctx, query,
		event.UserID,
		event.Title,
		event.Description,
		event.Date,
		event.Category,
		event.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update custom event: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// PatchCustomEvent partially updates an existing custom event in the database.
func (r *CustomEventRepo) PatchCustomEvent(ctx context.Context, event *models.PatchCustomEventModel) error {
	var args []interface{}
	count := 1
	query := `
		UPDATE custom_events
		SET 
	`

	filter := ""

	if event.Title != nil {
		filter += fmt.Sprintf(" title = $%d, ", count)
		args = append(args, *event.Title)
		count++
	}

	if event.Description != nil {
		filter += fmt.Sprintf(" description = $%d, ", count)
		args = append(args, *event.Description)
		count++
	}

	if event.Date != nil {
		filter += fmt.Sprintf(" date = $%d, ", count)
		args = append(args, *event.Date)
		count++
	}

	if event.Category != nil {
		filter += fmt.Sprintf(" category = $%d, ", count)
		args = append(args, *event.Category)
		count++
	}

	if filter == "" {
		return fmt.Errorf("at least one field to update is required")
	}

	filter = filter[:len(filter)-2] // Remove the trailing comma and space
	query += filter + fmt.Sprintf(" WHERE id = $%d", count)
	args = append(args, event.ID)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to patch custom event: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// DeleteCustomEvent deletes a custom event from the database.
func (r *CustomEventRepo) DeleteCustomEvent(ctx context.Context, id string) error {
	query := `
		DELETE FROM custom_events
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete custom event: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
