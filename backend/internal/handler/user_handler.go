package handler

import "github.com/gofiber/fiber/v2"

type UserHandler struct{}

func (h *UserHandler) Me(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"error": "invalid user_id",
		})
	}

	return c.JSON(fiber.Map{
		"user_id": userID,
	})
}
