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

func openDB(url string) *mongo.Client {
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

func closeDB(client *mongo.Client) {
	// Close the connection when you're done
	err := client.Disconnect(context.Background())
	if err != nil {
		return
	}
}

func (db *MongoDB) Create(data models.DataObject) error {
	var contact Contact

	client := openDB(db.databaseURL)

	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	for key, value := range data {
		contact = Contact{Phone: key, Name: value}
	}

	// Insert a document
	// person := Person{Name: "John Doe", Email: "johndoe@example.com", Age: 30}

	_, err := collection.InsertOne(context.Background(), contact)
	if err != nil {
		return fmt.Errorf("failed to insert item %w", err)
	}
	// fmt.Println("Document inserted successfully!")

	closeDB(client)

	return nil
}

func (db *MongoDB) Read(number string) (models.DataObject, error) {
	client := openDB(db.databaseURL)
	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	// Find a document
	filter := bson.M{"phone": "1234567890"}
	var result Contact

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return models.DataObject{}, fmt.Errorf("failed to find item %w", err)
	}

	contact := models.DataObject{
		result.Phone: result.Name,
	}

	// fmt.Printf("Found document: %+v\n", contact)

	closeDB(client)

	return contact, nil
}

func (db *MongoDB) Update(newData models.DataObject) error {
	client := openDB(db.databaseURL)

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
	// fmt.Println("Document updated successfully!")

	closeDB(client)

	return nil
}

func (db *MongoDB) Delete(number string) error {
	client := openDB(db.databaseURL)

	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	filter := bson.M{"phone": number}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete item %w", err)
	}

	closeDB(client)

	return nil
}

func (db *MongoDB) ReadAll() ([]models.DataObject, error) {
	client := openDB(db.databaseURL)
	// Find multiple documents

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

	closeDB(client)

	return results, nil
}

func NewMongoDB(url string) *MongoDB {
	return &MongoDB{databaseURL: url}
}
