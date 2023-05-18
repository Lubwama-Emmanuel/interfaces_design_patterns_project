package filesystem

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Lubwama-Emmannuel/Interfaces/models"
)

type FileSystemDatabase struct {
	filename string
}

type Contact struct {
	Path  string `json:"path"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

func (db *FileSystemDatabase) Create(path string, data models.DataObject) error {
	file, err := os.Create(db.filename)
	if err != nil {
		return fmt.Errorf("an error occurred creating a file %w", err)
	}
	defer file.Close()

	// Write the data to the file
	var contactName string
	var contactPhone string

	for key, value := range data {
		contactPhone = key
		contactName = value
	}

	contact := Contact{
		Path:  path,
		Phone: contactPhone,
		Name:  contactName,
	}

	// Encode Data to JSON
	jsonData, err := json.MarshalIndent(contact, "", " ")
	if err != nil {
		return fmt.Errorf("an error occured during json encoding %w", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("an error occured writing data %w", err)
	}

	return nil
}

func (db *FileSystemDatabase) Read(path string) (models.DataObject, error) {
	file, openErr := os.Open(db.filename)
	if openErr != nil {
		return models.DataObject{}, fmt.Errorf("an error occurred read contact from file %w", openErr)
	}

	defer file.Close()

	var contact Contact

	err := json.NewDecoder(file).Decode(&contact)
	if err != nil {
		return models.DataObject{}, fmt.Errorf("an error occured decoding to json %w", err)
	}

	if contact.Path == path {
		return models.DataObject{
			contact.Phone: contact.Name,
		}, nil

	}
	return models.DataObject{}, nil
}

func (db *FileSystemDatabase) Update(path string, data models.DataObject) error {
	file, openErr := os.Open(db.filename)
	if openErr != nil {
		return fmt.Errorf("an error occurred read contact from file %w", openErr)
	}

	defer file.Close()

	var contact Contact
	err := json.NewDecoder(file).Decode(&contact)
	if err != nil {
		return fmt.Errorf("an error occurred decoding to json %w", err)
	}

	if path == contact.Path {
		for key, value := range data {
			if key == contact.Phone {
				contact.Name = value
			}
		}

		updatedContact, err := json.MarshalIndent(&contact, "", " ")
		if err != nil {
			return fmt.Errorf("an error occurred during json encoding %w", err)
		}

		// _, writeErr := file.Write(updatedContact)
		writeErr := os.WriteFile(db.filename, updatedContact, 0644)
		if err != nil {
			return fmt.Errorf("an error occurred updating data %w", writeErr)
		}
	}

	return nil
}

func NewFileSytemDatabase(file string) *FileSystemDatabase {
	return &FileSystemDatabase{filename: file}
}
