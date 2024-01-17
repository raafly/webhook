package ewallet

import "github.com/gofiber/fiber/v2"

type EwalletHandler interface {
	EWalletPaymentStatus(c *fiber.Ctx) error
}

type ewalletHandlerImpl struct {
	ewalletService ewalletService
}

func NewEwalletHandler(ewalletService ewalletService) EwalletHandler {
	return &ewalletHandlerImpl{ewalletService}
}

func (h *ewalletHandlerImpl) EWalletPaymentStatus(c *fiber.Ctx) error {
	data := new(ewalletStatusResponse)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	ewallet, err := h.ewalletService.GetPaymentStatus(data.TransactionID)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(ewallet)
}
