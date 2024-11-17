# go-lesson

## This repository contains examples about the use of the goroutine, channel, mutex and panic

### Instructions
By building the main file in the root directory, you will create the application that allows you to run the various examples.
In the benchmark directory, there are additional examples that you need to run as follows:
* GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s (to use 8 cpu)
* GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s (to use 1 cpu)
