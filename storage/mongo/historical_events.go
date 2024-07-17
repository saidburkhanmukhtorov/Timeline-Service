package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// HistoricalEventRepo implements the storage.HistoricalEventRepoI interface for MongoDB.
type HistoricalEventRepo struct {
	db *mongo.Database
}

// NewHistoricalEventRepo creates a new HistoricalEventRepo instance.
func NewHistoricalEventRepo(db *mongo.Database) *HistoricalEventRepo {
	return &HistoricalEventRepo{
		db: db,
	}
}

// CreateHistoricalEvent creates a new historical event in the database.
func (r *HistoricalEventRepo) CreateHistoricalEvent(ctx context.Context, event *models.CreateHistoricalEventModel) (string, error) {
	var (
		objectID primitive.ObjectID
		err      error
	)

	if event.ID != "" {
		objectID, err = primitive.ObjectIDFromHex(event.ID)
		if err != nil {
			return "", err
		}
	} else {
		objectID = primitive.NewObjectID()
	}

	// Convert the model to a BSON document
	bsonEvent := bson.M{
		"_id":         objectID,
		"title":       event.Title,
		"description": event.Description,
		"date":        event.Date,
		"category":    event.Category,
		"source_url":  event.SourceURL,
		"created_at":  time.Now(),
	}

	// Insert the document into the collection
	result, err := r.db.Collection("historical_events").InsertOne(ctx, bsonEvent)
	if err != nil {
		return "", fmt.Errorf("failed to create historical event: %w", err)
	}

	// Get the inserted ID as a string
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert inserted ID to string")
	}
	return insertedID.Hex(), nil
}

// GetHistoricalEventByID retrieves a historical event by its ID.
func (r *HistoricalEventRepo) GetHistoricalEventByID(ctx context.Context, id string) (*timeline.HistoricalEvent, error) {
	// Convert the string ID to an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid historical event ID: %w", err)
	}

	// Find the document by ID
	var bsonEvent bson.M
	err = r.db.Collection("historical_events").FindOne(ctx, bson.M{"_id": objID}).Decode(&bsonEvent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, fmt.Errorf("failed to get historical event by ID: %w", err)
	}
	// Convert the BSON document to a proto message
	eventModel, err := bsonToHistoricalEvent(bsonEvent)
	if err != nil {
		return nil, err
	}

	return eventModel, nil
}

// GetAllHistoricalEvents retrieves all historical events, optionally filtered by parameters in the request.
func (r *HistoricalEventRepo) GetAllHistoricalEvents(ctx context.Context, req *timeline.GetAllHistoricalEventsRequest) ([]*timeline.HistoricalEvent, error) {
	// Build the filter query based on the request parameters
	filter := bson.M{}
	if req.Title != "" {
		filter["title"] = primitive.Regex{Pattern: req.Title, Options: "i"}
	}
	if req.Description != "" {
		filter["description"] = primitive.Regex{Pattern: req.Description, Options: "i"}
	}
	if req.Category != "" {
		filter["category"] = req.Category
	}
	if req.StartDate != "" {
		startTime, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start date format: %w", err)
		}
		filter["date"] = bson.M{"$gte": startTime}
	}
	if req.EndDate != "" {
		endTime, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid end date format: %w", err)
		}
		filter["date"] = bson.M{"$lte": endTime}
	}

	// Find the documents based on the filter
	cursor, err := r.db.Collection("historical_events").Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get all historical events: %w", err)
	}
	defer cursor.Close(ctx)

	var historicalEvents []*timeline.HistoricalEvent

	// Iterate through the cursor and convert each document to a proto message
	for cursor.Next(ctx) {
		var bsonEvent bson.M
		if err := cursor.Decode(&bsonEvent); err != nil {
			return nil, fmt.Errorf("failed to decode historical event: %w", err)
		}

		eventModel, err := bsonToHistoricalEvent(bsonEvent)
		if err != nil {
			return nil, err
		}
		log.Println(eventModel)
		historicalEvents = append(historicalEvents, eventModel)
	}

	return historicalEvents, nil
}

