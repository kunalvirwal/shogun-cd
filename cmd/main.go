package main

import "github.com/kunalvirwal/shogun-cd/internal/pipeline"

func main() {
	// api.StartAPIServer()
	pipeline.LoadPipeline("./examples/pipeline.yaml")
}
