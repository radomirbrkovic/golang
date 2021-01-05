/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/

package main

import "fmt"

 func main()  {
		 bdYear := 1991
		 
		 for {
			if bdYear > 2021 {
				break
			}

			 fmt.Println(bdYear)
			 bdYear++
		 }
 }