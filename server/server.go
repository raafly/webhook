package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/raafly/webhook/config"
	"github.com/raafly/webhook/auth"
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

func (s *Server) autoMigrate() error {
	db := database.NewPostgres(s.Conf)
	if db.Migrator().HasTable(&auth.User{}) {
		return fmt.Errorf("table already exists ...skipping migrations")
	}

	db.AutoMigrate(&auth.User{})
	return nil
}

func (s *Server) Run() error {
	db := database.NewPostgres(s.Conf)
	_ = s.autoMigrate()
	
	file, err := os.OpenFile(".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	s.App.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		Output: file,
	}))

	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Authorization, Content-Length",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowCredentials: true,
	}))

	NewUserRoutes(s.App, db)

	return s.App.Listen(":3000")
}

func (s *Server) GracefulShutdown(port string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.App.Listen(":" + port); err != nil {
			log.Fatalf("error when listening to :%s, %s", port, err)
		}
	}()

	log.Printf("server is running on :%s", port)

	<-stop

	log.Println("server gracefully shutdown")

	if err := s.App.Shutdown(); err != nil {
		log.Fatalf("error when shutting down the server, %s", err)
	}

	log.Println("process clean up...")
}
