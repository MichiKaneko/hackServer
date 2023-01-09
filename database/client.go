package database

import (
	"log"

	"github.com/MichiKaneko/hackServer/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect(connectionString string) {
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	log.Println("Database migrated")
}
