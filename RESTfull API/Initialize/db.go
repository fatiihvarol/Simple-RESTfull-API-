package Initialize

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=postgres password=pass123 dbname=tsl_employee port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
