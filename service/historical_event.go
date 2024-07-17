package service

import (
	"context"
	"fmt"

	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/storage"
)

// HistoricalEventService implements the timeline.HistoricalEventServiceServer interface.
type HistoricalEventService struct {
	storage storage.StorageIM // Assuming you're using MongoDB for historical events
	timeline.UnimplementedHistoricalEventServiceServer
}

// NewHistoricalEventService creates a new HistoricalEventService instance.
func NewHistoricalEventService(storage storage.StorageIM) *HistoricalEventService {
	return &HistoricalEventService{
		storage: storage,
	}
}

// GetHistoricalEventById retrieves a historical event by its ID.
func (s *HistoricalEventService) GetHistoricalEventById(ctx context.Context, req *timeline.GetHistoricalEventByIdRequest) (*timeline.HistoricalEvent, error) {
	event, err := s.storage.HistoricalEvent().GetHistoricalEventByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get historical event by ID: %w", err)
	}

	return event, nil
}

// GetAllHistoricalEvents retrieves all historical events, optionally filtered by parameters in the request.
func (s *HistoricalEventService) GetAllHistoricalEvents(ctx context.Context, req *timeline.GetAllHistoricalEventsRequest) (*timeline.GetAllHistoricalEventsResponse, error) {
	events, err := s.storage.HistoricalEvent().GetAllHistoricalEvents(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all historical events: %w", err)
	}

	return &timeline.GetAllHistoricalEventsResponse{
		HistoricalEvents: events,
		Count:            int32(len(events)), // Assuming you don't have pagination yet
	}, nil
}

// DeleteHistoricalEvent deletes a historical event by its ID.
func (s *HistoricalEventService) DeleteHistoricalEvent(ctx context.Context, req *timeline.DeleteHistoricalEventRequest) (*timeline.DeleteHistoricalEventResponse, error) {
	err := s.storage.HistoricalEvent().DeleteHistoricalEvent(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete historical event: %w", err)
	}

	return &timeline.DeleteHistoricalEventResponse{
		Success: true,
	}, nil
}
