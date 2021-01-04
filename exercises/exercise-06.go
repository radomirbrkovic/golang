// Write a program that prints a number in decimal, binary, and hex

package main

import ("fmt")

func main()  {
	a := 42
	fmt.Printf("%d\t%b\t%#x", a, a, a)
}