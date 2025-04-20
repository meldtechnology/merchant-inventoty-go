package config

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CreatAppServer() *fiber.App {
	return fiber.New()
}

func StartServer(app *fiber.App, port int) {
	err := app.Listen(":" + strconv.Itoa(port))
	if err != nil {
		panic("Could not start server ")
	}

}
