package main

import (
	"fmt"
	"time"
)

func main() {
	// Create channel with no receiver
	ch := make(chan int)
	// //This will block
	// <-ch
	// fmt.Println("here")

	// Create Goroutine and Send value to channel
	go func() {
		//Send number of the channel
		ch <- 353
	}()
	// Receive from channel
	val := <-ch
	fmt.Printf("got %d\n", val)

	//Send multiple
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("Got %d\n", val)
	}

	//close the channel to signal we are done
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("send %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	for i := range ch {
		fmt.Printf("Received: %d\n", i)
	}
}
