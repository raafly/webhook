package ewallet

import (
	"github.com/gofiber/fiber/v2"
	constants "github.com/raafly/webhook/constans"
)

type EwalletHandler interface {
	FindTransactionById(c *fiber.Ctx) error
}

type ewalletHandlerImpl struct {
	ewalletService ewalletService
}

func NewEwalletHandler(ewalletService ewalletService) EwalletHandler {
	return &ewalletHandlerImpl{ewalletService}
}

func (h *ewalletHandlerImpl) FindTransactionById(c *fiber.Ctx) error {
	param := c.Params("id")
	result, err := h.ewalletService.GetPaymentStatus(param)
	if err !=  nil {
		return c.Status(404).JSON(
			constants.NewNotFoundError("id transaction not found"),
		)
	}

	return c.Status(200).JSON(
		constants.NewSuccess("success get payment", result),
	)
}