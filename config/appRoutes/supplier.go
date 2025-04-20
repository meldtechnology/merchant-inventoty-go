package appRoutes

import (
	"github.com/gofiber/fiber/v2"
	"merchant_inventory/routes"
)

var SUPPLIERS_ENDPOINT = "/api/v1/suppliers"
var SUPPLIERS_WITH_ID_ENDPOINT = "/api/v1/suppliers/:uuid"

func ConfigureSupplierRoutes(app *fiber.App) {
	app.Post(SUPPLIERS_ENDPOINT, routes.CreateSupplier)
	app.Get(SUPPLIERS_ENDPOINT, routes.GetSuppliers)
	app.Get(SUPPLIERS_WITH_ID_ENDPOINT, routes.GetSupplierByUuid)
	app.Patch(SUPPLIERS_WITH_ID_ENDPOINT, routes.UpdateSupplier)
}
