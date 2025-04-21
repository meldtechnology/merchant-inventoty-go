package config

import (
	"fmt"
	_ "fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresDatabase struct {
	Db *gorm.DB
}

var Database PostgresDatabase

func buildConnectionString() string {
	DatabaseEnv()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		HOST, USER, PASSWORD, DATABASE_NAME, PORT)
}

func Connect() {
	dsn := buildConnectionString()
	// Open connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connecting to the database", err.Error())
	}

	log.Println("Connected to database successfully")
	Database = PostgresDatabase{Db: db}
}
