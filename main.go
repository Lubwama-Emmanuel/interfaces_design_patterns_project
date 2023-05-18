package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Lubwama-Emmannuel/Interfaces/app"
	"github.com/Lubwama-Emmannuel/Interfaces/storage/memory"
	// file_system "github.com/Lubwama-Emmannuel/Interfaces/fileSystem"
)

func main() {
	// storage := file_system.NewFileSytemDatabase("data.txt")
	storage := memory.NewMemoryStorage()

	db := app.NewApp(storage)

	name := "Emmanuel"
	phone := "0706039119"

	// Create a new record
	err := db.SavePhoneNumber(name, phone)
	if err != nil {
		log.Error("an error occurred creating file", err)
	}

	// Read created record
	data, err := db.GetName(phone)
	if err != nil {
		log.Error("an error occurred reading created file", err)
	}

	log.Info("saved data is: ", data)

	// Update the record
	updateName := "Lubwama"
	err = db.UpdateName(updateName, phone)
	if err != nil {
		log.Error("an error occurred reading updating file", err)
	}

	// Read the updated record
	updatedData, err := db.GetName(phone)
	if err != nil {
		log.Error("an error occurred reading updated file", err)
	}

	log.Info("updated data is: ", updatedData)

}
