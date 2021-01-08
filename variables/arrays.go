/*
	Array
Arrays are mostly used as a building block in the Go programming language. 
In some instances, we might use an array, but mostly the recommendation is to use slices instead.
https://golang.org/ref/spec#Array_types

DECALARATION:
	var x [length]<PRIMITIVE TYPE>
	x := [length]<PRIMITIVE TYPE>

	AFTER DECLARATION SIZE OF ARRAY CAN'T BE CHANGE BECAUSE IT WILL MAKE ANOTHER ARRAY
*/

package main

import "fmt"

var x [3]int 

func main()  {
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	x[0] = 1
	x[2] = 3
	fmt.Println(x)

	fmt.Println("length:" , len(x))
	fmt.Println("capacity:" , cap(x))
}