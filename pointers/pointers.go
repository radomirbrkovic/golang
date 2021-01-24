/*
Pointers

All values are stored in memory. Every location in memory has an address.
A pointer is a memory address.
*/

package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)

	fmt.Println("=============")
	var c = &a
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", &a)

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int


}
