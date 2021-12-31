package main 

import (
	"fmt"
)

type Node struct {
	data	int
	next	*Node
}

func (n *Node) appendToTail(d int) {
	addnode := &Node{data: d, next: nil}

	for {
		if n.next != nil {
			n = n.next
		} else {
			break
		}
	}

	n.next = addnode
}

func (n *Node) printAll() {
	iter := n
	for iter != nil {
	   fmt.Println(iter.data)
	   iter = iter.next
	}
}

func (n *Node) deleteNode(d int) *Node {
	if n.data == d {
		return n.next
	}

	for n.next != nil {
		if n.next.data == d {
			n.next = n.next.next
			return n
		}
	n = n.next 
	}

	return n
}


func main() {
	newList := Node{8, nil}
	newList.appendToTail(5)
	newList.appendToTail(6)
	newList.appendToTail(7)
	newList.deleteNode(6)
	newList.printAll()
	fmt.Println("Finished")
}