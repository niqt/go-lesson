package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

// CPUIntensiveTask performs a CPU-intensive calculation
func CPUIntensiveTask(start, end int, wg *sync.WaitGroup, results chan<- float64) {
	defer wg.Done()

	var sum float64
	// Perform complex mathematical operations
	for i := start; i < end; i++ {
		// Simulate heavy CPU computation with mathematical operations
		for j := 0; j < 1000; j++ {
			sum += math.Sqrt(float64(i*j)) * math.Sin(float64(i)) * math.Cos(float64(j))
			sum += math.Pow(float64(i), math.Mod(float64(j), 5))
			x := math.Log(float64(i + 1))
			sum += x * math.Exp(math.Min(x, 10))
		}
	}
	results <- sum
}

func CPU() {
	// Use all available CPU cores
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Running on %d CPUs\n", numCPU)

	// Parameters for the workload
	totalNumbers := 1000000
	numbersPerGoroutine := totalNumbers / numCPU

	// Create channels and wait group
	results := make(chan float64, numCPU)
	var wg sync.WaitGroup

	// Record start time
	startTime := time.Now()

	// Launch goroutines
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		start := i * numbersPerGoroutine
		end := start + numbersPerGoroutine
		if i == numCPU-1 {
			// Make sure we process all remaining numbers in the last goroutine
			end = totalNumbers
		}
		go CPUIntensiveTask(start, end, &wg, results)
	}

	// Start a goroutine to close results channel after all work is done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and sum all results
	var finalSum float64
	for result := range results {
		finalSum += result
	}

	// Calculate execution time
	duration := time.Since(startTime)

	fmt.Printf("Total sum: %f\n", finalSum)
	fmt.Printf("Time taken: %v\n", duration)
	fmt.Printf("Average time per goroutine: %v\n", duration/time.Duration(numCPU))
}
