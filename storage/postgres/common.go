package postgres

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

func toAppError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, sql.ErrNoRows) {
		return models.ErrNotFound
	}

	return err
}
