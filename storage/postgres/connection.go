package postgres

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}
type PostgresDB struct { //nolint:revive
	*gorm.DB
}

func NewPostgresDB(config Config, dialector gorm.Dialector) (*PostgresDB, error) {
	// DB string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	)

	if dialector == nil {
		// DB dialector
		dialector = postgres.Open(dsn)
	}

	// Open db
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB %w", err)
	}

	return &PostgresDB{
		DB: db,
	}, nil
}

func (p *PostgresDB) Conn() (*sql.DB, error) {
	db, err := p.DB.DB()
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to get DB connection: %w", err)
	}

	return db, nil
}
