// Print only integers from standard input file until client send "END" keyword
// and calculate sum of those numbers
package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main()  {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println("Please enter whole number. If you want to finish just type END keyword.")
	var sum int64 

	for scanner.Scan() {
		
		if scanner.Text() == "END" {
			fmt.Println("Sum is:", sum)
			break
		}

		s, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		sum += s

		fmt.Println(">", s)
	}
	
	
}