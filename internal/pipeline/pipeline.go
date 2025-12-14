package pipeline

type Kind string

const PipelineKind Kind = "Pipeline"

type Pipeline struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       Kind     `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Triggers []Trigger     `yaml:"triggers"`
	Steps    []StepWrapper `yaml:"steps"`
}

type Trigger struct {
	Type  string   `yaml:"type"`
	Paths []string `yaml:"paths,omitempty"`
}
