package service

import (
	"context"
	"fmt"

	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/storage"
)

// CustomEventService implements the timeline.CustomEventServiceServer interface.
type CustomEventService struct {
	storage storage.StorageIP // Assuming you're using PostgreSQL for custom events
	timeline.UnimplementedCustomEventServiceServer
}

// NewCustomEventService creates a new CustomEventService instance.
func NewCustomEventService(storage storage.StorageIP) *CustomEventService {
	return &CustomEventService{
		storage: storage,
	}
}

// GetCustomEventById retrieves a custom event by its ID.
func (s *CustomEventService) GetCustomEventById(ctx context.Context, req *timeline.GetCustomEventByIdRequest) (*timeline.CustomEvent, error) {
	event, err := s.storage.CustomEvent().GetCustomEventByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get custom event by ID: %w", err)
	}

	return event, nil
}

// GetAllCustomEvents retrieves all custom events, optionally filtered by parameters in the request.
func (s *CustomEventService) GetAllCustomEvents(ctx context.Context, req *timeline.GetAllCustomEventsRequest) (*timeline.GetAllCustomEventsResponse, error) {
	events, err := s.storage.CustomEvent().GetAllCustomEvents(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all custom events: %w", err)
	}

	return &timeline.GetAllCustomEventsResponse{
		CustomEvents: events,
		Count:        int32(len(events)), // Assuming you don't have pagination yet
	}, nil
}

// DeleteCustomEvent deletes a custom event by its ID.
func (s *CustomEventService) DeleteCustomEvent(ctx context.Context, req *timeline.DeleteCustomEventRequest) (*timeline.DeleteCustomEventResponse, error) {
	err := s.storage.CustomEvent().DeleteCustomEvent(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete custom event: %w", err)
	}

	return &timeline.DeleteCustomEventResponse{
		Success: true,
	}, nil
}
