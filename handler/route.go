package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/webhook/core/ewallet"
	"gorm.io/gorm"
)

func NewWebhookRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}

func NewEwalletRoutes(app *fiber.App, db *gorm.DB) {
	repoEwallet := ewallet.NewEwalletRepository(db)
	servEwallet := ewallet.NewEwalletService(repoEwallet)
	handlerEwallet := ewallet.NewEwalletHandler(servEwallet)

	app.Post("/ewallet/payment-status", handlerEwallet.EWalletPaymentStatus)
}
