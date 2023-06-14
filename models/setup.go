package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dbUser := "root"
	// dbPass := "dbpass"
	// dbName := "go_api"
	// dbHost := "172.17.0.1"
	// dbPort := "3306"
	// database, err := gorm.Open(mysql.Open(dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?tls=false"))
	// if err != nil {
	// 	panic(err)
	// }
	// database.AutoMigrate(&Product{})

	// DB =database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIMEZONE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to initialize database, got error", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal("failed to migrate schema, got error", err)
	}

	fmt.Println("Connected to database")
	DB = db
}
