package db

import (
	"log"

	"github.com/borisa99/gin-starter/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate()
	db.AutoMigrate(&models.Book{}, &models.User{}, &models.Role{}, &models.UserRole{})

	return db
}
