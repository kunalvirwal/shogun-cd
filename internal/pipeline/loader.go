package pipeline

import (
	"os"

	"go.yaml.in/yaml/v3"
)

func (p *Service) LoadPipeline(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		p.logger.LogNewError("failed to read pipeline file: %v", err)
		return
	}

	var pipeline Pipeline
	if err := yaml.Unmarshal(data, &pipeline); err != nil {
		p.logger.LogNewError("failed to unmarshal pipeline YAML: %v", err)
		return
	}

	// [TODO]: Validate the pipeline structure here if needed

	for i, sw := range pipeline.Spec.Steps {
		p.logger.LogInfo("Step %d: Type=%s, Details=%+v", i+1, sw.Step.Type(), sw.Step)
	}
	p.logger.Log("Pipeline %v loaded successfully", pipeline.Metadata.Name)

	// [TODO]: Further processing of the loaded pipeline, store into DB and all
}
