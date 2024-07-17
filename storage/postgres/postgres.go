package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/time_capsule/timeline-service/config"
	"github.com/time_capsule/timeline-service/storage"
)

// StorageP implements the storage.StorageIP interface for PostgreSQL.
type StorageP struct {
	db           *pgx.Conn
	CustomEventS storage.CustomEventRepoI
	MilestoneS   storage.MilestoneRepoI
}

// NewPostgresStorage creates a new PostgreSQL storage instance.
func NewPostgresStorage(cfg config.Config) (storage.StorageIP, error) {
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
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
		CustomEventS: NewCustomEventRepo(db),
		MilestoneS:   NewMilestoneRepo(db),
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
