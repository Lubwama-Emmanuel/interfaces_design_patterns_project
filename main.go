package main

import (
	// inmemory "github.com/Lubwama-Emmannuel/Interfaces/in_memory"
	"fmt"

	file_system "github.com/Lubwama-Emmannuel/Interfaces/fileSystem"
)

func main() {
	db := file_system.NewFile()

	// Create a new record
	err := db.Create("Hello, File!")
	if err != nil {
		fmt.Println("an error occurred creating file", err)
	}

	// Read created record
	data, err := db.Read()
	if err != nil {
		fmt.Println("an error occurred reading created file", err)
	}
	fmt.Println(data)

	// Update the record
	err = db.Update("Update worked very well")
	if err != nil {
		fmt.Println("an error occurred reading updating file", err)
	}

	// Read the updated record
	updatedData, err := db.Read()
	if err != nil {
		fmt.Println("an error occurred reading updated file", err)

	}
	fmt.Println(updatedData)

	// db := &inmemory.InMemoryDatabase{}
	// if err := db.Create("Hello, Emmanuel!"); err != nil {
	// 	fmt.Println("An error occurred during read")
	// }

	// data, err := db.Read()
	// if err != nil {
	// 	fmt.Println("An error occurred during reading data")
	// }
	// fmt.Println("Created data", data)

	// if err := db.Update("Hello, Rex!"); err != nil {
	// 	fmt.Println("An error occurred during update")
	// }

	// updatedData, err := db.Read()
	// if err != nil {
	// 	fmt.Println("An error occurred during reading data")
	// }
	// fmt.Println("Updated data", updatedData)

	// if err := db.Delete(); err != nil {
	// 	fmt.Println("An error occurred during delete")
	// }

	// deleted, err := db.Read()
	// if err != nil {
	// 	fmt.Println("An error occurred during reading data")
	// }
	// fmt.Println("Deleted", deleted)

}
