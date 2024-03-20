package databases

import (
	"fmt"
	"log"
	"selling-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	dbPort   = "5432"
	dbname   = "db-go"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting database: ", err)
	}

	fmt.Println("DB Connected")
	db.AutoMigrate(models.User{}, models.Product{}, models.Order{}, models.Invoice{}, models.Payment{})
}

func GetDB() *gorm.DB {
	return db
}
