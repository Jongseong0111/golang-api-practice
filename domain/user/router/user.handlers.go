package router

import (
	"github.com/gofiber/fiber/v2"
	"tutorial.sqlc.dev/app/domain/user/dto"
	userservice "tutorial.sqlc.dev/app/domain/user/service"
)

var (
	userService = userservice.UserService{}
)
func MappingUrl(app *fiber.App) {
	app.Post("/user", CreateUser)
}

func CreateUser(ctx *fiber.Ctx) error {
	var user dto.User
	err := ctx.BodyParser(&user);
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	newUser, err := userService.CreateUser(user)

	if err != nil {
		return ctx.SendStatus(fiber.StatusConflict)
	}

	return ctx.JSON(newUser)
}