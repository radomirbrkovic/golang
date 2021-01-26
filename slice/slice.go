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
	*/

	x = append(x, 43, 44, 88, 21, 46)
	fmt.Println("=============")
	fmt.Println(x)
	y := []int{100, 101, 108, 128}
	x = append(x, y...)
	fmt.Println(x)

	/*
	Slice - slicing a slice
	*/
	fmt.Println("=============")
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:5])

	/*
	Slice - deleting from a slice
	*/

	fmt.Println("=============")

	// Remove 2nd and 3rd element from array
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	/*
		Slice - make
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
	*/
	fmt.Println("=============")
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
