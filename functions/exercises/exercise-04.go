/*
Create a user defined struct with 
	- the identifier “person”
	- the fields:
		- first
		- last
		- age
attach a method to type person with
	- the identifier “speak”
	- the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/

package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main()  {
	p1 := person {
		first: "Nikola",
		last: "Tesla",
		age: 86,
	}

	p2 := person {
		first: "Mihajlo",
		last: "Pupin",
		age: 76,
	}

	p1.speak()
	p2.speak()
}

func (p person) speak()  {
	fmt.Println("My name is ", p.first, p.last, " and I'm ", p.age)
}
