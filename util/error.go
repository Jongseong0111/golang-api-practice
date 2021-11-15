package util

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

type HTTPStatusResponse struct {
	Status int `json:"status"`
	Message string `json:"msg"`
}

func SendError(errType int, err error, ctx *fiber.Ctx) HTTPStatusResponse {
	fmt.Fprintln(os.Stderr, err.Error())
	return HTTPStatusResponse{
		Status: errType,
		Message: err.Error(),
	}
}
