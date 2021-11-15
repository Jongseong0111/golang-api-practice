package product

import (
	"github.com/gofiber/fiber/v2"
	"tutorial.sqlc.dev/app/domain/product/router"
)

func Init(app *fiber.App) {
	router.MappingUrl(app)
}