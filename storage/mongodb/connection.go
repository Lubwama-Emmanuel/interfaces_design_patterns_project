package mongodb

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URL string
}

type MongoDB struct {
	*mongo.Collection
	client *mongo.Client
}

func connectToDB(ctx context.Context, url string) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(url)

	// Connect to mongodb
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	log.Info("----DATABASE CONNECTED SUCCESSFULLY")

	return client, nil
}

func NewMongoDB(ctx context.Context, config MongoConfig) (*MongoDB, error) {
	client, err := connectToDB(ctx, config.URL)
	if err != nil {
		return nil, fmt.Errorf("an error occurred: %w", err)
	}

	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	return &MongoDB{
		Collection: collection,
		client:     client,
	}, nil
}

func (m *MongoDB) Close(ctx context.Context) {
	if m.client != nil {
		err := m.client.Disconnect(ctx)
		if err != nil {
			log.WithError(err).Error("failed to close the db")
		}
	}
}
