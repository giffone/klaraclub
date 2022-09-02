package main

import (
	"fmt"
	"time"
)

// Generator returns a channel that produces the numbers 1, 2, 3,…
// To stop the underlying goroutine, send a number on this channel.
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func some() {
	fmt.Println("some")
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	time.Sleep(100 * time.Millisecond)
	number <- 0 // stops underlying goroutine
	// fmt.Println(<-number) // error, no one is sending anymore
	// …
}
