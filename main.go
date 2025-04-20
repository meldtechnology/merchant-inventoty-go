package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"merchant_inventory/config"
	"merchant_inventory/config/appRoutes"
)

func setupConfig(app *fiber.App) {
	// Configure Database connection
	config.Connect()
	// Configure Application routes
	appRoutes.ConfigureProductRoutes(app)
	appRoutes.ConfigureCustomerRoutes(app)
	appRoutes.ConfigureSupplierRoutes(app)
	appRoutes.ConfigurePurchaseOrderRoutes(app)

}

func main() {
	// Run the configuration
	log.Println("Running the configurations...")
	app := config.CreatAppServer()
	setupConfig(app)
	config.StartServer(app, 4400)
}
