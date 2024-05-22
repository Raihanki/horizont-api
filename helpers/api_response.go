package helpers

import "github.com/gofiber/fiber/v2"

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ValidatrionError struct {
	Key     string      `json:"key"`
	Tag     string      `json:"tag"`
	Message interface{} `json:"message"`
}

func ApiResponse(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	if data == nil {
		data = []interface{}{}
	}
	return ctx.Status(statusCode).JSON(DefaultResponse{
		Message: message,
		Data:    data,
	})
}
