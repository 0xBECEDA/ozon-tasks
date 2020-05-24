package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"strconv"
	"strings"
)
type Tree struct {
	root *Node
}

type Node struct {
	key   int
	left  *Node
	right *Node
}

// Tree
func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

func (t *Tree) remove(data int) {
	if t.root == nil {
		fmt.Printf("empty tree\n")

	} else {
		t.root.remove(data)
	}
}

func (t *Tree) getRootKey() {
	if t.root == nil {
		fmt.Printf("empty tree\n")

	} else {
		t.root.getRootKey()
	}
}


// Node
func (n *Node) insert(data int) {
	if data < n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	}
	if  data > n.key {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
	if data == n.key {
	n.remove(n.key)
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d\n", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func (t *Tree) fillBst(dataSlice []int) {

	var length int = len(dataSlice)

	for i := 0; i < length; i++ {

		t.insert(dataSlice[i])
	}
}


func getSmallestElt(n *Node) *Node {
	// fmt.Printf("getLeafLeftPointer: n.key %d\n n.left %d\n n.right %d\n ",
	// 	n.key, n.left, n.right)

	if n.left == nil {
		fmt.Printf("n %d \n", n)
		return n
	}

	return getSmallestElt(n.left)
}


func (n *Node) remove (data int) *Node {

	if data < n.key {
		if n.left == nil {
			fmt.Println("deleteNode:  дереве нет такого элемента: %d\n", data)
			return n
		}else{
			n.left = n.left.remove(data)
			return n
		}
	}

	if data > n.key {
		if n.right == nil {
			fmt.Println("deleteNode:  дереве нет такого элемента: %d\n", data)
			return n
		}else{
			n.right = n.right.remove(data)
			return n
		}
	}

	if n.right == nil && n.left == nil {
		n = nil
		return nil
	}

	if n.right == nil{
		return n.left
	}

	if n.left == nil{
		return n.right
	}

	var smallestElt *Node = getSmallestElt(n.right)

	n.key = smallestElt.key
	n.right = n.right.remove(smallestElt.key)
	return n
}

	var t Tree

func (n *Node) getRootKey() int {
	return n.key
}

func getUnicElt () int {

}

func getNumbersAndMakeTree (slice []byte) []int {

	var stringNumber = ""
	var convertedNumbersIntoInt []int
	var j int = 0

	for i :=0; i++; i< len(slice) {

		if (slice[i]) == io.EOF {
			break
		}

		if (string (slice[i])) != "\n" {
			stringNumber = stringNumber + (string (slice[i]))
		}

		if (string (slice[i])) == "\n" {
			number := strconv.Atoi(stringNumber)
			convertedNumbersIntoInt[j] = int(number)
			stringNumber = ""
		}
		return convertedNumbersIntoInt
	}
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := make([]byte, 100000)
	n, er := ffile.Read(data)

	if er == io.EOF {   // если конец файла
		fmt.Println(er)
		os.Exit(1)
	}


	// var dataSlice []int = []int{8, 7, 15, 24, 3, 6, 1, 0, 22, 28, 23, 20, 10, 9, 12, 8}

	// t.fillBst(dataSlice)
	// fmt.Println("before delete\n")
	// printPreOrder(t.root)

	// t.remove(25)
	// fmt.Println("after delete\n")
	// printPreOrder(t.root)


}
