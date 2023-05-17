package file_system

import (
	"io/ioutil"
	"os"
)

type FileSystemDatabase struct {
	filename string
}

func (db *FileSystemDatabase) Create(data string) error {
	file, err := os.Create(db.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func (db *FileSystemDatabase) Read() (string, error) {
	// Read the contents of the file
	data, err := ioutil.ReadFile(db.filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (db *FileSystemDatabase) Update(data string) error {
	// Open the file for writing, truncating it if it exists
	file, err := os.OpenFile(db.filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the updated data to the file
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func NewFileSytemDatabase(file string) *FileSystemDatabase {
	return &FileSystemDatabase{filename: file}
}
