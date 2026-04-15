package main

import (
	"log"

	"github.com/joho/godotenv"

	_ "finora-wealthlab/docs"
	"finora-wealthlab/internal/service"
	"finora-wealthlab/pkg/database"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Finora WealthLab API
// @version 1.0
// @description Personal finance tracker & planner
// @host localhost:8080
// @BasePath /
func main() {
	godotenv.Load()
	// connect DB
	err := database.Connect()
	if err != nil {
		log.Fatal("DB connect error:", err)
	}

	app := fiber.New()

	// routes
	app.Get("/", HealthCheck)
	app.Post("/api/register", RegisterHandler)
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Fatal(app.Listen(":8080"))
}

// =========================
// HANDLERS
// =========================

// HealthCheck godoc
// @Summary Check server status
// @Description Returns API status
// @Tags health
// @Success 200 {string} string "ok"
// @Router / [get]
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("API running 🚀")
}

// Register godoc
// @Summary Register user
// @Description Create new user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body RegisterRequest true "user info"
// @Success 200 {object} map[string]string
// @Router /api/register [post]
func RegisterHandler(c *fiber.Ctx) error {
	var body RegisterRequest

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	userService := service.NewUserService()

	err := userService.Register(body.Email, body.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "register success",
	})
}
