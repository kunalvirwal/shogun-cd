package app

import (
	"github.com/kunalvirwal/shogun-cd/internal/pipeline"
	"github.com/kunalvirwal/shogun-cd/internal/utils"
)

type App struct {
	Logger utils.Logger

	PipelineService pipeline.PipelineService
}

func NewApp(pipelineService pipeline.PipelineService, logger utils.Logger) *App {
	return &App{
		Logger:          logger,
		PipelineService: pipelineService,
	}
}
