package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"tutorial.sqlc.dev/app/db"
	"tutorial.sqlc.dev/app/domain/user"
	userService "tutorial.sqlc.dev/app/domain/user/service"
)

func main() {
	InitDb()

	app := fiber.New()

	db.Connect()
	defer db.Close()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Go Server")
	})

	user.Init(app)
	app.Listen(":3000")
}

func InitDb() {
	userService.Init()
}