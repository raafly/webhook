package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/webhook/core/auth"
	"github.com/raafly/webhook/utils"
	"gorm.io/gorm"
)

func NewUserRoutes(app *fiber.App, db *gorm.DB) {
	repo := auth.NewAuthRepository(db)
	service := auth.NewAuthService(repo, *utils.NewPassword(), &validator.Validate{})
	handler := auth.NewUserHandler(service)

	auth := app.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)	
}