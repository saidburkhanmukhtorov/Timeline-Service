package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/time_capsule/timeline-service/config"
	"github.com/time_capsule/timeline-service/genproto/timeline"
	"github.com/time_capsule/timeline-service/kafka/consumer"
	"github.com/time_capsule/timeline-service/service"

	mongodb "github.com/time_capsule/timeline-service/storage/mongo"
	"github.com/time_capsule/timeline-service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	// Initialize PostgreSQL storage
	postgresStorage, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatalf("failed to initialize PostgreSQL storage: %v", err)
	}

	// Initialize MongoDB storage
	mongoStorage, err := mongodb.NewMongoStorage(cfg)
	if err != nil {
		log.Fatalf("failed to initialize MongoDB storage: %v", err)
	}

	// Initialize Kafka consumers
	customEventConsumer := consumer.NewCustomEventConsumer(cfg.KafkaBrokers, "custom_event_topic", postgresStorage)
	milestoneConsumer := consumer.NewMilestoneConsumer(cfg.KafkaBrokers, "milestone_topic", postgresStorage)
	historicalEventConsumer := consumer.NewHistoricalEventConsumer(cfg.KafkaBrokers, "historical_event_topic", mongoStorage)

	// Start consumers in separate goroutines
	go func() {
		if err := customEventConsumer.Consume(context.Background()); err != nil {
			log.Fatalf("custom event consumer error: %v", err)
		}
	}()

	go func() {
		if err := milestoneConsumer.Consume(context.Background()); err != nil {
			log.Fatalf("milestone consumer error: %v", err)
		}
	}()

	go func() {
		if err := historicalEventConsumer.Consume(context.Background()); err != nil {
			log.Fatalf("historical event consumer error: %v", err)
		}
	}()

	// Initialize gRPC server
	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register gRPC services
	timeline.RegisterCustomEventServiceServer(s, service.NewCustomEventService(postgresStorage))
	timeline.RegisterMilestoneServiceServer(s, service.NewMilestoneService(postgresStorage))
	timeline.RegisterHistoricalEventServiceServer(s, service.NewHistoricalEventService(mongoStorage))

	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
