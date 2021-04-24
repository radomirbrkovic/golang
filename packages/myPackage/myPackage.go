package myPackage

import (
	"fmt"
)

func FunctionA()  {
	fmt.Println("functionA() from myPackage is called!")
}

func FunctionB()  {
	fmt.Println("functionB() from myPackage is called!")
}

// every constant, variable, or function which start with small character 
// will be privare 
func privateFunction() {
	fmt.Println("privateFunction() from myPackage is called!")
}