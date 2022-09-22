package database

import (
	"fmt"
	"os"
	"tempsensebackend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// function name has to be uppercase in order it to be exported
func Connect() *gorm.DB {

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DB")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to db: yeet: %s", err)
	}

	db.AutoMigrate(&model.Temperature{})

	return db
}
