package connection

import (
	"assignment_2/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	port     = os.Getenv("DB_PORT")
	dbName   = os.Getenv("DB_NAME")
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	db.Debug().AutoMigrate(models.Orders{}, models.Items{})
}

func GetDB() *gorm.DB {
	return db
}
