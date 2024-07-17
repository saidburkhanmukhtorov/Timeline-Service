package storage

import (
	"context"

	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/models"
)

// StorageIP defines the interface for interacting with the PostgreSQL storage layer.
type StorageIP interface {
	CustomEvent() CustomEventRepoI
	Milestone() MilestoneRepoI
}

// StorageIM defines the interface for interacting with the MongoDB storage layer.
type StorageIM interface {
	HistoricalEvent() HistoricalEventRepoI
}

// CustomEventRepoI defines methods for interacting with custom events in PostgreSQL.
type CustomEventRepoI interface {
	CreateCustomEvent(ctx context.Context, event *models.CreateCustomEventModel) (string, error)
	GetCustomEventByID(ctx context.Context, id string) (*timeline.CustomEvent, error)
	GetAllCustomEvents(ctx context.Context, req *timeline.GetAllCustomEventsRequest) ([]*timeline.CustomEvent, error)
	UpdateCustomEvent(ctx context.Context, event *models.UpdateCustomEventModel) error
	PatchCustomEvent(ctx context.Context, event *models.PatchCustomEventModel) error
	DeleteCustomEvent(ctx context.Context, id string) error
}

// MilestoneRepoI defines methods for interacting with milestones in PostgreSQL.
type MilestoneRepoI interface {
	CreateMilestone(ctx context.Context, milestone *models.CreateMilestoneModel) (string, error)
	GetMilestoneByID(ctx context.Context, id string) (*timeline.Milestone, error)
	GetAllMilestones(ctx context.Context, req *timeline.GetAllMilestonesRequest) ([]*timeline.Milestone, error)
	UpdateMilestone(ctx context.Context, milestone *models.UpdateMilestoneModel) error
	PatchMilestone(ctx context.Context, milestone *models.PatchMilestoneModel) error
	DeleteMilestone(ctx context.Context, id string) error
}

// HistoricalEventRepoI defines methods for interacting with historical events in MongoDB.
type HistoricalEventRepoI interface {
	CreateHistoricalEvent(ctx context.Context, event *models.CreateHistoricalEventModel) (string, error)
	GetHistoricalEventByID(ctx context.Context, id string) (*timeline.HistoricalEvent, error)
	GetAllHistoricalEvents(ctx context.Context, req *timeline.GetAllHistoricalEventsRequest) ([]*timeline.HistoricalEvent, error)
	UpdateHistoricalEvent(ctx context.Context, event *models.UpdateHistoricalEventModel) error
	PatchHistoricalEvent(ctx context.Context, event *models.PatchHistoricalEventModel) error
	DeleteHistoricalEvent(ctx context.Context, id string) error
}
