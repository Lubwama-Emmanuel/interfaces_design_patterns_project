package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Lubwama-Emmanuel/Interfaces/app"
	"github.com/Lubwama-Emmanuel/Interfaces/storage/memory"
)

func main() {
	storage := memory.NewMemoryStorage()
	// storage := filesystem.NewFileSytemDatabase("data.json")

	db := app.NewApp(storage)

	// name := "testName"
	// phone := "1234567890"

	// Create a new record
	err := db.SavePhoneNumber("testname", "1234567890")
	if err != nil {
		log.Error("an error occurred creating file: ", err)
	}

	err = db.SavePhoneNumber("rex", "625487469")
	if err != nil {
		log.Error("an error occurred creating file: ", err)
	}

	saveErr := db.SavePhoneNumber("name", "0987654321")
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
