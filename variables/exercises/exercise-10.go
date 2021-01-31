// Create a variable of type string using a "raw string literal". Print it.

package main

import "fmt"

func main()  {
	a := `
		This is 
		a 
		"raw string"
	`

	fmt.Println(a)
}