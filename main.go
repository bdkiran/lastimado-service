package main

import (
	"log"

	"github.com/bdkiran/lastimado-service/api"
	"github.com/bdkiran/lastimado-service/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting the server")
	app := fiber.New()
	db.SetupConnection()
	api.InitializeRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
