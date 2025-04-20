package routes

import (
	"github.com/gofiber/fiber/v2"
	"merchant_inventory/api"
	"merchant_inventory/model"
	"merchant_inventory/repository"
	"merchant_inventory/service"
)

func CreateProduct(ctx *fiber.Ctx) error {
	var newProduct model.Product
	err := ctx.BodyParser(&newProduct)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 400, "Failed to parse request")
	}

	product, err := service.AddProduct(newProduct)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Product added successfully", &product)
}

func GetProducts(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	pageable, err := service.GetProducts(repository.Pageable{
		Page: page, Limit: size,
	})
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Product sRetrieved successfully", &pageable)
}

func GetProductBySku(ctx *fiber.Ctx) error {
	sku := ctx.Params("sku", "")

	product, err := service.GetProductsBySku(sku)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Product Retrieved successfully", &product)
}

func GetProductByUuid(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid", "")

	product, err := service.GetProductsByUuid(uuid)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Product Retrieved successfully", &product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	var modifiedProduct model.Product
	err := ctx.BodyParser(&modifiedProduct)
	sku := ctx.Params("sku", "")

	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, "Failed to parse request")
	}
	product, err := service.UpdateProductsByUuid(sku, modifiedProduct)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 400, err.Error())
	}

	return api.SuccessResponse(ctx, "Product updated successfully", &product)
}
