package config

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
	"strconv"
)

var USER string = ""
var HOST = ""
var PASSWORD = ""
var DATABASE_NAME = ""
var PORT = 0

func initialize() {
	// load .env file from given path
	// we keep it empty it will load .env from the root directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}
}

func DatabaseEnv() {
	initialize()
	USER = os.Getenv("DATABASE_USER")
	HOST = os.Getenv("DATABASE_HOST")
	PASSWORD = os.Getenv("DATABASE_PASSWORD")
	DATABASE_NAME = os.Getenv("DATABASE")
	PORT, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
}
