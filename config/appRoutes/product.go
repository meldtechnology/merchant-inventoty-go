package appRoutes

import (
	"github.com/gofiber/fiber/v2"
	"merchant_inventory/routes"
)

var PRODUCTS_ENDPOINT = "/api/v1/products"
var PRODUCTS_WITH_ID_ENDPOINT = "/api/v1/products/:uuid"
var PRODUCTS_WITH_SKU_ENDPOINT = "/api/v1/products/sku/:sku"

func ConfigureProductRoutes(app *fiber.App) {
	app.Post(PRODUCTS_ENDPOINT, routes.CreateProduct)
	app.Get(PRODUCTS_ENDPOINT, routes.GetProducts)
	app.Get(PRODUCTS_WITH_ID_ENDPOINT, routes.GetProductByUuid)
	app.Get(PRODUCTS_WITH_SKU_ENDPOINT, routes.GetProductBySku)
	app.Patch(PRODUCTS_WITH_ID_ENDPOINT, routes.UpdateProduct)
}
