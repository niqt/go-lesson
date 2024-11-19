# go-lesson

## This repository contains examples about the use of the goroutine, channel, mutex and panic

### Install go
The official Go installer are here https://go.dev/dl/

### Instructions examples
In the main there are the application to run 16 examples, to run: make run

### Instructions benchmarks
In the benchmark directory, there are additional examples that you need to run as follows:
* GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s (to use 8 cpu)
* GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s (to use 1 cpu)


For contact: nicola.defilippo at gmail.com