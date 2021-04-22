// Example of implementation binary tree
package main 

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Left *Node
	Value int
	Right *Node
}

func traverse(node *Node) {
	if node == nil {
		return
	}
	traverse(node.Left)
	fmt.Print(node.Value, " ")
	traverse(node.Right)
}


func create(n int) *Node {
	var node *Node
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2*n; i++ {
		tmp := rand.Intn(n * 20)
		node = insert(node, tmp)
	}

	return node
}

func insert(node *Node, n int) *Node {
	if node == nil {
		return &Node{nil, n, nil}
	}

	// if given value already exist in tree return that node
	if node.Value == n {
		return node
	}

	if n < node.Value {
		node.Left = insert(node.Left, n)
		return node
	}

	
	node.Right = insert(node.Right, n)
	return node	
}

func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}