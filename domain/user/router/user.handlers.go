package router

import (
	"github.com/gofiber/fiber/v2"
	"tutorial.sqlc.dev/app/domain/user/dto"
	userservice "tutorial.sqlc.dev/app/domain/user/service"
	"tutorial.sqlc.dev/app/model"
)

var (
	userService = userservice.UserService{}
)
func MappingUrl(app *fiber.App) {
	app.Post("/user", CreateUser)
	app.Get("/user", GetUserList)
	app.Put("/user", UpdateUser)
}

func CreateUser(ctx *fiber.Ctx) error {
	var user dto.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	newUser, err := userService.CreateUser(user)

	if err != nil {
		return ctx.SendStatus(fiber.StatusConflict)
	}

	return ctx.JSON(newUser)
}

func GetUserList(ctx *fiber.Ctx) error {

	users, err := userService.GetUserList()
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(users)
}

func UpdateUser(ctx *fiber.Ctx) error {
	var req dto.UpdateBodyInfo
	err :=ctx.BodyParser(&req)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	currentUser, err := userService.GetSignedID(req.QuestionAccount)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = userService.UpdateUserAccount(model.UpdateUserParams{UserName: req.UpdateUserName, UserID: currentUser})
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	type Response struct {
		message string `json:"status"`
	}

	return ctx.JSON(Response{message: "Update Success"})
}