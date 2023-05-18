package memory

import (
	"github.com/Lubwama-Emmannuel/Interfaces/models"
)

type MemoryDatabase struct { //nolint:revive
	data map[string]models.DataObject
}

func (db *MemoryDatabase) Create(path string, data models.DataObject) error {
	newData := map[string]models.DataObject{
		path: data,
	}
	db.data = newData

	return nil
}

func (db MemoryDatabase) Read(path string) (models.DataObject, error) {
	for key, value := range db.data {
		if key == path {
			return value, nil
		}
	}

	return models.DataObject{}, nil
}

func (db *MemoryDatabase) Update(path string, data models.DataObject) error {
	var phone string
	var newName string

	for key, value := range data {
		phone = key
		newName = value
	}

	for value, key := range db.data {
		if value == path {
			key[phone] = newName
		}
	}

	return nil
}

func NewMemoryStorage() *MemoryDatabase {
	return &MemoryDatabase{}
}
