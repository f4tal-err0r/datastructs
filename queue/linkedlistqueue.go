package main

import "fmt"

type Node struct {
	data	string
	next	*Node	
}

func (n *Node) add(s string) {

	addnode := &Node{s, nil}
	for {
		if n.next == nil {
			n.next = addnode
			break
		} else {
			n = n.next
		}
	}
}

func (n *Node) remove() *Node {
	n = n.next
	return n 
}

func (n *Node) printAll() {
	iter := n
	for iter != nil {
	   fmt.Println(iter.data)
	   iter = iter.next
	}
}

func main() {
	queue := &Node{"test", nil}
	queue.add("is")
	queue.add("this")

	queue.remove()

	queue.printAll()

	
}