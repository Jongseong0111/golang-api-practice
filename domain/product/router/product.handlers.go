package router

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"tutorial.sqlc.dev/app/domain/product/dto"
	productservice "tutorial.sqlc.dev/app/domain/product/service"
)

var (
	productService = productservice.ProductService{}
)

func MappingUrl(app *fiber.App) {
	app.Post("/product", CreateProduct)
	app.Get("/product/:userid", GetProductList)
}

func CreateProduct(ctx *fiber.Ctx) error {
	var product dto.CreateProductRequest
	err := ctx.BodyParser(&product);
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	newProduct, err := productService.CreateProduct(product)

	if err != nil {
		if err.Error() == "duplicate Product"{
			return ctx.SendStatus(fiber.StatusConflict)
		}

		if err.Error() == "1"{
			return ctx.SendStatus(fiber.StatusForbidden)
		}
		if err.Error() == "2"{
			return ctx.SendStatus(fiber.StatusAccepted)
		}
		if err.Error() == "3"{
			return ctx.SendStatus(fiber.StatusContinue)
		}
		if err.Error() == "4"{
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(newProduct)
}

func GetProductList(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	productList, err := productService.GetProductList(int32(userId))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(productList)

}