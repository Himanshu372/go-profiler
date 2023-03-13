package runtime_profiler

import (
	"runtime"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Runtime Profile Handler Tests", func() {

	var (
		runtimeProfiler *RuntimeProfileHandler
		m               runtime.MemStats
		err             error
	)
	It("successfully returns the runtime profile handler", func() {
		runtimeProfiler, err = NewRuntimeProfileHandler(m)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(runtimeProfiler).ShouldNot(BeNil())
	})

	BeforeEach(func() {
		runtimeProfiler, err = NewRuntimeProfileHandler(m)
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("successfully executed ExecutionCollectMemStats without garbage collection", func() {
		runtimeProfiler.ExecutionCollectMemStats(false)
		// Compute something
		kv := make(map[int]bool)
		for i := 0; i < 10; i++ {
			kv[i] = true
		}
		runtimeProfiler.ExecutionCollectMemStats(false)
		memBytes := runtimeProfiler.GetExecutionMemInMB()
		Ω(memBytes).ShouldNot(BeZero())
	})

	It("successfully executed ExecutionCollectMemStats with garbage collection", func() {
		runtimeProfiler.ExecutionCollectMemStats(false)
		// Compute something
		kv := make(map[int]bool)
		for i := 0; i < 10; i++ {
			kv[i] = true
		}
		runtimeProfiler.ExecutionCollectMemStats(true)
		memBytes := runtimeProfiler.GetExecutionMemInMB()
		Ω(memBytes).ShouldNot(BeZero())
	})

	It("successfully execute GetExecutionTimeInMin", func() {
		// Executing func in a specific go routine
		go func() {
			runtimeProfiler.ExecutionTimeStart()
			d := 5 * time.Second
			Ω(err).ShouldNot(HaveOccurred())
			// Sleep current routine for 5s
			time.Sleep(d)
			runtimeProfiler.ExecutionTimeEnd()
			executionTime, err := runtimeProfiler.GetExecutionTimeInMin()
			Ω(err).ShouldNot(HaveOccurred())
			// Execution time should be > 5s/.12m
			Ω(executionTime).Should(BeNumerically(">", 0.12))
		}()
	})
})
