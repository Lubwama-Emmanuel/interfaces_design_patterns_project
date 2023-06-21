package filesystem

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

type FileSystemDatabase struct { //nolint:revive
	filename string
}

type Contact struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

func (db *FileSystemDatabase) Create(data models.DataObject) error {
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

func (db *FileSystemDatabase) Update(newData models.DataObject) error {
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
			err = saveDataToFile(data, db.filename)
			if err != nil {
				return fmt.Errorf("an error occurred %w", err)
			}
			return nil
		}
	}

	return errors.New("number not found")
}

// Delete function to be implemented here.
func (db *FileSystemDatabase) Delete(number string) error {
	data, err := loadDataFromFile(db.filename)
	if err != nil {
		return fmt.Errorf("an error occurred decoding to json %w", err)
	}

	for i, obj := range data {
		if obj.Phone == number {
			data = append(data[:i], data[i+1:]...)
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

	var numbers []models.DataObject

	for _, obj := range data {
		contact := models.DataObject{
			obj.Phone: obj.Name,
		}
		numbers = append(numbers, contact)
	}

	return numbers, nil
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
		return []Contact{}, fmt.Errorf("failed to load data from file %w", err)
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
	err = ioutil.WriteFile(filePath, jsonData, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write to file %w", err)
	}

	return nil
}

func NewFileSytemDatabase(file string) *FileSystemDatabase {
	return &FileSystemDatabase{filename: file}
}
