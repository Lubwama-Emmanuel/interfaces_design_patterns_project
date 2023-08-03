package mongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

func toAppError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.ErrNotFound
	}

	return err
}
