package appRoutes

import (
	"github.com/gofiber/fiber/v2"
	"merchant_inventory/routes"
)

var PURCHASSE_ORDER_ENDPOINT = "/api/v1/purchase-orders"
var PURCHASE_ORDER_WITH_ID_ENDPOINT = "/api/v1/purchase-orders/:uuid"
var PURCHASE_ORDER_UPDATE_ENDPOINT = "/api/v1/purchase-orders/:uuid/status/:status"
var PURCHASE_ORDER_STATUS_ENDPOINT = "/api/v1/purchase-orders/status/public"

func ConfigurePurchaseOrderRoutes(app *fiber.App) {
	app.Post(PURCHASSE_ORDER_ENDPOINT, routes.CreatePurchaseOrder)
	app.Get(PURCHASSE_ORDER_ENDPOINT, routes.GetPurchaseOrders)
	app.Get(PURCHASE_ORDER_WITH_ID_ENDPOINT, routes.GetPurchaseOrderById)
	app.Put(PURCHASE_ORDER_UPDATE_ENDPOINT, routes.UpdatePurchaseOrdersStatus)
	app.Get(PURCHASE_ORDER_STATUS_ENDPOINT, routes.GetPurchaseOrderStatuses)
}
