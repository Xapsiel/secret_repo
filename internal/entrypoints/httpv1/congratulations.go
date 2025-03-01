package httpv1

import (
	"girls/internal/model"

	"github.com/gofiber/fiber/v2"
)

type CongratulationResponse struct {
	*model.Page
	Congratulations []model.Congratulations `json:"congratulations"`
}

func (r *Router) CongratulationsHandler(c *fiber.Ctx) error {
	page := r.NewPage()
	cong, err := r.service.ReadCongratulations()
	if err != nil {
		return err
	}
	resp := CongratulationResponse{
		Page:            page,
		Congratulations: cong,
	}
	return c.Render("congratulations", resp, "layouts/main")

}
