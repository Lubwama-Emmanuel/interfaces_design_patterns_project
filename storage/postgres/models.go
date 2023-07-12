package postgres

import "github.com/Lubwama-Emmanuel/Interfaces/models"

type contact struct {
	// gorm.Model
	Phone string `gorm:"primaryKey"`
	Name  string
}

func (c contact) toDataObject() models.DataObject {
	return models.DataObject{
		c.Phone: c.Name,
	}
}
