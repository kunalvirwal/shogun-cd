package pipelineSteps

type ApplyStep struct {
	TriggerWhen string `yaml:"trigger_when,omitempty"`
	Target      string `yaml:"target"` // [TODO]: change this to pointer if needed
	Files       []string
}

func (*ApplyStep) Type() string {
	return ApplyType
}
