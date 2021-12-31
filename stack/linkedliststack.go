package main

import "fmt"

type Node struct {
	data	string
	next 	*Node
}

func (n *Node) pop() (string) {
	data := n.data
	n = n.next
	return data
}

func (n *Node) push(d string) {
	addnode := &Node{data: d, next: n}
	n = addnode
}

func (n *Node) peek() string {
	return n.data
}

func (n *Node) isEmpty() bool {
	return n.data == nil
}

func main() {
	stack := &Node{"this", nil}
	stack.push("is")
	stack.push("a")
	stack.push("test")

	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}