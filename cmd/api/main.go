package main

import (
	"github.com/gaboliveiradev/api-jobs/internal/config"
	"github.com/gaboliveiradev/api-jobs/internal/router"
)

var (
	logger *config.Logger
)

// @title API Jobs
// @version 1.0
// @description API para gerenciamento de vagas de emprego.
// @termsOfService https://api-jobs.com/terms

// @contact.name Gabriel Oliveira
// @contact.url https://github.com/gaboliveiradev
// @contact.email gabriel@email.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
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
