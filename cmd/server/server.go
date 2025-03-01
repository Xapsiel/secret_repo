package main

import (
	"flag"
	"log/slog"
	"os"

	"girls/internal/config"
	"girls/internal/entrypoints/httpv1"
	"girls/internal/repository"
	"girls/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func InitLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}

func main() {
	logger := InitLogger()
	slog.SetDefault(logger)
	configPath := flag.String("c", "config/config.yaml", "The path to the configuration file")
	flag.Parse()
	cfg, err := config.New(*configPath)
	if err != nil {
		slog.Error("unable to read config:", err.Error())
		os.Exit(1)
	}
	db, err := repository.NewPostgresDB(cfg.DatabaseConfig)
	if err != nil {
		slog.Error("unable to connect to database:", err.Error())
		os.Exit(1)
	}
	slog.Info("Initializing db")
	repos := repository.NewRepository(db)
	slog.Info("Initializing repos")
	services := service.NewService(*repos)

	engine := html.New("./web/views", ".gohtml")
	engine.Reload(true)
	router := httpv1.NewRouter(*services, cfg.HostConfig)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.Routes(app)

	slog.Error(app.Listen(":" + cfg.HostConfig.Port).Error())
}
