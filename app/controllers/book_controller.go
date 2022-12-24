package controllers

import (
	"github.com/gofiber/fiber/v2"
	"project/scratch/app/repository"
	"project/scratch/pkg/utils"
)

func GetBooks(c *fiber.Ctx) error {
	return c.Status(200).JSON(utils.ResponseSuccess(repository.GetUsers()))
}
