package app

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Lubwama-Emmannuel/Interfaces/models"
)

//go:generate mockgen -destination=mocks/mock_database.go -package=mocks . IDatabase
type IDatabase interface {
	Create(data models.DataObject) error
	Read() (models.DataObject, error)
	ReadAll() (models.DataObject, error)
	Update(data models.DataObject) error
	// Delete() error
}

type App struct {
	storage IDatabase
}

func (a *App) SavePhoneNumber(name, phoneNumber string) error {
	newMap := map[string]string{
		name:   phoneNumber,
		"josh": "07833727919",
	}

	return a.storage.Create(newMap)
}

func (a *App) GetName(number string) (string, error) {
	var name string

	result, err := a.storage.Read()
	if err != nil {
		return "", fmt.Errorf("an error occurred reading %w", err)
	}

	for value, key := range result {
		if key == number {
			name = value
		}
	}

	return name, nil
}

func (a *App) UpdatePhoneNumber(name, phoneNumber string) error {
	updatedName := models.DataObject{
		name: phoneNumber,
	}

	return a.storage.Update(updatedName)
}

func NewApp(storage IDatabase) *App {
	return &App{storage: storage}
}

func (a *App) GetAllPhoneNumbers() ([]string, error) {
	var contacts []string

	result, err := a.storage.Read()
	if err != nil {
		return nil, fmt.Errorf("an error occurred reading %w", err)
	}

	for value, key := range result {
		contact := fmt.Sprintf("%v:%v", value, key)
		contacts = append(contacts, contact)
	}

	log.Info("contacts here", contacts)

	return contacts, nil
}
