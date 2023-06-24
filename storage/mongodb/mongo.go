package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

type MongoDB struct {
	db_url string
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
		log.Fatal(err)
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func closeDB(client *mongo.Client) {
	// Close the connection when you're done
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func (db *MongoDB) Create(data models.DataObject) error {
	var contact Contact

	client := openDB(db.db_url)

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
		log.Fatal(err)
	}
	fmt.Println("Document inserted successfully!")

	closeDB(client)
	return nil
}

func (db *MongoDB) Read(number string) (models.DataObject, error) {
	client := openDB(db.db_url)
	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	// Find a document
	filter := bson.M{"phone": "1234567890"}
	var result Contact
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	contact := models.DataObject{
		result.Phone: result.Name,
	}

	fmt.Printf("Found document: %+v\n", contact)

	closeDB(client)
	return contact, nil
}

func (db *MongoDB) Update(newData models.DataObject) error {
	client := openDB(db.db_url)

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
		log.Fatal(err)
	}
	fmt.Println("Document updated successfully!")

	closeDB(client)
	return nil
}

func (db *MongoDB) Delete(number string) error {
	client := openDB(db.db_url)

	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	filter := bson.M{"phone": number}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	closeDB(client)
	return nil
}

func (db *MongoDB) ReadAll() ([]models.DataObject, error) {
	client := openDB(db.db_url)
	// Find multiple documents

	// Access the database and collection
	database := client.Database("test_contacts")
	collection := database.Collection("contacts")

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var results []Contact
	var results2 []models.DataObject
	for cur.Next(context.Background()) {
		var contact Contact
		err := cur.Decode(&contact)
		if err != nil {
			log.Fatal(err)
		}
		finalResult := models.DataObject{
			contact.Phone: contact.Name,
		}
		results = append(results, contact)
		results2 = append(results2, finalResult)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found documents: %+v\n", results)
	fmt.Printf("Found two: %+v\n", results2)

	closeDB(client)
	return results2, nil
}

func NewMongoDB(url string) *MongoDB {
	return &MongoDB{db_url: url}
}
