package app

import (
	"fmt"

	"github.com/Lubwama-Emmannuel/Interfaces/models"
)

//go:generate mockgen -destination=mocks/mock_database.go -package=mocks . IDatabase
type IDatabase interface {
	Create(path string, data models.DataObject) error
	Read(path string) (models.DataObject, error)
	// ReadAll() (models.DataObject, error)
	Update(path string, data models.DataObject) error
	// Delete(path string) error
}

type App struct {
	storage IDatabase
}

func NewApp(storage IDatabase) *App {
	return &App{storage: storage}
}

func (a *App) SavePhoneNumber(name, phoneNumber string) error {
	newMap := map[string]string{
		phoneNumber: name,
	}

	return a.storage.Create("a", newMap)
}

func (a *App) GetName(number string) (string, error) {
	var name string

	result, err := a.storage.Read("a")

	if err != nil {
		return "", fmt.Errorf("an error occurred reading %w", err)
	}

	for key, value := range result {
		if key == number {
			name = value
		}
	}

	return name, nil
}

func (a *App) UpdateName(name, phoneNumber string) error {
	updatedName := models.DataObject{
		phoneNumber: name,
	}

	return a.storage.Update("a", updatedName)
}
