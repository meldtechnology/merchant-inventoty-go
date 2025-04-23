package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrchantinevntory/pkg/adding"
	"github.com/mrchantinevntory/pkg/listing"
)

func Handler(a adding.Service, l listing.Service) *fiber.App {
	router := fiber.New()

	router.Get("/api/v1/products", getProducts(l))
	router.Get("/api/v1/products/:id", getProduct(l))
	router.Get("/api/v1/products/:sku", getProductSku(l))
	router.Post("/api/v1/products", addProduct(a))

	return router
}

// addProduct returns a handler for POST /products requests
func addProduct(s adding.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var newPrd adding.Product
		if err := c.BodyParser(&newPrd); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		err := s.AddProduct(newPrd)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		// error handling omitted for simplicity
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "New Product added.",
		})
	}
}

// getProducts returns a handler for GET /products requests
func getProducts(l listing.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := c.QueryInt("page", 1)
		lm := c.QueryInt("limit", 10)

		pg, err := l.GetProducts(listing.Pageable{Page: p, Limit: lm})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&pg)
	}
}

// getProduct returns a handler for GET /product by id requests
func getProduct(l listing.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id", -1)

		ps, err := l.GetProduct(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&ps)
	}
}

// getProduct returns a handler for GET /product by id requests
func getProductSku(l listing.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("sku", "")

		ps, err := l.GetProductSku(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&ps)
	}
}
