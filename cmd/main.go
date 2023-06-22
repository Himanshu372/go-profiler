package main

import (
	"fmt"
	"github.com/Himanshu372/go-profiler/runtime_profiler"
	"math/rand"
	"runtime"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Init profiler")
	var memstats runtime.MemStats
	var startHeapAlloc, endHeapAlloc uint64
	profiler, err := runtime_profiler.NewRuntimeProfileHandler(memstats, runtime_profiler.MemInKB)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Starting Collection of Stats")
	profiler.ExecutionTimeStart()
	fmt.Println("Collecting Memstats at start with garbage collection")
	// Collecting mem stats without garbage collection
	profiler.ExecutionCollectMemStats(true)
	fmt.Println("Starting time")
	// Fetching HeapAlloc from mem stats
	startHeapAlloc = profiler.GetAllocatedMem()

	fmt.Println("Start compute/service")
	memArray := make([]string, 0)
	for i := 0; i < 5; i++ {
		fmt.Println(fmt.Sprintf("Inside %d loop", i))
		j := 10
		for {
			if j == i {
				break
			}
			memArray = append(memArray, string(letters[0:rand.Intn(len(letters))]))
			j -= 1
		}
	}
	fmt.Println("End compute/service")

	fmt.Println("Ending collection of Memstats without garbage collection")
	// Collecting mem stats without garbage collection
	profiler.ExecutionCollectMemStats(false)
	// Fetching HeapAlloc from mem stats
	endHeapAlloc = profiler.GetAllocatedMem()

	fmt.Println("Ending time")
	// Fetching HeapAlloc from mem stats
	// End executionTime timer
	profiler.ExecutionTimeEnd()
	elapsedTime, err := profiler.GetExecutionTimeInMin()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Time elapsed during compute: %f min", elapsedTime))
	fmt.Println(fmt.Sprintf("Heap usage during compute: %d KB", endHeapAlloc-startHeapAlloc))

}
