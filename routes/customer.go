package routes

import (
	"github.com/gofiber/fiber/v2"
	"merchant_inventory/api"
	"merchant_inventory/model"
	"merchant_inventory/repository"
	"merchant_inventory/service"
)

func CreateCustomer(ctx *fiber.Ctx) error {
	var newCustomer model.Customer
	err := ctx.BodyParser(&newCustomer)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 400, "Failed to parse request")
	}

	Customer, err := service.AddCustomer(newCustomer)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Customer added successfully", &Customer)
}

func GetCustomers(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("size", 10)

	pageable, err := service.GetCustomers(repository.Pageable{
		Page: page, Limit: size,
	})
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Customer sRetrieved successfully", &pageable)
}

func GetCustomerByPhone(ctx *fiber.Ctx) error {
	phone := ctx.Params("phone", "")

	Customer, err := service.GetCustomersByPhone(phone)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Customer Retrieved successfully", &Customer)
}

func GetCustomerByUuid(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid", "")

	Customer, err := service.GetCustomersByUuid(uuid)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, err.Error())
	}

	return api.SuccessResponse(ctx, "Customer Retrieved successfully", &Customer)
}

func UpdateCustomer(ctx *fiber.Ctx) error {
	var modifiedCustomer model.Customer
	err := ctx.BodyParser(&modifiedCustomer)
	phone := ctx.Params("phone", "")

	if err != nil {
		return api.ErrorHandlerResponse(ctx, 404, "Failed to parse request")
	}
	Customer, err := service.UpdateCustomersByUuid(phone, modifiedCustomer)
	if err != nil {
		return api.ErrorHandlerResponse(ctx, 400, err.Error())
	}

	return api.SuccessResponse(ctx, "Customer updated successfully", &Customer)
}
