package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/webhook/utils/constans"
)

type handlerUser interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
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

	err := h.port.insertOne(data)
	if err != nil {
		return c.Status(400).JSON(constans.NewBadRequestError(err.Error()))
	}

	return c.Status(201).JSON(constans.NewCreated("success create account"))
}

func (h *handlerUserImpl) Login(c *fiber.Ctx) error {
	data := new(login)	
	c.BodyParser(data)

	respon, err := h.port.login(data)
	if err != nil {
		return c.Status(404).JSON(constans.NewBadRequestError(err.Error()))
	}

	c.Cookie(&fiber.Cookie{
		Name: "access_token",
		Value: respon.AccessToken,
		MaxAge: int(respon.AccessTokenExpired),
		SameSite: "disable",
		Domain: "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name: "refresh_token",
		Value: respon.RefreshToken,
		MaxAge: int(respon.RefreshTokenExpired),
		SameSite: "disable",
		Domain: "localhost",
	})

	c.Set("Authorizated", respon.AccessToken)

	return c.Status(200).JSON(response{
		Code: 200,
		Status: "OK",
		Message: "success login",
		Data: payload{
			UserID: respon.UserID,
			Token: respon.AccessToken,
		},
	})
}