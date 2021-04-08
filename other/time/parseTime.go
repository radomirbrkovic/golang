// This is example of parsing and manipulating time from string in golang from standard input file
package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main()  {
	var input []string
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println("Please enter any time in format HH:MM and add how many minutes you want to add. If you want to finish just type STOP keyword.")

	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			fmt.Println("Have a nice day ;)")
			break
		} else {

			input = strings.Split(scanner.Text(), " ")
			if(len(input) > 1) {
				inputTime, err := time.Parse("15:04", input[0])
				if err == nil {
					addMinute := 0
					if(len(input) == 2) {
						addMinute, _= strconv.Atoi(input[1])
					}

					newTime := inputTime.Add(time.Minute * time.Duration(addMinute))
					fmt.Println("Full:", newTime)
					fmt.Println("Time:", newTime.Hour(), newTime.Minute())
					
				} else {
					fmt.Println(err)
				}
			}
			
		}
	}

}