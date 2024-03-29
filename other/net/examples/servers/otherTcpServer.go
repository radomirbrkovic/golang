// Other example of tcp server

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	SERVER := "localhost:"+args[1]

	s, err := net.ResloveTCPAddr("tcp", SERVER)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	conn, err := l.Accept()

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting TCP server!")
			conn.Close()
			return
		}

		fmt.Print("> ", string(buffer[0:n-1]))
		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}