package app

import (
	"fmt"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

//go:generate mockgen -destination=mocks/mock_database.go -package=mocks . IDatabase
type IDatabase interface {
	Create(data models.DataObject) error
	Read(number string) (models.DataObject, error)
	ReadAll() ([]models.DataObject, error)
	Update(data models.DataObject) error
	Delete(number string) error
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

	if err := a.storage.Create(newMap); err != nil {
		return fmt.Errorf("an error occurred during creating contact %w", err)
	}

	return nil
}

func (a *App) GetName(number string) (string, error) {
	var name string

	result, err := a.storage.Read(number)
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

func (a *App) UpdateName(phoneNumber, name string) error {
	updatedName := map[string]string{
		phoneNumber: name,
	}

	if err := a.storage.Update(updatedName); err != nil {
		return fmt.Errorf("an error occurred during updating contact %w", err)
	}

	return nil
}

func (a *App) DeleteContact(phoneNumber string) error {
	err := a.storage.Delete(phoneNumber)
	if err != nil {
		return fmt.Errorf("an error occurred reading %w", err)
	}

	return nil
}

func (a *App) GetAllPhoneNumbers() ([]string, error) {
	data, err := a.storage.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("an error occurred reading %w", err)
	}

	var phoneNumbers []string

	for _, contact := range data {
		for phoneNumber := range contact {
			phoneNumbers = append(phoneNumbers, phoneNumber)
		}
	}

	return phoneNumbers, nil
}
