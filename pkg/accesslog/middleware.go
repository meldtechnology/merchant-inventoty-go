package accesslog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meldtechnology/merchant-inventory-go/pkg/log"
	"time"
)

// Handler returns a middleware that records an access log message for every HTTP request being processed.
func Handler(logger log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		ctx := log.WithRequest(c.Context(), c)

		err := c.Next()

		// generate an access log message
		logger.With(ctx, "duration", time.Now().Sub(start).Milliseconds(), "status", c.Response().StatusCode()).
			Infof("%s %s %s %d %d", c.Method(), c.OriginalURL(), c.Protocol(), c.Response().StatusCode(), c.BodyRaw())

		return err
	}
}
