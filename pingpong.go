package main

import (
	"fmt"
	"sync"
)

func pingPong() {
	ping := make(chan string)
	pong := make(chan string)

	var wg sync.WaitGroup

	// Pinger
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ping <- "ping"
			msg := <-pong
			fmt.Println("Pinger received:", msg)
		}
		close(ping)
	}()

	// Ponger
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range ping {
			fmt.Println("Ponger received:", msg)
			pong <- "pong"
		}
	}()

	wg.Wait()
	fmt.Println("Game over!")
}
