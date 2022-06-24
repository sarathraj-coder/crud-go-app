package db

import (
	"log"
	"sample_crud/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"
	db, error := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}
	db.AutoMigrate(&models.Book{})

	return db
}
