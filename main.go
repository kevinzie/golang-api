package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"project/scratch/pkg/configs"
	"project/scratch/pkg/utils"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	if os.Getenv("APP_ENV") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
