package appRoutes

import (
	"github.com/gofiber/fiber/v2"
	"merchant_inventory/routes"
)

var CUSTOMERS_ENDPOINT = "/api/v1/customers"
var CUSTOMERS_WITH_PHONE_ENDPOINT = "/api/v1/customers/phone/:phone"
var CUSTOMERS_WITH_ID_ENDPOINT = "/api/v1/customers/:uuid"

func ConfigureCustomerRoutes(app *fiber.App) {
	app.Post(CUSTOMERS_ENDPOINT, routes.CreateCustomer)
	app.Get(CUSTOMERS_ENDPOINT, routes.GetCustomers)
	app.Get(CUSTOMERS_WITH_ID_ENDPOINT, routes.GetCustomerByUuid)
	app.Get(CUSTOMERS_WITH_PHONE_ENDPOINT, routes.GetCustomerByPhone)
	app.Patch(CUSTOMERS_WITH_ID_ENDPOINT, routes.UpdateCustomer)
}
