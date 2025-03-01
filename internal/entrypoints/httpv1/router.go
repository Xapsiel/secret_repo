package httpv1

import (
	"girls/internal/config"
	"girls/internal/model"
	"girls/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	service service.Service
	Domain  string
	Title   string
}

func NewRouter(service service.Service, cfg config.HostConfig) *Router {
	return &Router{
		service: service,
		Domain:  cfg.Domain,
		Title:   cfg.Title,
	}
}
func (r *Router) NewPage() *model.Page {
	return &model.Page{
		Domain: r.Domain,
		Title:  r.Title,
	}
}
func (r *Router) Routes(app fiber.Router) {
	app.Static("assets", "web/assets")
	app.Static("", "web/assets")
	app.Static("static", "web/assets")
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} - ${ua}\\n\n",
	}))

	app.Get("/", r.IndexHandler)
	app.Get("/congratulations", r.CongratulationsHandler)
}
