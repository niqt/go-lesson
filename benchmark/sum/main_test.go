package main

// Run with GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s
// Run with GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s

import (
	"runtime"
	"testing"
)

var numbers = generateList(1e7)

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}
