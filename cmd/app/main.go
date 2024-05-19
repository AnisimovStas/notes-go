package main

import (
	"log"
	"log/slog"
	"notes-go/internal/app"
	"notes-go/internal/config"
	"os"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()
	_ = cfg

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	app, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Error("%w", err)
	}

	app.Run()
}
