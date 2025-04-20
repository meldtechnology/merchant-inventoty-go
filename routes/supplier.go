package routes

import (
	"github.com/gofiber/fiber/v2"
	"merchant_inventory/api"
	"merchant_inventory/model"
	"merchant_inventory/repository"
	"merchant_inventory/service"
)

func CreateSupplier(ctx *fiber.Ctx) error {
	var newSupplier model.Supplier
	err := ctx.BodyParser(&newSupplier)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 400, "Failed to parse request")
	}

	Supplier, err := service.AddSupplier(newSupplier)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Supplier added successfully", &Supplier)
}

func GetSuppliers(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	pageable, err := service.GetSuppliers(repository.Pageable{
		Page: page, Limit: size,
	})
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Supplier sRetrieved successfully", &pageable)
}

func GetSupplierByUuid(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid", "")

	Supplier, err := service.GetSuppliersByUuid(uuid)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Supplier Retrieved successfully", &Supplier)
}

func UpdateSupplier(ctx *fiber.Ctx) error {
	var modifiedSupplier model.Supplier
	err := ctx.BodyParser(&modifiedSupplier)
	uuid := ctx.Params("uuid", "")

	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, "Failed to parse request")
	}
	Supplier, err := service.UpdateSuppliersByUuid(uuid, modifiedSupplier)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 400, err.Error())
	}

	return api.SuccessResponse(ctx, "Supplier updated successfully", &Supplier)
}
