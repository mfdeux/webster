package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// NewDatabase creates a new data store/database
func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=administrator dbname=sentry password=mysecretpassword sslmode=disable")
	// db, err := gorm.Open("sqlite3", "test.db")
	return db, err
}
