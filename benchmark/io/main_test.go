package main

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
