// Example of Atomic Counters

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main()  {
	// We’ll use an unsigned integer to represent our (always-positive) counter.
	var ops unit64

	// A WaitGroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup

	// We’ll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i ++ {
		wg.Add(1)

		go func ()  {
			for c := 0; c < 1000; c ++ {
				atomic.AddUnit64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Wait until all the goroutines are done.
	wg.Wait()
	
	/*
		It’s safe to access ops now because we know no other goroutine is writing to it. 
		Reading atomics safely while they are being updated is also possible, 
		using functions like atomic.LoadUint64.
	*/
	
	fmt.Println("ops:", ops)


}