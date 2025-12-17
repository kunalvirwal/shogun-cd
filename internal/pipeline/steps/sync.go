package pipelineSteps

// ExecSteps are only valid for target type server

type SyncStep struct {
	TriggerWhen string       `yaml:"trigger_when,omitempty"`
	Target      string       `yaml:"target"` // [TODO]: change this to pointer if needed
	Files       []FileUpdate `yaml:"files"`
}

type FileUpdate struct {
	Dst string `yaml:"dst"`
	Src string `yaml:"source"`
}

func (*SyncStep) Type() string {
	return SyncType
}
