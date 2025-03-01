package httpv1

import (
	"girls/internal/model"

	"github.com/gofiber/fiber/v2"
)

type IndexResponse struct {
	*model.Page
}

func (r *Router) IndexHandler(c *fiber.Ctx) error {
	page := r.NewPage()
	page.Title = "фор вуманс"
	resp := IndexResponse{
		page,
	}
	return c.Render("index", resp, "layouts/main")

}
