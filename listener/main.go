package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateServer() *fiber.App {
	app := fiber.New()
	return app
}

func main() {
	app := CreateServer()

	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Println("[INFO] Server started on :5002")
	log.Fatal(app.Listen(":5002"))
}