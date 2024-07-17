package service

import (
	"context"
	"fmt"

	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/storage"
)

// MilestoneService implements the timeline.MilestoneServiceServer interface.
type MilestoneService struct {
	storage storage.StorageIP // Assuming you're using PostgreSQL for milestones
	timeline.UnimplementedMilestoneServiceServer
}

// NewMilestoneService creates a new MilestoneService instance.
func NewMilestoneService(storage storage.StorageIP) *MilestoneService {
	return &MilestoneService{
		storage: storage,
	}
}

// GetMilestoneById retrieves a milestone by its ID.
func (s *MilestoneService) GetMilestoneById(ctx context.Context, req *timeline.GetMilestoneByIdRequest) (*timeline.Milestone, error) {
	milestone, err := s.storage.Milestone().GetMilestoneByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get milestone by ID: %w", err)
	}
	return milestone, nil
}

// GetAllMilestones retrieves all milestones, optionally filtered by parameters in the request.
func (s *MilestoneService) GetAllMilestones(ctx context.Context, req *timeline.GetAllMilestonesRequest) (*timeline.GetAllMilestonesResponse, error) {
	milestones, err := s.storage.Milestone().GetAllMilestones(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all milestones: %w", err)
	}

	return &timeline.GetAllMilestonesResponse{
		Milestones: milestones,
		Count:      int32(len(milestones)), // Assuming you don't have pagination yet
	}, nil
}

// DeleteMilestone deletes a milestone by its ID.
func (s *MilestoneService) DeleteMilestone(ctx context.Context, req *timeline.DeleteMilestoneRequest) (*timeline.DeleteMilestoneResponse, error) {
	err := s.storage.Milestone().DeleteMilestone(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete milestone: %w", err)
	}

	return &timeline.DeleteMilestoneResponse{
		Success: true,
	}, nil
}
