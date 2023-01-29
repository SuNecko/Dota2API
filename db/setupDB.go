package db

import (
	"ApiCRUD/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DSN = "host=localhost user=postgres password=nejihiuga75387 dbname=dota port=5432"
var DB *gorm.DB

func SetupDB() {
	var err error

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Database Connected")
	}

	err = DB.AutoMigrate(models.Hero{}, models.Skill{}, models.Item{})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Migrate Successful")
	}
}
