package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // Load env automatically
	"os"
	"project/scratch/pkg/config"
	"project/scratch/pkg/routes"
	"project/scratch/pkg/utils"
)

func main() {
	fiberConfig := config.FiberConfig()
	app := fiber.New(fiberConfig)

	//database.Config

	err := config.Connect()
	if err != nil {
		return
	}
	routes.PublicRoutes(app)

	if os.Getenv("APP_ENV") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
