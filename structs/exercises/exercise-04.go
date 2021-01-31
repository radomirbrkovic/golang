/*
	Create and use an anonymous struct.
*/
package main 

import "fmt"

func main()  {
	
	s := struct {
		firstName string
		lastName string
		companies map[string]string
	}{
		firstName: "Steve",
		lastName: "Jobs",
		companies: map[string]string{
			"Apple": "Founder",
			"Next": "Founder",
			"Pixar": "Buyer",
		},
	}

	fmt.Println("First name:", s.firstName)
	fmt.Println("Last name:", s.lastName)
	fmt.Printf("\nCompany\tStatus\n")
	for i,v := range s.companies {
		fmt.Println("-----------------------")
		fmt.Printf("%s\t%s\n", i, v)
	}
}