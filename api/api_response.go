package api

import "github.com/gofiber/fiber/v2"

type AppResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}

func Response(status bool, message string, data interface{}) AppResponse {
	return AppResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func SuccessResponse(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.Status(200).JSON(Response(true, message, data))
}

func ErrorHandlerResponse(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(Response(false, message, nil))
}
