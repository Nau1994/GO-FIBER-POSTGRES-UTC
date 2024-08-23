package routes

import (
	"GoFiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	userGroup.Post("/", controllers.CreateUser)
	userGroup.Get("/", controllers.GetAllUsers)
}
