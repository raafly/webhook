package ewallet

import (
	"github.com/gofiber/fiber/v2"
	constants "github.com/raafly/webhook/utils/constans"
)

type EwalletHandler interface {
	FindTransactionById(c *fiber.Ctx) error
	ChangeStatus(c *fiber.Ctx) error
}

type ewalletHandlerImpl struct {
	port ewalletService
}

func NewEwalletHandler(port ewalletService) EwalletHandler {
	return &ewalletHandlerImpl{port: port}
} 

func (h *ewalletHandlerImpl) FindTransactionById(c *fiber.Ctx) error {
	param := c.Params("id")
	result, err := h.port.GetPaymentStatus(param)
	if err !=  nil {
		return c.Status(404).JSON(
			constants.NewNotFoundError("id transaction not found"),
		)
	}

	return c.Status(200).JSON(
		constants.NewSuccess("success get payment", result),
	)
}

func (h *ewalletHandlerImpl) ChangeStatus(c *fiber.Ctx) error {
	param := c.Params("id")
	payload := new(payload)
	c.BodyParser(payload)

	err := h.port.ChangeStatus(param, payload.Status.Status)
	if err != nil {
		return c.Status(400).JSON(
			constants.NewBadRequestError(err.Error()),
		)
	}

	return c.Status(200).JSON(fiber.Map{"massage": "success"})
}
