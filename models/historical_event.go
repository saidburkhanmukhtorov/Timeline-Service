package models

import "time"

// CreateHistoricalEventModel represents the data structure for creating a new historical event.
type CreateHistoricalEventModel struct {
	ID          string    `json:"id" bson:"_id"`
	UserID      string    `json:"user_id" bson:"user_id"` // Added user_id field
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Date        time.Time `json:"date" bson:"date"`
	Category    string    `json:"category" bson:"category"`
	SourceURL   string    `json:"source_url" bson:"source_url"`
}

// UpdateHistoricalEventModel represents the data structure for updating an existing historical event (PUT).
type UpdateHistoricalEventModel struct {
	ID          string    `json:"id" bson:"_id"`
	UserID      string    `json:"user_id" bson:"user_id"` // Added user_id field
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Date        time.Time `json:"date" bson:"date"`
	Category    string    `json:"category" bson:"category"`
	SourceURL   string    `json:"source_url" bson:"source_url"`
}

// PatchHistoricalEventModel represents the data structure for partially updating an existing historical event (PATCH).
type PatchHistoricalEventModel struct {
	ID          string     `json:"id" bson:"_id"`
	UserID      *string    `json:"user_id,omitempty" bson:"user_id,omitempty"` // Added user_id field (optional)
	Title       *string    `json:"title,omitempty" bson:"title,omitempty"`
	Description *string    `json:"description,omitempty" bson:"description,omitempty"`
	Date        *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Category    *string    `json:"category,omitempty" bson:"category,omitempty"`
	SourceURL   *string    `json:"source_url,omitempty" bson:"source_url,omitempty"`
}
