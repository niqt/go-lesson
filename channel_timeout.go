package main

import (
	"fmt"
	"time"
)

func FinishAfterFixedTime() {
	c := Boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}

func FinishIfSlow() {
	c := Boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case s := <-time.After(800 * time.Millisecond):
			fmt.Println("You're too slow." + s.String())
			return
		}
	}
}
