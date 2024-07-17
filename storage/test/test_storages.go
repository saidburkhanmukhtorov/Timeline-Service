package test

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/time_capsule/timeline-service/config"
	"github.com/time_capsule/timeline-service/storage"
	mongodb "github.com/time_capsule/timeline-service/storage/mongo"
	"github.com/time_capsule/timeline-service/storage/postgres"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// StorageP implements the storage.StorageIP interface for PostgreSQL.
type StorageP struct {
	db           *pgx.Conn
	CustomEventS storage.CustomEventRepoI
	MilestoneS   storage.MilestoneRepoI
}

// StorageM implements the storage.StorageIM interface for MongoDB.
type StorageM struct {
	db               *mongo.Database
	HistoricalEventS storage.HistoricalEventRepoI
}

// NewPostgresStorage creates a new PostgreSQL storage instance.
func NewPostgresStorageTest(cfg config.Config) (storage.StorageIP, error) {
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"postgres",
	)

	db, err := pgx.Connect(context.Background(), dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:", err)
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		slog.Warn("Unable to ping database:", err)
		return nil, err
	}

	return &StorageP{
		db:           db,
		CustomEventS: postgres.NewCustomEventRepo(db),
		MilestoneS:   postgres.NewMilestoneRepo(db),
	}, nil
}

// CustomEvent returns the CustomEventRepoI implementation for PostgreSQL.
func (s *StorageP) CustomEvent() storage.CustomEventRepoI {
	return s.CustomEventS
}

// Milestone returns the MilestoneRepoI implementation for PostgreSQL.
func (s *StorageP) Milestone() storage.MilestoneRepoI {
	return s.MilestoneS
}

// NewMongoStorage creates a new MongoDB storage instance.
func NewMongoStorageTest(cfg config.Config) (storage.StorageIM, error) {

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
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
		HistoricalEventS: mongodb.NewHistoricalEventRepo(db),
	}, nil
}

// HistoricalEvent returns the HistoricalEventRepoI implementation for MongoDB.
func (s *StorageM) HistoricalEvent() storage.HistoricalEventRepoI {
	return s.HistoricalEventS
}
