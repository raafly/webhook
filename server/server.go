package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/webhook/config"
	// "github.com/raafly/webhook/core/auth"
	"github.com/raafly/webhook/database"
)

type Server struct {
	App  *fiber.App
	Conf *config.AppConfig
}

func NewServer() *Server {
	app := fiber.New()
	conf := config.NewAppConfig()

	return &Server{
		App:  app,
		Conf: conf,
	}
}

func (s *Server) Run() error {
	db := database.NewPostgres(s.Conf)
	// db.AutoMigrate(&auth.User{})

	NewUserRoutes(s.App, db)

	return s.App.Listen(":3000")
}