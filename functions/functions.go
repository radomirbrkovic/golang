package main

import "fmt"

func main() {
	defer foo4()
	foo()
	foo1("Some example")
	fmt.Println("Result from foo2 function is: ", foo2())
	x, y :=foo3()
	fmt.Println("Result from foo3 function is: ",x, y )
	a := multiplication(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Result from multiplication function is: ",a)

	b := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println("Result from multiplication function is: ", multiplication(b...))

	s1 := scientists {
		person: person {
			firstName: "John",
			lastName: "Nash",
		},
		category: "Mathematic",
	}

	s2 := scientists {
		person: person {
			firstName: "Albert",
			lastName: "Einstein",
		},
		category: "Physic",
	}

	s1.speak()
	s2.speak()
}

// Function without parameters and returns
func foo() {
	fmt.Println("Hello world, I am a foo function")
}

// Function with parameter string
func foo1(s string) {
	fmt.Println("Hello from second function I have arguments and you pass me: ", s)
}

// Function with returns
func foo2() string {
	return "RETURN"
}

// Function with multiple return 
func foo3() (string, int) {
	return "Test", 42
}


/*
Variadic parameter
*/

func multiplication(x ...int) int {
	result := 1
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	for _, v := range x {
		result *= v
	}
	
	return result
}

/*
Defer
*/

func foo4() {
	fmt.Println("This is defer function and it will be the last executed independent where in code it will be called")
}

/*
Methods
*/

type person struct {
	firstName string
	lastName string
}

type scientists struct {
	person
	category string
}

func (s scientists) speak() {
	fmt.Println("My name is ", s.firstName, s.lastName, " and I work in ", s.category, " field")
}