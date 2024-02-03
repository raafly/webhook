package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raaafly/powerup-client-service-golang/utils/constans"
)

type handlerUser interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ForgetPassword(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
}

type handlerUserImpl struct {
	port authService
}

func NewUserHandler(port authService) handlerUser {
	return &handlerUserImpl{port: port}
}

func (h *handlerUserImpl) Register(c *fiber.Ctx) error {
	data := new(register)
	c.BodyParser(data)

	token, err := h.port.insertOne(data)
	if err != nil {
		return c.Status(400).JSON(constans.NewBadRequestError(err.Error()))
	}

	return c.Status(201).JSON(
		constans.NewCreated("success create account", token),
	)
}

func (h *handlerUserImpl) Login(c *fiber.Ctx) error {
	req := new(login)
	c.BodyParser(req)

	token, err := h.port.login(req)
	if err != nil {
		return c.Status(404).JSON(constans.NewBadRequestError(err.Error()))
	}

	c.Set("Authorizated", token)

	return c.Status(200).JSON(response{
		Code:    200,
		Status:  true,
		Message: "success login",
		Data: data{
			Token: token,
		},
	})
}

func (h *handlerUserImpl) ForgetPassword(c *fiber.Ctx) error {
	req := new(forgetPassword)
	c.BodyParser(req)

	err := h.port.forgetPassword(req.Email)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"code":    404,
			"message": "could find the email",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"code":    200,
		"message": "Email send Successfully",
	})
}

func (h *handlerUserImpl) ResetPassword(c *fiber.Ctx) error {
	token := c.Query("token")

	if token == "" {
		return c.Status(402).JSON(fiber.Map{
			"status": false,
			"code": 402,
			"message": "Token is invalid",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"code": 402,
		"message": "Token is valid",
	})
}
