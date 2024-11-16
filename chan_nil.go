package main

func ChanNil() {
	c := make(chan int)

	close(c)
	c = nil
	c <- 42
}
