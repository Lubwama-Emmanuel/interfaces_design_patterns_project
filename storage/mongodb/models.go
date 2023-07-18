package mongodb

import "github.com/Lubwama-Emmanuel/Interfaces/models"

type contact struct {
	Phone string
	Name  string
}

func (c contact) toDataObject() models.DataObject {
	return models.DataObject{
		c.Phone: c.Name,
	}
}
