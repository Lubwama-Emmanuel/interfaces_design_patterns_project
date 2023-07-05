package mongodb

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

type MongoDB struct {
	databaseURL string
}

type Contact struct {
	Phone string
	Name  string
}

func connectToDB(url string) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(url)

	// Connect to mongodb
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb %w", err)
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb %w", err)
	}

	log.Info("----DATABASE CONNECTED SUCCESSFULLY")

	return client, nil
}

func disconnectFromDB(client *mongo.Client) {
	// Close the connection when you're done
	err := client.Disconnect(context.Background())
	if err != nil {
		return
	}
}

func (db *MongoDB) Create(data models.DataObject) error {
	var contact Contact

	client, conncetionErr := connectToDB(db.databaseURL)
	if conncetionErr != nil {
		return fmt.Errorf("failed to connect to database: %w", conncetionErr)
	}

	defer disconnectFromDB(client)

	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	for key, value := range data {
		contact = Contact{Phone: key, Name: value}
	}

	_, err := collection.InsertOne(context.Background(), contact)
	if err != nil {
		return fmt.Errorf("failed to insert item %w", err)
	}

	return nil
}

func (db *MongoDB) Read(number string) (models.DataObject, error) {
	client, conncetionErr := connectToDB(db.databaseURL)
	if conncetionErr != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", conncetionErr)
	}

	defer disconnectFromDB(client)

	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	// Find a document
	filter := bson.M{"phone": number}
	var result Contact

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return models.DataObject{}, fmt.Errorf("failed to find item %w", err)
	}

	contact := models.DataObject{
		result.Phone: result.Name,
	}

	return contact, nil
}

func (db *MongoDB) Update(newData models.DataObject) error {
	client, conncetionErr := connectToDB(db.databaseURL)
	if conncetionErr != nil {
		return fmt.Errorf("failed to connect to database: %w", conncetionErr)
	}

	defer disconnectFromDB(client)

	var phoneNumber string
	var newName string

	for key, value := range newData {
		phoneNumber = key
		newName = value
	}

	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	// Update a document
	filter := bson.M{"phone": phoneNumber}
	update := bson.M{"$set": bson.M{"name": newName}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update item %w", err)
	}

	return nil
}

func (db *MongoDB) Delete(number string) error {
	client, conncetionErr := connectToDB(db.databaseURL)
	if conncetionErr != nil {
		return fmt.Errorf("failed to connect to database: %w", conncetionErr)
	}

	defer disconnectFromDB(client)

	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	filter := bson.M{"phone": number}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete item %w", err)
	}

	return nil
}

func (db *MongoDB) ReadAll() ([]models.DataObject, error) {
	client, conncetionErr := connectToDB(db.databaseURL)
	if conncetionErr != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", conncetionErr)
	}

	defer disconnectFromDB(client)
	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return []models.DataObject{}, fmt.Errorf("failed to insert item %w", err)
	}
	defer cur.Close(context.Background())

	var results []models.DataObject

	for cur.Next(context.Background()) {
		var contact Contact

		err := cur.Decode(&contact)
		if err != nil {
			return []models.DataObject{}, fmt.Errorf("failed to get items %w", err)
		}
		finalResult := models.DataObject{
			contact.Phone: contact.Name,
		}

		results = append(results, finalResult)
	}

	if err := cur.Err(); err != nil {
		return []models.DataObject{}, fmt.Errorf("failed to get items %w", err)
	}

	return results, nil
}

func NewMongoDB(url string) *MongoDB {
	return &MongoDB{databaseURL: url}
}
