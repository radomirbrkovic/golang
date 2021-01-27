/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/

package main

import "fmt"

func main()  {
	defer foo()
	fmt.Println("main")
}

func foo() {
	defer func(){
		fmt.Println("Anonym function")
	}()

	fmt.Println("foo")	
}