package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

type PhoneNumberStorage struct {
	coll *mongoDB
}

func NewPhoneNumberStorage(db *mongoDB) *PhoneNumberStorage {
	return &PhoneNumberStorage{
		coll: db,
	}
}

func (db *PhoneNumberStorage) Create(data models.DataObject) error {
	var ct contact

	for key, value := range data {
		ct = contact{Phone: key, Name: value}
	}

	_, err := db.coll.InsertOne(context.Background(), ct)
	if err != nil {
		return fmt.Errorf("failed to insert item %w", err)
	}

	return nil
}

func (db *PhoneNumberStorage) Read(number string) (models.DataObject, error) {
	// Find a document
	filter := bson.M{"phone": number}
	var result contact

	err := db.coll.FindOne(context.Background(), filter).Decode(&result)

	return result.toDataObject(), toAppError(err)
}

func (db *PhoneNumberStorage) Update(newData models.DataObject) error {
	var phoneNumber string
	var newName string

	for key, value := range newData {
		phoneNumber = key
		newName = value
	}

	// Update a document
	filter := bson.M{"phone": phoneNumber}
	update := bson.M{"$set": bson.M{"name": newName}}

	_, err := db.coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update item %w", err)
	}

	return nil
}

func (db *PhoneNumberStorage) Delete(number string) error {
	filter := bson.M{"phone": number}

	_, err := db.coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete item %w", err)
	}

	return nil
}

func (db *PhoneNumberStorage) ReadAll() ([]models.DataObject, error) {
	cur, err := db.coll.Find(context.Background(), bson.D{})
	if err != nil {
		return []models.DataObject{}, fmt.Errorf("failed to insert item %w", err)
	}
	defer cur.Close(context.Background())

	var results []models.DataObject

	for cur.Next(context.Background()) {
		var ct contact

		err := cur.Decode(&ct)
		if err != nil {
			return []models.DataObject{}, fmt.Errorf("failed to get items %w", err)
		}
		finalResult := ct.toDataObject()

		results = append(results, finalResult)
	}

	if err := cur.Err(); err != nil {
		return []models.DataObject{}, fmt.Errorf("failed to get items %w", err)
	}

	return results, nil
}
