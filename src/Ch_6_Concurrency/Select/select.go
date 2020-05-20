package main

import (
	"fmt"
	"time"
)

func main() {

	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		ch1 <- 42
		ch2 <- 54
		ch1 <- 33
	}()
	select {
	case val := <-ch1:
		fmt.Printf("Got %d from ch 1\n ", val)
	case val := <-ch2:
		fmt.Printf("Got %d from ch 2 \n", val)
	}

	// create another channel and func that sleeps for 100 milliseconds
	out := make(chan float64)
	go func() {
		time.Sleep(300 * time.Millisecond)
		out <- 3.14
	}()
	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)
	case <-time.After(200 * time.Millisecond):
		fmt.Printf("Timeout")
	}
}
