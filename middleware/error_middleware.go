package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/webhook/utils/constans"
)

func ErrorMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			c.Status(fiber.StatusInternalServerError).JSON(
				constans.NewInternalServerError("INTERNAL SERVER ERROR"),
			)
		}
	}()

	return c.Next()
}