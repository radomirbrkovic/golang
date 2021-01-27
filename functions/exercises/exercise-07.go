/*
	Create a func which returns a func
	assign the returned func to a variable
	call the returned func
*/


package main 

import "fmt"

func main()  {
	f := foo()
	f()
}

func foo() func() {
	return func() {
		fmt.Println("Hello from foo - sub function")
	}
}