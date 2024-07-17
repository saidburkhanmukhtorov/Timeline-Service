package models

import "time"

// CreateMilestoneModel represents the data structure for creating a new milestone.
type CreateMilestoneModel struct {
	ID       string    `json:"id" bson:"_id"`
	UserID   string    `json:"user_id" bson:"user_id"`
	Title    string    `json:"title" bson:"title"`
	Date     time.Time `json:"date" bson:"date"`
	Category string    `json:"category" bson:"category"`
}

// UpdateMilestoneModel represents the data structure for updating an existing milestone (PUT).
type UpdateMilestoneModel struct {
	ID       string    `json:"id" bson:"_id"`
	UserID   string    `json:"user_id" bson:"user_id"`
	Title    string    `json:"title" bson:"title"`
	Date     time.Time `json:"date" bson:"date"`
	Category string    `json:"category" bson:"category"`
}

// PatchMilestoneModel represents the data structure for partially updating an existing milestone (PATCH).
type PatchMilestoneModel struct {
	ID       string     `json:"id" bson:"_id"`
	Title    *string    `json:"title,omitempty" bson:"title,omitempty"`
	Date     *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Category *string    `json:"category,omitempty" bson:"category,omitempty"`
}
