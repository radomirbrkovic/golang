/*
Write a program that: 
- assigns an int to a variable
- prints that int in decimal, binary, and hex
- shifts the bits of that int over 1 position to the left, and assigns that to a variable
- prints that variable in decimal, binary, and hex
*/

package main

import "fmt"

func main()  {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n\n", x, x, x)
	y := x << 1 
	fmt.Printf("%d\t%b\t%#x\n\n", y, y, y)
}