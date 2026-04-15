package handler

import (
	"finora-wealthlab/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err := h.Service.Register(body.Email, body.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "registered"})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	ok, err := h.Service.Login(body.Email, body.Password)
	if err != nil || !ok {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	return c.JSON(fiber.Map{"message": "login success"})
}
