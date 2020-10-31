package api

import "github.com/gofiber/fiber/v2"

func SetupMoviesRoutes(app *fiber.App) {
	s := start()
	grp := app.Group("/movies")
	grp.Get("/", s.SearchMovieHandler)
}
