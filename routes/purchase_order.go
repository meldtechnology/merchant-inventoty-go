package routes

import (
	"github.com/gofiber/fiber/v2"
	"merchant_inventory/api"
	"merchant_inventory/model"
	"merchant_inventory/repository"
	"merchant_inventory/service"
	"merchant_inventory/util"
)

func CreatePurchaseOrder(ctx *fiber.Ctx) error {
	var newPurchaseOrder model.PurchaseOrder
	err := ctx.BodyParser(&newPurchaseOrder)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 400, "Failed to parse request")
	}

	PurchaseOrder, err := service.AddPurchase(newPurchaseOrder)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "PurchaseOrder added successfully", &PurchaseOrder)
}

func UpdatePurchaseOrdersStatus(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid", "")
	status := ctx.Params("status", "IN_TRANSIT")

	err := util.ValidateStatus(status)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	// TODO: Create async for Status and validate the status entry here
	updated, err := service.UpdatePurchase(uuid, status)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "PurchaseOrder Status Updated successfully", &updated)
}

func GetPurchaseOrders(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	pageable, err := service.GetPurchaseOrders(repository.Pageable{
		Page: page, Limit: size,
	})
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "PurchaseOrders Retrieved successfully", &pageable)
}

func GetPurchaseOrderById(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid", "")

	if uuid == "" {
		return api.ErrorHandlerResponse(ctx, 404, "Please provide the purchase order UUID")
	}

	purchaseOrder, err := service.GetPurchaseOrderByUuid(uuid)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "PurchaseOrders Retrieved successfully", &purchaseOrder)
}

func GetPurchaseOrderStatuses(ctx *fiber.Ctx) error {
	statuses, err := service.GetStatues()
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "PurchaseOrders Retrieved successfully", &statuses)
}
