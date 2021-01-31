/*
 - Using goroutines, create an incrementer program
	 - have a variable to hold the incrementer value
	 - launch a bunch of goroutines
			- each goroutine should
				- read the incrementer value
					- store it in a new variable
				- yield the processor with runtime.Gosched()
				- increment the new variable
				- write the value in the new variable back to the incrementer variable
 - use waitgroups to wait for all of your goroutines to finish
 - the above will create a race condition. 
 - Prove that it is a race condition by using the -race flag
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}