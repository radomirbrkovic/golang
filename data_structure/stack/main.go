package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println("Please enter keywords PUSH or POP for adding or removing element from stack . If you want to finish just type STOP keyword.")

	stack := new(Stack)


	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			fmt.Println("Have a nice day ;)")
			break
		} else {

			if scanner.Text() == "PUSH" {
				stack.Push(len(stack.Items) + 1)
			} else if scanner.Text() == "POP" {
				stack.Pop()
			}

			fmt.Printf("> %v\n", stack.Items)
		}
		
	}
}
