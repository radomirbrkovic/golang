// Example od simple tcp client
package main 

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	
	if len(os.Args) == 1 {
		fmt.Println("Please provide host:post")
		return
	}

	connect := os.Args[1]

	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprint(c, text + "\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP Client exiting ...")
			return 
		}

	}


}