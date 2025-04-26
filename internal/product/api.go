package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrchantinevntory/internal/errors"
	"github.com/mrchantinevntory/pkg/log"
	"github.com/mrchantinevntory/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r fiber.Router, service Service, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/products/<id>", res.get)
	r.Get("/products", res.query)

	//TODO: Will be implemented in revision 3
	//r.Use(authHandler)

	// the following endpoints require a valid JWT
	r.Post("/products", res.create)
	r.Put("/products/<id>", res.update)
	r.Delete("/products/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *fiber.Ctx) error {
	product, err := r.service.Get(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}

	return c.JSON(product)
}

func (r resource) query(c *fiber.Ctx) error {
	ctx := c.Context()
	params := ctx.QueryArgs()
	page, _ := params.GetUint("page")
	size, _ := params.GetUint("size")

	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c, count)
	products, err := r.service.Query(ctx, page, size)
	if err != nil {
		return err
	}
	pages.Rows = products
	return c.JSON(pages)
}

func (r resource) create(c *fiber.Ctx) error {
	var input CreateProductRequest
	if err := c.BodyParser(&input); err != nil {
		r.logger.With(c.Context()).Info(err)
		return errors.BadRequest("")
	}
	product, err := r.service.Create(c.Context(), input)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(product)
}

func (r resource) update(c *fiber.Ctx) error {
	var input UpdateProductRequest
	if err := c.BodyParser(&input); err != nil {
		r.logger.With(c.Context()).Info(err)
		return errors.BadRequest("")
	}

	product, err := r.service.Update(c.Context(), c.Params("id"), input)
	if err != nil {
		return err
	}

	return c.JSON(product)
}

func (r resource) delete(c *fiber.Ctx) error {
	product, err := r.service.Delete(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}

	return c.JSON(product)
}
