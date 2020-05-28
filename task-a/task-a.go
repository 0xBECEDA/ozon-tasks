package main

import (
	"fmt"
	"os"
	"io"
	"strconv"
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
		n.key = data
		t.remove(n.key)
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d \n", n.key)
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
	if n.left == nil {
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


func (n *Node) getRootKey() int {
	return n.key
}

func getNumbersFromFile (slice []byte) []int {

	var	curNum int = 0
	var numsSlice []int

	for i :=0; i< len(slice);  i++ {

		if slice[i] == 10 || slice[i] == 32 {
			if curNum != 0 {
				numsSlice = append(numsSlice, curNum)
				curNum = 0
			}

		}else if slice[i] ==  0 {
			break

		}else {
			curN, errr := (strconv.Atoi(string(slice[i])))
			if errr != nil {
				fmt.Println(" errr : ",  errr)
				os.Exit(1)
			}
			curNum =  int(curN) + (curNum * 10)
		}
	}
	return numsSlice
}


var t Tree

func main() {
	file, err := os.Open("input-201.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data:= make([]byte, 100)
	var buf []byte
	for {
		n, er := file.Read(data)

		if er == io.EOF {   // если конец файла
			break
		}else {

			for i:=0; i < n; i++ {
				buf = append(buf, data[i])
			}
		}
	}

	var dataSlice []int = getNumbersFromFile(buf)

	t.fillBst(dataSlice)
	unicN := t.root.getRootKey()

	fmt.Println(unicN)
}
