package main

import (
	"github.com/kunalvirwal/shogun-cd/internal/app"
	"github.com/kunalvirwal/shogun-cd/internal/pipeline"
	"github.com/kunalvirwal/shogun-cd/internal/utils"
)

func main() {
	initServices()
	// api.StartAPIServer()
	// pipeline.LoadPipeline("./examples/pipeline.yaml")
}

func initServices() {
	// Initialize logger

	logger := utils.NewLogger(utils.DebugLevel, true)
	pipelineService := pipeline.NewPipelineService(logger)

	// Initialize main application
	app := app.NewApp(pipelineService, logger)

	_ = app

}
