// reading from standard input (console)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println("Please enter any text. If you want to finish just type STOP keyword.")

	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			fmt.Println("Have a nice day ;)")
			break
		}
		fmt.Println(">", scanner.Text())
	}
	
}