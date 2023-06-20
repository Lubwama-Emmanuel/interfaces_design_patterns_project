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
	var fileData []Contact

	// reading existing data
	existingData, err := loadDataFromFile(db.filename)
	if err != nil {
		return fmt.Errorf("an error occurred decoding to json %w", err)
	}

	fileData = append(fileData, existingData...)

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

	fileData = append(fileData, contact)

	err = saveDataToFile(fileData, db.filename)
	if err != nil {
		return fmt.Errorf("an error occurred %w", err)
	}

	return nil
}

func (db *FileSystemDatabase) Read(number string) (models.DataObject, error) {
	data, err := loadDataFromFile(db.filename)
	if err != nil {
		return models.DataObject{}, fmt.Errorf("an error occurred decoding to json %w", err)
	}

	for _, value := range data {
		if number == value.Phone {
			return models.DataObject{
				value.Phone: value.Name,
			}, nil
		}
	}

	return models.DataObject{}, nil
}

func (db *FileSystemDatabase) Update(path string, newData models.DataObject) error {
	var phoneNumber string
	var newName string

	for key, value := range newData {
		phoneNumber = key
		newName = value
	}

	data, err := loadDataFromFile(db.filename)
	if err != nil {
		return fmt.Errorf("an error occurred decoding to json %w", err)
	}

	for i := range data {
		if data[i].Phone == phoneNumber {
			data[i].Name = newName
			break
		}
	}

	err = saveDataToFile(data, db.filename)
	if err != nil {
		return fmt.Errorf("an error occurred %w", err)
	}

	return nil
}

// Delete function to be implemented here.
func (db *FileSystemDatabase) Delete(path string) error {
	data, err := loadDataFromFile(db.filename)
	if err != nil {
		return fmt.Errorf("an error occurred decoding to json %w", err)
	}

	for i, obj := range data {
		if obj.Phone == path { 
			data = append(data[:i], data[i + 1:]...)
			break
		}
	}

	err = saveDataToFile(data, db.filename)
	if err != nil {
		return fmt.Errorf("an error occurred %w", err)
	}

	return nil
}

func (db *FileSystemDatabase) ReadAll() ([]models.DataObject, error) {
	data, err := loadDataFromFile(db.filename)
	if err != nil {
		return []models.DataObject{}, fmt.Errorf("an error occurred decoding to json %w", err)
	}

	var contacts []models.DataObject

	for _, obj := range data {
		contact := models.DataObject{
			obj.Phone: obj.Name,
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
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

	// Unamrshal the JSON data into a slice of Contact objects
	var data []Contact

	// err = json.NewDecoder(fileData).Decode(&data)

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall data %w", err)
	}

	return data, nil
}

func saveDataToFile(data []Contact, filePath string) error {
	// Convert the data to JSON format
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("failed to convert data to json %w", err)
	}

	// Write the JSON data to file
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %w", err)
	}

	return nil
}

func NewFileSytemDatabase(file string) *FileSystemDatabase {
	return &FileSystemDatabase{filename: file}
}
