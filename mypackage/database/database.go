package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mypackage/model"
	"os"
)

func NewDB() (*gorm.DB, error) {

	// Connection string comes from DATABASE_URL; the fallback is a local development database.
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&model.Student{})

	if err != nil {
		return nil, err
	}

	fmt.Println("database connected successfully")

	return db, nil
}
