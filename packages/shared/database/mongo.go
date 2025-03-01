package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func OpenMongo(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	var lastPingErr error
	for {
		// Try to ping the database
		if err := client.Ping(ctx, nil); err == nil {
			return client, nil
		} else {
			lastPingErr = err
		}

		// Wait for 1 second before trying again, unless the context is done
		select {
		case <-time.After(1 * time.Second):
			// continue looping
		case <-ctx.Done():
			return nil, fmt.Errorf("failed to connect to MongoDB: %w", lastPingErr)
		}
	}
}
