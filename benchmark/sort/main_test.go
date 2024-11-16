package main

// Run with GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s
// Run with GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s

import (
	"math/rand"
	"runtime"
	"testing"
)

var numbers = []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0} //generateList(1e5)

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)
	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers)
	}
	return numbers
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSortConcurrent(runtime.NumCPU(), numbers)
	}
}
