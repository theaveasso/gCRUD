package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	envErr := godotenv.Load("../../.env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}

	dbSecretKey := os.Getenv("DB_SECRET_KEY")
	d, err := gorm.Open("mysql", dbSecretKey)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
