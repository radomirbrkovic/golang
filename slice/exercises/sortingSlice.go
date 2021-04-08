// This is example of using sorting slice of structs 
package main

import (
	"fmt"
	"sort"
)

type employee struct {
	name string
	age int
	salary float32
}

func main()  {
	slice := make([]employee, 0)
	slice = append(slice, employee{"John", 32, 2500})
	slice = append(slice, employee{"Mike", 36, 4900})
	slice = append(slice, employee{"Liza", 26, 3900.92})
	slice = append(slice, employee{"Kim", 49, 5100})


	fmt.Println("Slice:", slice)

	sort.Slice(slice, func (i, j int) bool {
		return slice[i].age > slice[j].age
	})

	fmt.Println("Order by age:", slice)

	sort.Slice(slice, func (i, j int) bool {
		return slice[i].salary > slice[j].salary
	})

	fmt.Println("Order by salary:", slice)

	sort.Slice(slice, func (i, j int) bool {
		return slice[i].name < slice[j].name
	})

	fmt.Println("Order by name:", slice)
}