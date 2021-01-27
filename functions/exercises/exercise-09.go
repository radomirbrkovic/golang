/*
Closure is when we have “enclosed” the scope of a variable in some code block. 
For this hands-on exercise, create a func which “encloses” the scope of a variable:
*/

package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println("f", f())
	fmt.Println("f", f())
	g := foo()
	fmt.Println("g", g())
	fmt.Println("g", g())
	fmt.Println("g", g())
	fmt.Println("f", f())
	fmt.Println("g", g())
	fmt.Println("f", f())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}