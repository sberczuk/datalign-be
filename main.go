package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type App struct {
	app *fiber.App
}

func NewApp() *App {
	app := fiber.New()
	a := &App{app: app}
	return a
}

func SetupRoutes(app *fiber.App) *fiber.App {
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/hello", home2)
	app.Post("/eval", eval)
	return app
}

func main() {
	// Initialize a new Fiber app
	app := NewApp()

	app.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	// Define a route for the GET method on the root path '/'
	SetupRoutes(app.app)
	// Start the server on port 3000

	//TODO: This should be configurable
	log.Fatal(app.app.Listen(":3000"))

}

func home2(c fiber.Ctx) error {
	// Send a string response to the client
	return c.SendString("Hello, World WWWW 👋!")
}
