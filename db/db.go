package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getConfig() string {
	if db_host := os.Getenv("DATABASE_URL"); db_host != "" {
		return db_host
	}
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	return fmt.Sprintf("host=localhost user=weekendinator password=%s dbname=weekend port=5432 sslmode=disable", db_pass)
}

func Connect() error {
	db, err := gorm.Open(postgres.Open(getConfig()), &gorm.Config{})

	if err != nil {
		return err
	}
	log.Print("Connection established")

	log.Print("Migrating models")
	err = db.AutoMigrate(&User{}, &Token{})
	if err != nil {
		return err
	}
	log.Print("Done")

	DB = db

	return nil
}
