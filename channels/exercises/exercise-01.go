/*
	Create channel with using func literal, aka, anonymous self-executing func
*/
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
	
}