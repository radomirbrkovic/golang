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

func (L *List) Remove(key interface{}) {
	node := L.head
	for node != nil {
		if node.key == key {
			if node.next != nil {
				node.next.prev = node.prev
			}
			if node.prev != nil {
				node.prev.next = node.next
			}
			break
		}
		node = node.next
	}

}

func (L *List) Display() {
	list := L.head
	for list != nil {
		prev := 0 
		if list.prev != nil {
			prev, _ = list.prev.key.(int)	
		}

		if list.next != nil {
			fmt.Printf("%v(%v) -> ", list.key, prev)
		} else {
			fmt.Printf("%v(%v)", list.key,  prev)
		}
		list = list.next
	}

	fmt.Printf("\n")		
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
}