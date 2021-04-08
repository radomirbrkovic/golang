// This is example of parsing date from string in golang from standard input file
package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
)

func main()  {
	var inputDate string

	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println("Please enter any date in format Day/Mont/Year. If you want to finish just type STOP keyword.")


	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			fmt.Println("Have a nice day ;)")
			break
		} else {
			inputDate = scanner.Text()

			date, err := time.Parse("02/01/2006", inputDate)
			if err == nil {
				fmt.Println("Full:", date)
				fmt.Println("Time:", date.Day(), date.Month(), date.Year(), date.Weekday())
			} else {
				fmt.Println(err)
			}
		}
		
	}
}