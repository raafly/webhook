package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/webhook/core/ewallet"
	"gorm.io/gorm"
)

func NewEwalletRoutes(app *fiber.App, db *gorm.DB) {
	repoEwallet := ewallet.NewEwalletRepository(db)
	servEwallet := ewallet.NewEwalletService(repoEwallet)
	handlerEwallet := ewallet.NewEwalletHandler(servEwallet)

	ewallet := app.Group("/payment")
	ewallet.Get("/:id", handlerEwallet.FindTransactionById)
}