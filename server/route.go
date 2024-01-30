package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/webhook/core/auth"
	"github.com/raafly/webhook/utils"
	"github.com/raafly/webhook/utils/constans"
	"gorm.io/gorm"
)

func NewUserRoutes(app *fiber.App, db *gorm.DB) {
	repo := auth.NewAuthRepository(db)
	service := auth.NewAuthService(repo, *utils.NewPassword(), constans.NewValidationError())
	handler := auth.NewUserHandler(service)

	public := app.Group("/public")
	public.Post("/register", handler.Register)
	public.Post("/login", handler.Login)	
}