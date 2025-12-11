package database

import (
	"log"
	"os"

	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	dsn := os.Getenv("DSN")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	println("connected to PostgreSQL")

}

func AutoMigrate() {
	err := DB.AutoMigrate(&model.User{},&model.MobileNumber{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return DB
}