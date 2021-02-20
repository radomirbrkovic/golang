package main

import "fmt"

func main() {
	// Set up the pipline
	c := gen(2,3,4,5,6,7,8,9)
	out := sq(c)
	for i := range out {
		fmt.Println(i)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func ()  {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	
	out := make(chan int)

	go func ()  {
		for n:= range in {
			out <- n * n
		}
		close(out)
	}()
	return out

}

