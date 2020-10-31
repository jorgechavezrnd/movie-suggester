package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jorgechavezrnd/movies-suggester/api"
)

func main() {
	app := fiber.New()
	api.SetupMoviesRoutes(app)
	_ = app.Listen(":3001")
}
