// Linked list 
package main

import "fmt"

type List struct {
	head *Node
	tail *Node
} 


func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key: key,
	}

	if L.head != nil {
		L.head.prev = list
	}

	L.head = list
	head := L.head

	for head.next != nil {
		head = head.next
	}

	L.tail = head

}

func (L *List) Display() {
	list := L.head
	for list != nil {
		if list.next != nil {
			fmt.Printf("%v -> ", list.key)
		} else {
			fmt.Printf("%v", list.key)
		}
		list = list.next
	}

	fmt.Printf("\n")		
}

// Private function for printing list
func Display(list *Node) {
	for list != nil {
		if list.next != nil {
			fmt.Printf("%v -> ", list.key)
		} else {
			fmt.Printf("%v", list.key)
		}
		
		list = list.next
	}
	fmt.Println()
}

func (L *List) Reverse() {
	curr := L.head
	var prev *Node
	L.tail = L.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	L.head = prev
	Display(L.head)
}