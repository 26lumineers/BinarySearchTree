package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert
func (n *Node) Insert(k int) {

	if n.Key < k {
		// move to right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move to left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}

}

//Serch
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	}else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}
func main() {
	fmt.Println("BinarySearchTree...")
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(400)
	tree.Insert(130)
	tree.Insert(60)
	tree.Insert(70)
	tree.Insert(20)
	fmt.Println(tree)
	fmt.Println(tree.Search(60))

}
