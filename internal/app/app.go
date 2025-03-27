// Package app is responsible of wiring up and initializing all the dependencies of the application and modules.
package app

import (
	"github.com/alfredoprograma/filez-server/internal/config"
	"github.com/alfredoprograma/filez-server/internal/repositories"
	"github.com/alfredoprograma/filez-server/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type appRepos struct {
	UserRepository repositories.UserRepository
}

type appServices struct {
	UserService services.UserService
}

type Application struct {
	Server   *fiber.App
	Config   *config.Config
	Repos    appRepos
	Services appServices
}

func NewApplication(config *config.Config, db *gorm.DB) Application {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	fiber := fiber.New()

	return Application{
		Server: fiber,
		Config: config,
		Repos: appRepos{
			UserRepository: userRepo,
		},
		Services: appServices{
			UserService: userService,
		},
	}
}
