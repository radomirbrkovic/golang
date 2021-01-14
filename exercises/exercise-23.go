/*
	Create your own type “person” which will have an underlying type of “struct” so that 
	it can store the following data:
		- first name
		- last name
		- favorite ice cream flavors
	Create two VALUES of TYPE person. Print out the values, 
	ranging over the elements in the slice which stores the favorite flavors. 

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

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	fmt.Println("Favorite ice cream flavors are:")
	for _, v := range p1.iceCream {
		fmt.Println("\t - ",v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	fmt.Println("Favorite ice cream flavors are:")
	for _, v := range p2.iceCream {
		fmt.Println("\t - ",v)
	}

}