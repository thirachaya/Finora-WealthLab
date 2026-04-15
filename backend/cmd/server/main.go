package main

import (
	"log"

	"github.com/joho/godotenv"

	_ "finora-wealthlab/docs"
	"finora-wealthlab/internal/handler"
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
	// load env
	godotenv.Load()

	// connect DB
	err := database.Connect()
	if err != nil {
		log.Fatal("DB connect error:", err)
	}

	app := fiber.New()

	// =========================
	// INIT HANDLERS
	// =========================
	authHandler := &handler.AuthHandler{
		AuthService: service.NewAuthService(),
		UserService: service.NewUserService(),
	}
	userHandler := &handler.UserHandler{}

	// =========================
	// ROUTES
	// =========================
	app.Get("/", HealthCheck)
	app.Get("/api/me", userHandler.Me)

	app.Post("/api/register", authHandler.Register)
	app.Post("/api/login", authHandler.Login)

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
func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("API running")
}
