package api

import "github.com/gofiber/fiber/v2"

func (w *WebServices) SearchMovieHandler(c *fiber.Ctx) error {
	res, err := w.s.search.Search(&MovieFilter{})

	if err != nil {
		return fiber.NewError(400, "cannot bring movies")
	}

	return c.JSON(res)
}
