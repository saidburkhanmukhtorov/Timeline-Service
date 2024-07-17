package models

import "time"

// CreateCustomEventModel represents the data structure for creating a new custom event.
type CreateCustomEventModel struct {
	ID          string    `json:"id" bson:"_id"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Date        time.Time `json:"date" bson:"date"`
	Category    string    `json:"category" bson:"category"`
}

// UpdateCustomEventModel represents the data structure for updating an existing custom event (PUT).
type UpdateCustomEventModel struct {
	ID          string    `json:"id" bson:"_id"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Date        time.Time `json:"date" bson:"date"`
	Category    string    `json:"category" bson:"category"`
}

// PatchCustomEventModel represents the data structure for partially updating an existing custom event (PATCH).
type PatchCustomEventModel struct {
	ID          string     `json:"id" bson:"_id"`
	Title       *string    `json:"title,omitempty" bson:"title,omitempty"`
	Description *string    `json:"description,omitempty" bson:"description,omitempty"`
	Date        *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Category    *string    `json:"category,omitempty" bson:"category,omitempty"`
}
