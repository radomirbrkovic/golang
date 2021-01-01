package main

import "fmt"

var sGlobal string

func main()  {
	fmt.Printf("The value of `sGlobal` variable is '%v'. So it has \"Zero Value\"\n", sGlobal)
	sGlobal = "New value"

	fmt.Printf("The new value of `sGlobal` variable is '%v'.\n", sGlobal)
	sLocal := "My local value"
	
	fmt.Printf("The value of `sLocal` variable is '%v'.\n", sLocal)

	for i := 0; i < 10; i ++ {
		// Concatenation
		sLocal += "!"
	}

	fmt.Printf("The new value of `sLocal` variable is '%v'.\n\n\n", sLocal)

	// Reference
	myRef := &sLocal

	fmt.Printf("The new value of `myRef` variable is '%v'.\n", myRef)

	*myRef = "New text example"

	fmt.Printf("The new value of `myRef` variable is '%v'.\n", myRef)
	fmt.Printf("The value of `sLocal` variable is '%v'.\n", sLocal)

}