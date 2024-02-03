package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raaafly/powerup-client-service-golang/auth"
	"github.com/raaafly/powerup-client-service-golang/utils"
	"github.com/raaafly/powerup-client-service-golang/utils/constans"
	"gorm.io/gorm"
)

func NewUserRoutes(app *fiber.App, db *gorm.DB) {
	repo := auth.NewAuthRepository(db)
	service := auth.NewAuthService(repo, *utils.NewPassword(), constans.NewValidationError())
	handler := auth.NewUserHandler(service)

	public := app.Group("/auth")
	public.Post("/register", handler.Register)
	public.Post("/login", handler.Login)
	public.Post("/forget-password", handler.ForgetPassword)
	public.Get("/reset-password/", handler.ResetPassword)
}
