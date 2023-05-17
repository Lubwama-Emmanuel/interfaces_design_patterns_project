package memory

import "github.com/Lubwama-Emmannuel/Interfaces/models"

type MemoryDatabase struct {
	data models.DataObject
}

func (db *MemoryDatabase) Create(data models.DataObject) error {
	db.data = data
	return nil
}

func (db MemoryDatabase) Read() (models.DataObject, error) {
	return db.data, nil
}

func (db *MemoryDatabase) Update(data models.DataObject) error {
	db.data = data
	return nil
}

func (db *MemoryDatabase) ReadAll() (models.DataObject, error) {
	return db.data, nil
}

func NewMemoryStorage() *MemoryDatabase {
	return &MemoryDatabase{}
}
