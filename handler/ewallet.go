package handler

import "github.com/gofiber/fiber/v2"

type WebhookHandler interface {
	eWalletPaymentStatus(c *fiber.Ctx) error
}

type webhookHandler struct {
}

func NewWebhookHandler() WebhookHandler {
	return &webhookHandler{}
}

func (h *webhookHandler) eWalletPaymentStatus(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
