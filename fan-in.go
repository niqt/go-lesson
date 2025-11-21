package main

import "fmt"

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func fanInWithSelect(input1, input2 <-chan string) <-chan string {
	//var c chan string
	c := make(chan string)
	go func() {
		for {
			select {
			case s1 := <-input1:
				c <- s1
			case s2 := <-input2:
				c <- s2
			}
		}
	}()
	return c
}

func FanInExample() {
	c := fanIn(Boring("Bob"), Boring("Alice"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func FanInWithSelectExample() {
	c := fanInWithSelect(Boring("Bob"), Boring("Alice"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}
