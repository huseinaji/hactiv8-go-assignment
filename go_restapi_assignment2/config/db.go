package config

import (
	"fmt"
	"go_restapi_assignment2/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	dbPort   = "5432"
	dbname   = "orders_by"
	DB       *gorm.DB
)

func DBInit() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	database, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting database: ", err)
	}

	database.AutoMigrate(models.Order{}, models.Item{})
	fmt.Println("connected to database")
	DB = database
}
