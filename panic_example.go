package main

import "fmt"

func PanicExampleWithoutRecover() {
	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 2")
	defer fmt.Println("Deferred statement 3")
	c := make(chan int)
	close(c)
	c <- 42
}

func PanicExampleWithRecover() {
	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 2")
	defer fmt.Println("Deferred statement 3")
	defer func() {
		if r := recover(); r != nil {
			// handle the panic
			fmt.Println("Recovered from panic:", r)
		}
	}()
	c := make(chan int)

	close(c)
	c <- 42
}
