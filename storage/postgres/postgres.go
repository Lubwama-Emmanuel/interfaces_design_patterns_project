package postgres

import (
	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

type PhoneNumberStorage struct { //nolint:revi
	conn *PostgresDB
}

func NewPhoneNumberStorage(db *PostgresDB) *PhoneNumberStorage {
	return &PhoneNumberStorage{
		conn: db,
	}
}

func (db *PhoneNumberStorage) Create(data models.DataObject) error {
	var ct contact

	for key, value := range data {
		ct = contact{
			Phone: key,
			Name:  value,
		}
	}

	query := db.conn.Create(&ct)

	return toAppError(query.Error)
}

func (db *PhoneNumberStorage) Read(number string) (models.DataObject, error) {
	var contact contact

	query := db.conn.First(&contact, number)

	return contact.toDataObject(), toAppError(query.Error)
}

func (db *PhoneNumberStorage) Update(newData models.DataObject) error {
	var contact contact
	var phoneNumber string
	var phoneName string

	for key, value := range newData {
		phoneNumber = key
		phoneName = value
	}

	db.conn.Where("phone=?", phoneNumber).First(&contact)

	contact.Name = phoneName

	query := db.conn.Save(&contact)

	return toAppError(query.Error)
}

func (db *PhoneNumberStorage) Delete(number string) error {
	var contact contact

	query := db.conn.Delete(&contact, number)

	return toAppError(query.Error)
}

func (db *PhoneNumberStorage) ReadAll() ([]models.DataObject, error) {
	var contacts []contact

	query := db.conn.Find(&contacts)

	var results []models.DataObject //nolint:prealloc

	for _, value := range contacts {
		finalResult := value.toDataObject()
		results = append(results, finalResult)
	}

	return results, toAppError(query.Error)
}
