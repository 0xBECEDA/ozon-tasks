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
		//fmt.Printf("remove: %d \n", n.key)
		n.key = data
		// n.remove(n.key)
		// printPreOrder(t.root)
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

		// 	if dataSlice[i] == 0 {
		// 		break
		// 	}
		//fmt.Printf("insert: %d \n", dataSlice[i])
		t.insert(dataSlice[i])
		//printPreOrder(t.root)


	}
}



func getSmallestElt(n *Node) *Node {
	// fmt.Printf("getLeafLeftPointer: n.key %d\n n.left %d\n n.right %d\n ",
	// 	n.key, n.left, n.right)

	if n.left == nil {
		//fmt.Printf("n %d \n", n)
		return n
	}

	return getSmallestElt(n.left)
}


func (n *Node) remove (data int) *Node {
	// fmt.Printf("remove: %d \n", n.key)
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

	// fmt.Printf("remove: found! %d \n", n.key)
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

// func getUnicElt () int {

// }

func getNumbersFromFile (slice []byte) []int {

	var j int = 0
	var	curNum int = 0
	var numsSlice []int
	numsSlice = make([]int, len(slice) / 4, len(slice) / 4)

	for i :=0; i< len(slice);  i++ {

		if slice[i] == 10 || slice[i] == 32 {
			// fmt.Println("curNum: ",  curNum)
			// fmt.Println("j: ",  j)
			if curNum != 0 {
				numsSlice[j] = curNum
				j = j + 1
				curNum = 0
			}
			// fmt.Println("numsSlice[j]: ",  numsSlice[j-1])

		}else if slice[i] ==  0 {
			break

		}else {
			// fmt.Println("slice[i]): ",  slice[i])
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

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data:= make([]byte, 100)
	var buf []byte
	for {
		n, er := file.Read(data)

		if er == io.EOF {   // если конец файла
			fmt.Println(er)
			break
		}else {

			for i:=0; i < n; i++ {
				buf = append(buf, data[i])
			}
		}
	}

	var dataSlice []int = getNumbersFromFile(buf)

	// for i := 0; i < len(dataSlice); i++ {

	// 	if dataSlice[i] == 0 {
	// 		break
	// 	}

	// 	fmt.Println("elt: ", dataSlice[i], "i ", i, "\n", )
	// }


	t.fillBst(dataSlice)
	fmt.Println("after delete\n")
	printPreOrder(t.root)
	unicN := t.root.getRootKey()

	fmt.Println("unicN ", unicN)

}
