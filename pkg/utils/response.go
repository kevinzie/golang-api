package utils

import "github.com/gofiber/fiber/v2"

func ResponseSuccess(data any) fiber.Map {
	return fiber.Map{
		"status": "success",
		"data":   data,
	}
}

func ResponseError(data any) fiber.Map {
	return fiber.Map{
		"status": "success",
		"data":   data,
	}
}
