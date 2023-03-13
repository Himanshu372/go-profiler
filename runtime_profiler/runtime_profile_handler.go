package runtime_profiler

import (
	"errors"
	"runtime"
	"time"
)

type RuntimeProfileHandler struct {
	memStatsCollector runtime.MemStats
	startTime         time.Time
	endTime           time.Time
}

func NewRuntimeProfileHandler(f runtime.MemStats) (*RuntimeProfileHandler, error) {
	return &RuntimeProfileHandler{
		memStatsCollector: f,
	}, nil
}

// ExecutionCollectMemStats function should be called to record current mem stats
func (r *RuntimeProfileHandler) ExecutionCollectMemStats(runGC bool) {
	// Init garbage collection for up-to-date stats at end of compute
	if runGC {
		runtime.GC()
	}
	runtime.ReadMemStats(&r.memStatsCollector)
}

// GetExecutionMemInGB function fetches collected mem in GB
func (r *RuntimeProfileHandler) GetExecutionMemInMB() uint64 {
	return byteToMB(r.memStatsCollector.HeapAlloc)
}

// ExecutionTimeStart marks start of execution time
func (r *RuntimeProfileHandler) ExecutionTimeStart() {
	r.startTime = time.Now()
}

// ExecutionTimeEnd marks end of execution time
func (r *RuntimeProfileHandler) ExecutionTimeEnd() {
	r.endTime = time.Now()
}

// GetExecutionTimeInMin is used to fetch execution time
func (r *RuntimeProfileHandler) GetExecutionTimeInMin() (float64, error) {
	if r.startTime.After(r.endTime) {
		return 0, errors.New("starTime can't be greater than endTime")
	}
	return r.endTime.Sub(r.startTime).Minutes(), nil
}

// byteToGB to convert bytes to MBs
func byteToMB(b uint64) uint64 {
	return b / 1024 / 1024
}
