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

	L.head = list
	head := L.head

	for head.next != nil {
		head = head.next
	}

	L.tail = head

}

func (L *List) Remove(key interface{}) {
	node := L.head
	for node.next != nil {
		if node.next.key == key {
			if node.next != nil {
				node.next = node.next.next
			}
			
			break
		}
		node = node.next
	}

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
