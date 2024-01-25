package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/webhook/core/auth"
	"github.com/raafly/webhook/utils"
	"gorm.io/gorm"
)

func NewUserRoutes(app *fiber.App, db *gorm.DB) {
	repo := auth.NewAuthRepository(db)
	service := auth.NewAuthService(repo, *utils.NewPassword())
	handler := auth.NewUserHandler(service)

	auth := app.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)	
}