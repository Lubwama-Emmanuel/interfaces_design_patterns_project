package filesystem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Lubwama-Emmannuel/Interfaces/models"
)

type FileSystemDatabase struct { //nolint:revive
	filename string
}

type Contact struct {
	Path  string `json:"path"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

func (db *FileSystemDatabase) Create(path string, data models.DataObject) error {
	// Open the file in append mode
	file, err := os.OpenFile(db.filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("an error occurred creating a file %w", err)
	}
	defer file.Close()

	got, err := loadDataFromFile(db.filename)
	fmt.Println("error", err)
	fmt.Println("got", got)

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
		return fmt.Errorf("an error occurred during json encoding %w", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("an error occurred writing data %w", err)
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
		return models.DataObject{}, fmt.Errorf("an error occurred decoding to json %w", err)
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

		writeErr := os.WriteFile(db.filename, updatedContact, 0o644) //nolint:gosec
		if err != nil {
			return fmt.Errorf("an error occurred updating data %w", writeErr)
		}
	}

	return nil
}

func loadDataFromFile(filePath string) ([]Contact, error) {
	// Check if file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return []Contact{}, nil
	}

	// Read the JSON data from the file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load data from file %w", err)
	}

	fmt.Println("fileData", string(fileData))

	// Unamrshal the JSON data into a slice of Contact objects
	var data []Contact

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall data %w", err)
	}

	fmt.Println("data is", data)

	return data, nil
}

func NewFileSytemDatabase(file string) *FileSystemDatabase {
	return &FileSystemDatabase{filename: file}
}
