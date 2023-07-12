package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Lubwama-Emmanuel/Interfaces/app"
	// "github.com/Lubwama-Emmanuel/Interfaces/storage/mongodb"
	"github.com/Lubwama-Emmanuel/Interfaces/storage/postgres"
)

func main() {
	// storage := memory.NewMemoryStorage()
	// storage := filesystem.NewFileSytemDatabase("data.json")
	// storage := mongodb.NewMongoDB("mongodb://localhost:27017")
	pg, err := postgres.NewPostgresDB("phonebook", nil)
	if err != nil {
		log.WithError(err).Fatal("an error occurred while connecting to postgresql")
	}

	storage := postgres.NewPhoneNumberStorage(pg)

	db := app.NewApp(storage)

	// name := "testName"
	// phone := "1234567890"

	// Create a new record
	saveErr := db.SavePhoneNumber("Gift", "07047286821")
	if saveErr != nil {
		log.Error("an error occurred creating file: ", saveErr)
	}

	// saveErr := db.SavePhoneNumber("Emmanuel", "0782640437")
	// if saveErr != nil {
	// 	log.Error("an error occurred creating file: ", saveErr)
	// }

	// Read created record
	data, err := db.GetName("0706039119")
	if err != nil {
		log.Error("an error occurred reading created file: ", err)
	}

	log.Info("saved data is: ", data)

	// Update the record

	// updateErr := db.UpdateName("0706039119", "Uncle Drizzy")
	// if updateErr != nil {
	// 	log.Error("an error occurred reading updating file: ", updateErr)
	// }

	// // Read the updated record
	// updatedData, err := db.GetName("0706039119")
	// if err != nil {
	// 	log.Error("an error occurred reading updated file: ", err)
	// }

	// log.Info("updated data is: ", updatedData)

	// db.DeleteContact("1234567890")

	phoneNumbers, err := db.GetAllPhoneNumbers()
	if err != nil {
		log.Error("an error occurred getting all numbers: ", err)
	}

	log.Info("numbers", phoneNumbers)
}
