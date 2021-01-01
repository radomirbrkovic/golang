package main 

import ("fmt")
var a int = 47
var b float64 = 4.7 

func main() {
	fmt.Println(a)
	fmt.Printf("Type of variable 'a' is :%T\n", a)
	fmt.Printf("Binary representation of variable 'a' is :%b\n", a)
	fmt.Printf("Hexadecimal representation of variable 'a' is: %x\n", a)

	fmt.Println(b)
	fmt.Printf("Type of variable 'b' is :%T\n", b)

	fmt.Printf("%v + %v = %v\n", float64(a), b, float64(a) + b) 
	fmt.Printf("%v * %v = %v\n", float64(a), b, float64(a) * b)
	fmt.Printf("%v - %v = %v\n", float64(a), b, float64(a) - b) 
	fmt.Printf("%v / %v = %v\n", float64(a), b, float64(a) / b)  
	fmt.Printf("%v mod %v = %v \n\n", a, int(b), a % int(b)) 
}