package handler

import (
	"finora-wealthlab/internal/service"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	AuthService *service.AuthService
	UserService *service.UserService
}

type MessageResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Register godoc
// @Summary Register user
// @Description Create new user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body RegisterRequest true "user info"
// @Success 200 {object} MessageResponse
// @Router /api/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err := h.UserService.Register(body.Email, body.Password)
	if err != nil {
		return err
	}

	return c.JSON(MessageResponse{
		Message: "registered",
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param body body LoginRequest true "login info"
// @Success 200 {object} LoginResponse
// @Router /api/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	token, err := h.AuthService.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(LoginResponse{
		Token: token,
	})
}
