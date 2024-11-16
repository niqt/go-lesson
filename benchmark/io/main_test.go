package main

// Run with GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s
// Run with GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s

import (
	"runtime"
	"testing"
)

var docs = generateList(1e3)

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find("test", docs)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "test", docs)
	}
}
