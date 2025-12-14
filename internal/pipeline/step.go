package pipeline

const (
	MutateType = "mutate"
	SyncType   = "sync"
	ExecType   = "exec"
	ApplyType  = "apply"
)

type StepWrapper struct {
	Step Step
}
type Step interface {
	Type() string
}
