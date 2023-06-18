package interfaces

type ProfileHandler interface {
	ExecutionCollectMemStats(bool)
	GetAllocatedMem() uint64
	ExecutionTimeStart()
	ExecutionTimeEnd()
	GetExecutionTimeInMin() (float64, error)
}
