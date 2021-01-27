/*
- create a type SQUARE
- create a type CIRCLE
- attach a method to each that calculates AREA and returns it
	- circle area= π r 2
	- square area = L * W
- create a type SHAPE that defines an interface as anything that has the AREA method
- create a func INFO which takes type shape and then prints the area
- create a value of type square
- create a value of type circle
- use func info to print the area of square
- use func info to print the area of circle
*/

package main

import (
	"fmt"
	"math"
	)

type circle struct {
	radius float64 
}	

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type sharpe interface {
	area() float64
}

func info(s sharpe) {
	fmt.Println(s.area())
}

func main() {
	c := circle {
		radius: 10.34,
	}

	s := square {
		length: 5,
	}
	info(c)
	info(s)
}




