package postgres

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct { //nolint:revive
	*gorm.DB
}

func NewPostgresDB(database string, dialector gorm.Dialector) (*PostgresDB, error) {
	host := viper.GetString("PG_HOST")
	port := viper.GetString("PG_PORT")
	user := viper.GetString("PG_USER")
	password := viper.GetString("PG_PASSWORD")
	// dbName := "phonebook"

	// DB string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database) //nolint:lll

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
