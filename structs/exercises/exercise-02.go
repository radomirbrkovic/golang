/*
	Take the code from the previous exercise, then store the values of type person in a map with 
	the key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/

package main 

import "fmt"

type person struct {
	firstName string
	lastName string
	iceCream []string
}



func main()  {
	
	p1 := person {
		firstName: "John",
		lastName: "Nash",
		iceCream: []string{"Chocolate", "Blueberries"},
	}

	p2 := person {
		firstName: "Albert",
		lastName: "Einstein",
		iceCream: []string{"Vanilia", "Blackberry"},
	}

	m1 := map[string]person {
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m1 {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		fmt.Println("Favorite ice cream flavors are:")
		for _, v := range v.iceCream {
			fmt.Println("\t - ",v)
		}

		fmt.Println("-------")
	}
	

}