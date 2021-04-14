// client code
package main
import "fmt"

func main() {
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)
	
	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
	fmt.Println("\n==============================\n")
	fmt.Println("Remove 28")
	link.Remove(28)
	fmt.Println("Remove 9")
	link.Remove(9)
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
	fmt.Println("\n==============================\n")
	link.Reverse()
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	link.Display()
	fmt.Println("\n==============================\n")
}