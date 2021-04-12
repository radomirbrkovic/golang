// Create program which will print invalid part(s) of ip address from given string
package main 

import (
	"bufio"
	"fmt"
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
	fmt.Println("Please enter any text and check is it valid IPv4 address. If you want to finish just type STOP keyword.")

	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			fmt.Println("Have a nice day ;)")
			break
		} else {
			input = strings.Split(scanner.Text(), ".")

			if(len(input) != 4) {
				fmt.Println("IP address must have 4 integer parts. This is invalid FORMAT")
			} else {

				passed := 0
				for  _, i := range input {

					ipPart, err := strconv.Atoi(i)

					if err != nil {
						fmt.Printf("%v is invalid part of IP \n", i)
						break
					}

					if (ipPart < 0 || ipPart > 255) {
						fmt.Println(ipPart, " is out of IP range")
						break
					}

					passed ++
				}

				if passed == len(input) {
					fmt.Println(scanner.Text(), " is VALID IP address!")
				}
			}
		}	
	}
}

