package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Lubwama-Emmanuel/Interfaces/app"
	// filesystem "github.com/Lubwama-Emmanuel/Interfaces/storage/file_system"
	"github.com/Lubwama-Emmanuel/Interfaces/storage/memory"
)

func main() {
	storage := memory.NewMemoryStorage()
	// storage := filesystem.NewFileSytemDatabase("data.json")

	db := app.NewApp(storage)

	// name := "testName"
	// phone := "1234567890"

	// Create a new record
	saveErr := db.SavePhoneNumber("Test Name", "1234567890")
	if saveErr != nil {
		log.Error("an error occurred creating file: ", saveErr)
	}

	// Read created record
	data, err := db.GetName("1234567890")
	if err != nil {
		log.Error("an error occurred reading created file: ", err)
	}

	log.Info("saved data is: ", data)

	// Update the record

	updateErr := db.UpdateName("1234567890", "Uncle Drizzy")
	if updateErr != nil {
		log.Error("an error occurred reading updating file: ", updateErr)
	}

	// Read the updated record
	updatedData, err := db.GetName("1234567890")
	if err != nil {
		log.Error("an error occurred reading updated file: ", err)
	}

	log.Info("updated data is: ", updatedData)

	// deleteErr := db.DeleteContact("0706039111")
	// if deleteErr != nil {
	// 	log.Error("an error occurred reading updated file", deleteErr)
	// }

	// db.DeleteContact("1234567890")

	phoneNumbers, err := db.GetAllPhoneNumbers()
	if err != nil {
		log.Error("an error occurred getting all numbers: ", err)
	}

	log.Info("numbers", phoneNumbers)
}
