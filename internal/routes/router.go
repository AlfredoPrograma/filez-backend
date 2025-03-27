package routes

import (
	"github.com/alfredoprograma/filez-server/internal/app"
	"github.com/alfredoprograma/filez-server/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *app.Application) {
	api := app.Server.Group("/api/v1")

	userRoutes(app, api)
}

func userRoutes(app *app.Application, parent fiber.Router) {
	users := parent.Group("/users")
	userController := controllers.NewUserController(app.Services.UserService)

	users.Get(":email", userController.GetByEmail)
	users.Post("", userController.Create)
}
