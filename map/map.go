/*
MAP
*/

package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}

	/*
	Here is how we add an element to a map. 
	I also show you how to use the range loop to print out a map’s keys and values.
	*/

	fmt.Println("=============")
	m["todd"] = 33
	for k, v := range m {
		fmt.Println(k, v)
	}

	/*
	Deleting element from map
	*/

	fmt.Println("=============")
	delete(m, "not exist key")
	delete(m, "todd")
	
	for k, v := range m {
		fmt.Println(k, v)
	}
}