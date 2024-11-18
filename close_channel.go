package main

import (
	"fmt"
	"sync"
	"time"
)

func TerminateCallingClose() {
	var wg sync.WaitGroup
	wg.Add(1)

	strChan := make(chan string)
	go func() {
		for {
			element, ok := <-strChan
			// check if channel is closed
			if !ok {
				println("Goroutines killed!")
				wg.Done()
				return
			}
			println(element)
		}
	}()
	strChan <- "this"
	strChan <- "is"
	strChan <- "a"
	strChan <- "message"
	close(strChan)

	// wait all goroutine to stop
	wg.Wait()
	// print the last message
	fmt.Println("This is the end of TerminateCallingClose func!")
}

func TerminateWithChannel() {
	fmt.Println("ExampleWithChannel")

	quitChan := make(chan bool)
	go func() {
		for {
			select {
			case <-quitChan:
				return
			default:
				// print a message every 3 seconds
				fmt.Println("Test goroutine")
				time.Sleep(time.Second * 3)
			}
		}
	}()

	// sleep to print some message from the goroutine
	time.Sleep(time.Second * 10)

	// stop the goroutine
	quitChan <- true
	fmt.Println("corountine stopped!")

	// test if the gorountine stopped or not
	time.Sleep(time.Second * 10)
	fmt.Println("End of the TerminateWithChannel")
}
