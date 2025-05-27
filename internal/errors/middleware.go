package errors

import (
	"database/sql"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/meldtechnology/merchant-inventory-go/pkg/log"
	"runtime/debug"
	"strconv"
)

// Handler creates a middleware that handles panics and errors encountered during HTTP request processing.
func Handler(logger log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			l := logger.With(c.Context())
			if e := recover(); e != nil {
				var ok bool
				if err, ok = e.(error); !ok {
					err = fmt.Errorf("%v", e)
				}

				l.Errorf("recovered from panic (%v): %s", err, debug.Stack())
			}

			if err != nil {
				res := buildErrorResponse(err)
				if res.StatusCode() == fiber.StatusInternalServerError {
					l.Errorf("encountered internal server error: %v", err)
				}
				c.Set("X-Custom-Header", strconv.Itoa(res.StatusCode()))
				if err != nil {
					l.Errorf("failed writing error response: %v", err)
					return
				}
				// skip any pending handlers since an error has occurred
				err := abort(c)
				if err != nil {
					return
				}
				err = nil // return nil because the error is already handled
			}
		}()
		return c.Next()
	}
}

// buildErrorResponse builds an error response from an error.
func buildErrorResponse(err error) ErrorResponse {
	switch err.(type) {
	case ErrorResponse:
		return err.(ErrorResponse)
	case validation.Errors:
		return InvalidInput(err.(validation.Errors))
	case fiber.ConversionError:
	default:
		return ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		}

	}

	if errors.Is(err, sql.ErrNoRows) {
		return NotFound("")
	}
	return InternalServerError("")
}

func abort(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}
