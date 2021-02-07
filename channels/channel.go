package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	fmt.Println("=========")			
	// Using Buffer
	/*
	The capacity, in number of elements, sets the size of the buffer in the channel. 
	If the capacity is zero or absent, the channel is unbuffered and communication 
	succeeds only when both a sender and receiver are ready.
	*/
	b := make(chan int, 2)
	go func() {
		for i := 0; i < 5; i++ {
				b <- i
		}
		
		close(b)
	}()

		// receive
	for v := range b {
			fmt.Println(v)
	}	

	// send & receive (bidirectional)
	fmt.Println("=========")

	cc := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Printf("cc\t%T\n", cc)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

}
