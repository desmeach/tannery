package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultTimeout     = 10 * time.Second
	defaultMaxPoolSize = 100
	defaultMinPoolSize = 0
)

// ClientConfig config for MongoDB Client
type ClientConfig struct {
	URI         string
	Username    string
	Password    string
	MaxPoolSize uint64
	MinPoolSize uint64
}

// NewClient create new MongoDB client with optional parameters
func NewClient(uri, username, password string, maxPoolSize, minPoolSize int) (*mongo.Client, error) {
	maxPoolSizeVal := getMaxPoolSize(maxPoolSize)
	minPoolSizeVal := getMinPoolSize(minPoolSize)

	opts := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(defaultTimeout).
		SetServerSelectionTimeout(defaultTimeout).
		SetMaxPoolSize(maxPoolSizeVal).
		SetMinPoolSize(minPoolSizeVal)

	if username != "" && password != "" {
		opts.SetAuth(options.Credential{
			Username: username,
			Password: password,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	return client, nil
}

// getMaxPoolSize returns max pool size for mongo config
func getMaxPoolSize(size int) uint64 {
	if size <= 0 {
		return defaultMaxPoolSize
	}
	return uint64(size)
}

// getMinPoolSize returns min pool size for mongo config
func getMinPoolSize(size int) uint64 {
	if size <= 0 {
		return defaultMinPoolSize
	}
	return uint64(size)
}

// NewClientWithConfig create mongo client using config as parameter
func NewClientWithConfig(cfg ClientConfig) (*mongo.Client, error) {
	maxPoolSize := cfg.MaxPoolSize
	if maxPoolSize == 0 {
		maxPoolSize = defaultMaxPoolSize
	}

	minPoolSize := cfg.MinPoolSize

	opts := options.Client().
		ApplyURI(cfg.URI).
		SetConnectTimeout(defaultTimeout).
		SetServerSelectionTimeout(defaultTimeout).
		SetMaxPoolSize(maxPoolSize).
		SetMinPoolSize(minPoolSize)

	if cfg.Username != "" && cfg.Password != "" {
		opts.SetAuth(options.Credential{
			Username: cfg.Username,
			Password: cfg.Password,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	return client, nil
}
