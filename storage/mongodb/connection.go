package mongodb

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	*mongo.Collection
}

func connectToDB(url string) *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(url)

	// Connect to mongodb
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil
	}

	log.Info("----DATABASE CONNECTED SUCCESSFULLY")

	return client
}

func NewMongoDB(url string) *MongoDB {
	client := connectToDB(url)

	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	return &MongoDB{
		Collection: collection,
	}
}
