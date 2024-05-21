package helpers

import "github.com/gofiber/fiber/v2"

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ValidatrionError struct {
	Key     string `json:"key"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

func ApiResponse(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return ctx.Status(statusCode).JSON(DefaultResponse{
		Message: message,
		Data:    data,
	})
}
