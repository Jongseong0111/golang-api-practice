package user

import (
	"github.com/gofiber/fiber/v2"
	"tutorial.sqlc.dev/app/domain/user/router"
)

func Init(app *fiber.App) {
	router.MappingUrl(app)
}