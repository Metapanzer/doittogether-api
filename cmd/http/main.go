package main

import (
	"DoItTogether/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	viperConfig := config.LoadConfig()

	// Connect to MySQL
	DBConn := config.NewDatabase(viperConfig)
	app := gin.Default()
	cfg := config.AppConfig{
		DB:     DBConn,
		App:    app,
		Config: viperConfig,
	}

	cfg.Init()

	err := app.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
