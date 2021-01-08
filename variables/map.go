/*
MAP

A map is a key-value store. This means that you store some value and you access that value by a key.
For instance, I might store the value “phoneNumber” and access it be the key “friendName”.
This way I could enter my friend’s name and have the map return their phone number. 
The syntax for a map is map[key]value. The key can be of any type which allows comparison 
(eg, I could use a string or an int, for example, as I can compare if two strings are equal,
or if two ints are equal). It is important to note that maps are unordered. 
If you print out all of the keys and values in a map, they will print out in random order. 
The comma ok idiom is also covered in this video. 
A map is the perfect data structure when you need to look up data fast.
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
	You delete an entry from a map using delete(<map name>, “key”). 
	No error is thrown if you use a key which does not exist. 
	To confirm you delete a key/value, verify that the key/value exists with the comma ok idiom.
	*/

	fmt.Println("=============")
	delete(m, "not exist key")
	delete(m, "todd")
	
	for k, v := range m {
		fmt.Println(k, v)
	}
}