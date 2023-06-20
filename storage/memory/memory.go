package memory

import (
	"fmt"

	"github.com/Lubwama-Emmannuel/Interfaces/models"
)

type MemoryDatabase struct { //nolint:revive
	data []models.DataObject
}

func (db *MemoryDatabase) Create(data models.DataObject) error {
	db.data = append(db.data, data)

	return nil
}

func (db MemoryDatabase) Read(number string) (models.DataObject, error) {

	var data models.DataObject
	
	for _, value := range db.data {
		for phoneNumber, phoneName := range value {
			if phoneNumber == number { 
				data = models.DataObject{
					phoneNumber: phoneName,
				}
				return data, nil
			}
		}
	}

	return models.DataObject{}, nil
}

func (db *MemoryDatabase) Update(data models.DataObject) error {
	var phone string
	var newName string

	for key, value := range data {
		phone = key
		newName = value
	}

	for _, key := range db.data {
		for phoneNumber, _ := range key {
			if phoneNumber == phone {
				key[phoneNumber] = newName
				return nil
			}
		}
	}

	return fmt.Errorf("number not saved")
}

func (db *MemoryDatabase) Delete(number string) error {
	for i, obj := range db.data {
		for phoneNumber, _ := range obj {
			if phoneNumber == number {
				db.data = append(db.data[:i], db.data[i + 1:]... )

			}
		}
		fmt.Println(i, obj)
	}
	return nil
}

func (db *MemoryDatabase) ReadAll() ([]models.DataObject, error) {
	return db.data, nil
}

func NewMemoryStorage() *MemoryDatabase {
	return &MemoryDatabase{}
}
