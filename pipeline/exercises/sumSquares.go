/*
	Create pipeline for summing square of int numbers in given range 
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

var DATA = make(map[int]bool)

var signal chan struct{}

func square(n int) int {
	return n * n
}

func sum(in <-chan int) {
	sum := 0
	for x := range in {
		sum += x
	}
	fmt.Printf("The sum of the squares is %d.\n", sum)
}

func first(min, max int, out chan<- int) {
	
	for i := min; i <= max; i ++ {
		out <- square(i)	
	}

	close(out)

}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		_, ok := DATA[x]
		if ok {
			signal <- struct{}{}
		} else {
			fmt.Print(x, " ")
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		return
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d.\n", n1, n2)
		return
	}

	signal = make(chan struct{})
	A := make(chan int)
	B := make(chan int)

	
	go first(n1, n2, A)
	go second(B, A)

	sum(B)
	
}
