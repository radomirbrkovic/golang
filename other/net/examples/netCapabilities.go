// Example of getting capabilities of all net interfaces on UNIX os
package main 

import (
	"fmt"
	"net"
)

func main()  {
	
	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Println(err)
		return
	} 

	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Println("Interface Flags: ", i.Flags.String())
		fmt.Println("Interface MTU: ", i.MTU)
		fmt.Println("Interface Hardware Address:", i.HardwareAddr)
		fmt.Println("====================")
	}
}
