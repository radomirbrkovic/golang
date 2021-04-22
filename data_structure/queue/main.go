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
	fmt.Println("Please enter keywords PUSH or POP for adding or removing element from queue . If you want to finish just type STOP keyword.")

	queue := new(Queue)

	elem := 1
	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			fmt.Println("Have a nice day ;)")
			break
		} else {

			if scanner.Text() == "PUSH" {
				queue.Push(elem)
				elem ++
			} else if scanner.Text() == "POP" {
				queue.Pop()
			}

			fmt.Printf("> %v\n", queue.Items)
		}
		
	}
}
