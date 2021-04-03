// Create program for calculating average number of all arguments from comand line
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	arguments := os.Args
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more numbers.")
		os.Exit(1)
	}

	var sum float64 
	count := 0

	for i:= 1; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64) 
	
		if n > 0 {
			sum += n	
			count ++
		}
		
	}

	fmt.Println("Average number is:", sum / float64(count))

}