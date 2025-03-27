package controllers

import (
	"net/http"

	"github.com/alfredoprograma/filez-server/internal/domain"
	"github.com/alfredoprograma/filez-server/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) GetByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")

	user, err := c.userService.GetByEmail(email)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"data": user,
	})
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	var payload domain.CreateUserDTO

	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if err := c.userService.Create(payload); err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusCreated)
}
