package mongodb

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/time_capsule/timeline-service/config"
	"github.com/time_capsule/timeline-service/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// StorageM implements the storage.StorageIM interface for MongoDB.
type StorageM struct {
	db               *mongo.Database
	HistoricalEventS storage.HistoricalEventRepoI
}

// NewMongoStorage creates a new MongoDB storage instance.
func NewMongoStorage(cfg config.Config) (storage.StorageIM, error) {
	// Construct MongoDB connection URI
	uri := fmt.Sprintf("mongodb://%s:%d",
		cfg.MongoHost,
		cfg.MongoPort,
	)
	clientOptions := options.Client().ApplyURI(uri).
		SetAuth(options.Credential{Username: cfg.MongoUser, Password: cfg.MongoPassword})

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		slog.Warn("Unable to connect to MongoDB:", err)
		return nil, err
	}

	// Ping the database to verify the connection
	if err := client.Ping(context.Background(), nil); err != nil {
		slog.Warn("Unable to ping MongoDB:", err)
		return nil, err
	}

	db := client.Database(cfg.MongoDB)

	return &StorageM{
		db:               db,
		HistoricalEventS: NewHistoricalEventRepo(db),
	}, nil
}

// HistoricalEvent returns the HistoricalEventRepoI implementation for MongoDB.
func (s *StorageM) HistoricalEvent() storage.HistoricalEventRepoI {
	return s.HistoricalEventS
}
