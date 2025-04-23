package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrchantinevntory/pkg/adding"
	"github.com/mrchantinevntory/pkg/listing"
	"github.com/mrchantinevntory/pkg/storage/postgres"
)

func main() {

	var add adding.Service
	var list listing.Service

	pg, _ := postgres.NewStorage()

	add = adding.NewService(pg)
	list = listing.NewService(pg)

	http := fiber.New()

	// set up the HTTP server
	//router := http.Handler()

	//http.All("/api/v1", adding, listing)

}
