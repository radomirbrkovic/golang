package main

import "fmt"

func main() {
	fmt.Println("2 + 3=", mySum(2, 3))
	fmt.Println("4 + 3=", mySum(4, 3))
	fmt.Println("5 + 12=", mySum(5, 12))
}

func mySum(x ...int) int  {
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}