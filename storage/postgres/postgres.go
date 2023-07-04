package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

type PostgresDB struct { //nolint:revive
	Dialector gorm.Dialector
}

type Contact struct {
	// gorm.Model
	Phone string `gorm:"primaryKey"`
	Name  string
}

func openDB(dialector gorm.Dialector) (*gorm.DB, error) {
	// Open db
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB %w", err)
	}

	return db, nil
}

func (db *PostgresDB) Create(data models.DataObject) error {
	var contact Contact

	for key, value := range data {
		contact = Contact{
			Phone: key,
			Name:  value,
		}
	}

	DB, err := openDB(db.Dialector)
	if err != nil {
		return fmt.Errorf("failed to create contact %w", err)
	}

	DB.Create(&contact)

	return nil
}

func (db *PostgresDB) Read(number string) (models.DataObject, error) {
	var contact Contact

	DB, err := openDB(db.Dialector)
	if err != nil {
		return nil, fmt.Errorf("failed to create contact %w", err)
	}

	DB.First(&contact, number)

	result := models.DataObject{
		contact.Phone: contact.Name,
	}

	return result, nil
}

func (db *PostgresDB) Update(newData models.DataObject) error {
	var contact Contact
	var phoneNumber string
	var phoneName string

	for key, value := range newData {
		phoneNumber = key
		phoneName = value
	}

	DB, err := openDB(db.Dialector)
	if err != nil {
		return fmt.Errorf("failed to create contact %w", err)
	}

	DB.Where("phone=?", phoneNumber).First(&contact)

	contact.Name = phoneName

	DB.Save(&contact)

	return nil
}

func (db *PostgresDB) Delete(number string) error {
	var contact Contact

	DB, err := openDB(db.Dialector)
	if err != nil {
		return fmt.Errorf("failed to create contact %w", err)
	}

	DB.Delete(&contact, number)

	return nil
}

func (db *PostgresDB) ReadAll() ([]models.DataObject, error) {
	var contacts []Contact

	DB, err := openDB(db.Dialector)
	if err != nil {
		return nil, fmt.Errorf("failed to create contact %w", err)
	}

	DB.Find(&contacts)

	var results []models.DataObject //nolint:prealloc

	for _, value := range contacts {
		finalResult := models.DataObject{
			value.Phone: value.Name,
		}
		results = append(results, finalResult)
	}

	return results, nil
}

func NewPostgresDB(database string) *PostgresDB {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "1234567890"
	// dbName := "phonebook"

	// DB string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database) //nolint:lll

	dia := postgres.Open(dsn)

	return &PostgresDB{Dialector: dia}
}
