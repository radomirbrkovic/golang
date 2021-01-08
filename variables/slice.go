/*
	A SLICE HOLDS VALUES OF THE SAME TYPE. 
	
	If I wanted to store all of my favorite numbers,
	 I would use a slice of int. If I wanted to store all of my favorite foods, 
	 I would use a slice of string. We will use a COMPOSITE LITERAL to create a slice.
	  A composite literal is created by having the TYPE followed by CURLY BRACES 
	  and then putting the appropriate values in the curly brace area.

	  https://golang.org/ref/spec#Slice_types
*/

package main 

import "fmt"

func main()  {

	// COMPOSITE LITERAL
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println("=============")
	/*
	We can loop over the values in a slice with the range clause. 
	We can also access items in a slice by index position
	*/
	for _, v:= range x {
		fmt.Println(v)
	}
	fmt.Println("=============")
	for i:= 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	/*
	APPEND
	To append values to a slice, we use the special built in function append.
	 This function returns a slice of the same type. 
	*/

	x = append(x, 43, 44, 88, 21, 46)
	fmt.Println("=============")
	fmt.Println(x)
	y := []int{100, 101, 108, 128}
	x = append(x, y...)
	fmt.Println(x)

	/*
	Slice - slicing a slice
	We can slice a slice which means that we can cut parts of the slice away.
	We do this with the colon operator.
	*/
	fmt.Println("=============")
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:5])

	/*
	Slice - deleting from a slice
	We can delete from a slice using both append and slicing. 
	This is a wonderful and elegant example of why Go is great and how Go provides ease of programming.
	*/

	fmt.Println("=============")

	// Remove 2nd and 3rd element from array
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	/*
		Slice - make
		Slices are built on top of arrays. A slice is dynamic in that it will grow in size. 
		The underlying array, however, does not grow in size. 
		When we create a slice, we can use the built in function make to specify 
		how large our slice should be and also how large the underlying array should be. 
		This can enhance performance a little bit.
			make([]T, length, capacity) 
			make([]int, 50, 100) 
	*/
	
	fmt.Println("=============")
	z :=  make([]int, 10, 12)
	fmt.Println(z)
	z[0] = 10
	z[6] = 11
	z[9] = 23
	fmt.Println(z)
	fmt.Println("length:", len(z))
	fmt.Println("capacity:", cap(z))
	z = append(z, 21, 22, 25, 24)
	fmt.Println(z)
	fmt.Println("length:", len(z))
	fmt.Println("capacity:", cap(z))

	/*
	Slice - multi-dimensional slice
	A multi-dimensional slice is like a spreadsheet. You can have a slice of a slice of some type.
	Does that sound confusing? Watch this video and it will all be clarified.
	*/
	fmt.Println("=============")
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
