package runtime_profiler

import (
	"errors"
	"runtime"
	"time"
)

type MemorySizeUnit string

const (
	MemInMB MemorySizeUnit = "MB"
	MemInKB MemorySizeUnit = "KB"
)

type RuntimeProfileHandler struct {
	MemStatsCollector runtime.MemStats
	MemSizeUnit       MemorySizeUnit
	StartTime         time.Time
	EndTime           time.Time
}

func NewRuntimeProfileHandler(f runtime.MemStats, sizeUnit MemorySizeUnit) (*RuntimeProfileHandler, error) {
	return &RuntimeProfileHandler{
		MemStatsCollector: f,
		MemSizeUnit:       sizeUnit,
	}, nil
}

// ExecutionCollectMemStats function should be called to record current mem stats
func (r *RuntimeProfileHandler) ExecutionCollectMemStats(runGC bool) {
	// Init garbage collection for up-to-date stats at end of compute
	if runGC {
		runtime.GC()
	}
	runtime.ReadMemStats(&r.MemStatsCollector)
}

// GetExecutionMemInGB function fetches collected mem in GB
func (r *RuntimeProfileHandler) GetAllocatedMem() uint64 {
	val := r.byteToHumanReadable(r.MemStatsCollector.HeapAlloc)
	return val
}

// ExecutionTimeStart marks start of execution time
func (r *RuntimeProfileHandler) ExecutionTimeStart() {
	r.StartTime = time.Now()
}

// ExecutionTimeEnd marks end of execution time
func (r *RuntimeProfileHandler) ExecutionTimeEnd() {
	r.EndTime = time.Now()
}

// GetExecutionTimeInMin is used to fetch execution time
func (r *RuntimeProfileHandler) GetExecutionTimeInMin() (float64, error) {
	if r.StartTime.After(r.EndTime) {
		return 0, errors.New("starTime can't be greater than endTime")
	}
	return r.EndTime.Sub(r.StartTime).Minutes(), nil
}

// byteToGB to convert bytes to MBs
func (r *RuntimeProfileHandler) byteToHumanReadable(b uint64) uint64 {
	var humanReadableSize uint64
	switch r.MemSizeUnit {
	case MemInMB:
		humanReadableSize = b / 1024 / 1024
	case MemInKB:
		humanReadableSize = b / 1024
	}
	return humanReadableSize
}
