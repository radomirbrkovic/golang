package main

import "fmt"

// The custom type can be decalred via: type <name> <PRIMITIVE TYPE>
// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS

type myType int
var x myType

func main() {
	x = 42
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	
	// converting certain type to integer
	y := int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
}