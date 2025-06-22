// Package config provides configuration for the application.
package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

func NewMongoClient(cfg *Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoTimeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %w", err)
	}

	return client, nil
}
