package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/webhook/constans"
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
		return constans.NewBadRequestError(err.Error())
	}

	return constans.NewCreated("success create acccount")
}

func (h *handlerUserImpl) Login(c *fiber.Ctx) error {
	data := new(login)	
	c.BodyParser(data)

	respon, err := h.port.login(data)
	if err != nil {
		return constans.NewBadRequestError(err.Error())
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"
	cookie.Value = respon.AccessToken
	// cookie.Expires =  time.Now().Add(time.Hour * 24).Unix()

	return constans.NewSuccess("success login", nil)
}