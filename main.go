package main

import (
	"GoFiber/config"
	"GoFiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize database
	config.ConnectDatabase()

	// Set up routes
	routes.PostRoutes(app)
	routes.UserRoutes(app)

	// Start server
	app.Listen(":3000")
}
