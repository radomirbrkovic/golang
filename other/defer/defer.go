// Example of using defer keyword 
package main

import "fmt"

func df1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
	fmt.Printf("\n------\n")
}

func df2()  {
	for i:= 3; i > 0; i-- {
		defer func() {
			defer fmt.Print(i, " ")	
		}()
	}
	fmt.Printf("\n------\n")
}


func df3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	fmt.Printf("\n------\n")
}

func main()  {
	df1()
	df2()
	df3()
	fmt.Println()
}
