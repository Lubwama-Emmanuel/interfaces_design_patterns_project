package mongodb

import (
	"errors"
	"fmt"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

func toAppError(err error) error {
	if err == nil {
		return nil
	}

	target := errors.New("no document in result")

	if errors.Is(err, target) {
		fmt.Println("worked")
		return models.ErrNotFound
	}

	return err
}
