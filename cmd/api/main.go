package main

import (
	"github.com/gaboliveiradev/api-jobs/internal/config"
	"github.com/gaboliveiradev/api-jobs/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("Main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Failed initialize config: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
