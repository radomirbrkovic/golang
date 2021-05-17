// Example of using select instruction for fibonacci numbers

package main 

import (
	"fmt"
	"os"
	"strconv"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case c <- x:
				x, y = y, x + y
			case <-quit:
				fmt.Println("quit")
				return	
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	n := 10

	if len(os.Args) == 2 {
		n, _ = strconv.Atoi(os.Args[1])
	}


	go func (n int)  {
		for i:= 0; i < n; i++ {
			fmt.Printf("%d - %v \n", i+1, <-c)
		}
		quit <- 0
	}(n)

	fibonacci(c, quit)

}