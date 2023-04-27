package database

import (
	"log"

	"github.com/vauzi/go-crud-api/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	database, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	database.AutoMigrate(&models.Book{})

	return database
}