// UpdateHistoricalEvent updates an existing historical event in the database.
func (r *HistoricalEventRepo) UpdateHistoricalEvent(ctx context.Context, event *models.UpdateHistoricalEventModel) error {
	// Convert the string ID to an ObjectID
	objID, err := primitive.ObjectIDFromHex(event.ID)
	if err != nil {
		return fmt.Errorf("invalid historical event ID: %w", err)
	}

	// Convert the model to a BSON document
	bsonEvent := bson.M{
		"title":       event.Title,
		"description": event.Description,
		"date":        event.Date,
		"category":    event.Category,
		"source_url":  event.SourceURL,
	}

	// Update the document in the collection
	result, err := r.db.Collection("historical_events").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bsonEvent})
	if err != nil {
		return fmt.Errorf("failed to update historical event: %w", err)
	}

	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// PatchHistoricalEvent partially updates an existing historical event in the database.
func (r *HistoricalEventRepo) PatchHistoricalEvent(ctx context.Context, event *models.PatchHistoricalEventModel) error {
	// Convert the string ID to an ObjectID
	objID, err := primitive.ObjectIDFromHex(event.ID)
	if err != nil {
		return fmt.Errorf("invalid historical event ID: %w", err)
	}

	// Build the update document based on the provided fields
	update := bson.M{}
	if event.Title != nil {
		update["title"] = *event.Title
	}
	if event.Description != nil {
		update["description"] = *event.Description
	}
	if event.Date != nil {
		update["date"] = *event.Date
	}
	if event.Category != nil {
		update["category"] = *event.Category
	}
	if event.SourceURL != nil {
		update["source_url"] = *event.SourceURL
	}

	// Update the document in the collection
	result, err := r.db.Collection("historical_events").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("failed to patch historical event: %w", err)
	}

	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// DeleteHistoricalEvent deletes a historical event from the database.
func (r *HistoricalEventRepo) DeleteHistoricalEvent(ctx context.Context, id string) error {
	// Convert the string ID to an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid historical event ID: %w", err)
	}

	// Delete the document from the collection
	result, err := r.db.Collection("historical_events").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return fmt.Errorf("failed to delete historical event: %w", err)
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// bsonToHistoricalEvent converts a BSON document to a timeline.HistoricalEvent proto message.
func bsonToHistoricalEvent(bsonEvent bson.M) (*timeline.HistoricalEvent, error) {
	eventModel := &timeline.HistoricalEvent{} // Initialize an empty message

	// Handle _id field as string
	if val, ok := bsonEvent["_id"].(string); ok {
		eventModel.Id = val
	} else if oid, ok := bsonEvent["_id"].(primitive.ObjectID); ok {
		eventModel.Id = oid.Hex()
	} else {
		return nil, fmt.Errorf("invalid _id type: %T", bsonEvent["_id"])
	}

	// Handle potentially nil fields
	if val, ok := bsonEvent["user_id"].(string); ok {
		eventModel.UserId = val
	}
	if val, ok := bsonEvent["title"].(string); ok {
		eventModel.Title = val
	}
	if val, ok := bsonEvent["description"].(string); ok {
		eventModel.Description = val
	}
	if val, ok := bsonEvent["category"].(string); ok {
		eventModel.Category = val
	}
	if val, ok := bsonEvent["source_url"].(string); ok {
		eventModel.SourceUrl = val
	}

	// Convert date and created_at fields
	if val, ok := bsonEvent["date"].(primitive.DateTime); ok {
		eventModel.Date = val.Time().Format("2006-01-02")
	}
	if val, ok := bsonEvent["created_at"].(primitive.DateTime); ok {
		eventModel.CreatedAt = val.Time().Format(time.RFC3339)
	}

	return eventModel, nil
}
