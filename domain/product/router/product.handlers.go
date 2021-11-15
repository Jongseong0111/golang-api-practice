package router

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"tutorial.sqlc.dev/app/domain/product/dto"
	productservice "tutorial.sqlc.dev/app/domain/product/service"
	"tutorial.sqlc.dev/app/util"
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
		return ctx.Status(http.StatusBadRequest).JSON(util.SendError(http.StatusBadRequest, err, ctx))
	}

	newProduct, err := productService.CreateProduct(product)

	if err != nil {
		if err.Error() == "duplicate Product"{
			return ctx.Status(http.StatusConflict).JSON(util.HTTPStatusResponse{
				Status:  http.StatusConflict,
				Message: "duplicate productunit",
			})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(util.SendError(http.StatusInternalServerError, err, ctx))
	}

	return ctx.JSON(newProduct)
}

func GetProductList(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	productList, err := productService.GetProductList(userId)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(util.SendError(http.StatusInternalServerError, err, ctx))
	}

	return ctx.JSON(productList)

}