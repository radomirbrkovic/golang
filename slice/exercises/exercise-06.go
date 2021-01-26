/*
Create a slice to store the names of all of the states in the United States of America. 
 - What is the length of your slice? 
 - What is the capacity?  
 - Print out all of the values, along with their index position in the slice, 
  without using the range clause. Here is a list of the states:

"Alabama", 
" Alaska", 
" Arizona", 
" Arkansas", 
" California", 
" Colorado", 
" Connecticut", 
" Delaware", 
" Florida", 
" Georgia", 
" Hawaii", 
" Idaho", 
" Illinois",
" Indiana", 
" Iowa",
" Kansas",
" Kentucky", 
" Louisiana", 
" Maine", 
" Maryland", 
" Massachusetts", 
" Michigan", 
" Minnesota", 
" Mississippi", 
" Missouri", 
" Montana", 
" Nebraska", 
" Nevada", 
" New Hampshire", 
" New Jersey", 
" New Mexico", 
" New York", 
" North Carolina", 
" North Dakota", 
" Ohio", 
" Oklahoma", 
" Oregon",
" Pennsylvania", 
" Rhode Island", 
" South Carolina", 
" South Dakota", 
" Tennessee", 
" Texas", 
" Utah", 
" Vermont", 
" Virginia",
" Washington", 
" West Virginia", 
" Wisconsin", 
" Wyoming", 
*/


package main

import "fmt"

func main()  {
	x := make([]string, 50, 50)
	x = []string{"Alabama", 
	" Alaska", 
	" Arizona", 
	" Arkansas", 
	" California", 
	" Colorado", 
	" Connecticut", 
	" Delaware", 
	" Florida", 
	" Georgia", 
	" Hawaii", 
	" Idaho", 
	" Illinois",
	" Indiana", 
	" Iowa",
	" Kansas",
	" Kentucky", 
	" Louisiana", 
	" Maine", 
	" Maryland", 
	" Massachusetts", 
	" Michigan", 
	" Minnesota", 
	" Mississippi", 
	" Missouri", 
	" Montana", 
	" Nebraska", 
	" Nevada", 
	" New Hampshire", 
	" New Jersey", 
	" New Mexico", 
	" New York", 
	" North Carolina", 
	" North Dakota", 
	" Ohio", 
	" Oklahoma", 
	" Oregon",
	" Pennsylvania", 
	" Rhode Island", 
	" South Carolina", 
	" South Dakota", 
	" Tennessee", 
	" Texas", 
	" Utah", 
	" Vermont", 
	" Virginia",
	" Washington", 
	" West Virginia", 
	" Wisconsin", 
	" Wyoming", }

	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
