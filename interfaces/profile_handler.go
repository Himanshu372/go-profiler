package interfaces

type ProfileHandler interface {
	ExecutionCollectMemStats(bool)
	GetExecutionMemInMB() uint64
	ExecutionTimeStart()
	ExecutionTimeEnd()
	GetExecutionTimeInMin() (float64, error)
}
