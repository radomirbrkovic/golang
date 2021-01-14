/*
Struct
A struct is an composite data type. (composite data types, aka, aggregate data types, aka, complex data types). 
Struct allow us to compose together values of different types.

All properties in struct must be initialized 

https://golang.org/ref/spec#Struct_types
*/


package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
}

type basketballPlayer struct {
	person // untyped property (nested structure)
	position string
}

func main()  {
	
	p1 := person {
		firstName: "Kobe",
		lastName: "Brayant",
		age: 42, 
	}

	p2 := person {
		"Micheal",
		"Jordan",
		57,
	}

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)


	fmt.Println("=============")
	/*
	Embedded structures
		We can take one struct and embed it in another struct. 
		When you do this the inner type gets promoted to the outer type.
	*/

	player1 := basketballPlayer{
		person: p1,
		position: "Small forward, Shooting guard",
	}

	player2 := basketballPlayer{
		person: p2,
		position: "Outfielder, Small forward, Shooting guard",
	}

	fmt.Println(player1)
	fmt.Println(player2)

	fmt.Println("=============")

	// Override struct data by other structure 
	player2.person = p1
	fmt.Println(player2)
}