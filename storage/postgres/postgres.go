package postgres

import (
	"fmt"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresBD struct {
	database string
}

type Contact struct {
	// gorm.Model
	Phone string
	Name  string
}

func openDB(database string) (*gorm.DB, error) {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "1234567890"
	// dbName := "phonebook"

	// DB string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	// Open db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB %w", err)
	}

	return db, nil
}

func (db *PostgresBD) Create(data models.DataObject) error {
	var contact Contact

	for key, value := range data {
		contact = Contact{
			Phone: key,
			Name:  value,
		}
	}

	DB, err := openDB(db.database)
	if err != nil {
		return fmt.Errorf("failed to create contact %w", err)
	}

	DB.AutoMigrate(&contact)

	DB.Create(&contact)

	return nil
}

func (db *PostgresBD) Read(number string) (models.DataObject, error) {
	var contact Contact
	DB, err := openDB(db.database)
	if err != nil {
		return nil, fmt.Errorf("failed to create contact %w", err)
	}

	DB.Where("phone=?", number).First(&contact)

	result := models.DataObject{
		contact.Phone: contact.Name,
	}

	return result, nil
}

func (db *PostgresBD) Update(newData models.DataObject) error {
	var contact Contact
	var phoneNumber string
	var phoneName string

	for key, value := range newData {
		phoneNumber = key
		phoneName = value
	}

	fmt.Println(phoneNumber)

	DB, err := openDB(db.database)
	if err != nil {
		return fmt.Errorf("failed to create contact %w", err)
	}

	DB.Where("phone=?", phoneNumber).First(&contact)

	contact.Name = phoneName

	DB.Save(&contact)

	return nil
}

func (db *PostgresBD) Delete(number string) error {
	var contact Contact
	DB, err := openDB(db.database)
	if err != nil {
		return fmt.Errorf("failed to create contact %w", err)
	}

	DB.Where("phone=?", number).Delete(&contact)
	return nil
}

func (db *PostgresBD) ReadAll() ([]models.DataObject, error) {
	var contacts []Contact
	DB, err := openDB(db.database)
	if err != nil {
		return nil, fmt.Errorf("failed to create contact %w", err)
	}

	DB.Find(&contacts)

	var results []models.DataObject

	for _, value := range contacts {
		finalResult := models.DataObject{
			value.Phone: value.Name,
		}
		results = append(results, finalResult)
	}

	return results, nil
}

func NewPostgresDB(database string) *PostgresBD {
	return &PostgresBD{database}
}
