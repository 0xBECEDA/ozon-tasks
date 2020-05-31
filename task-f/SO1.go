package main

import (
	"fmt"
	"os"
	"io"
	"strconv"
	"sort"
)

func getNumbersFromFile (slice []byte) []int {

	var curNum int = 0
	var numsSlice []int

	// do before the end of slice
	for i :=0; i< len(slice);  i++ {

		// if we got "\n" or "\t"
		if slice[i] == 10 || slice[i] == 32 {
			// and numer isn't 0
			if curNum != 0 {
				// write it to new slice
				numsSlice = append(numsSlice, curNum)
				//fmt.Println(" curNum : ",  curNum)
				curNum = 0
			}
			// oterwise create number
		}else {
			curN, errr := (strconv.Atoi(string(slice[i])))
			if errr != nil {
				fmt.Println(" errr : ",  errr)
				os.Exit(1)
			}
			// make a int instead of byte
			curNum =  int(curN) + (curNum * 10)
		}
	}
	return numsSlice
}

func findSum (array[]int, target int) {

	var arrLen = len(array)

	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}

	for i:=0; i < arrLen; i++ {
		term := target - array[i]
		indx := sort.Search(arrLen, func(indx int) bool { return array[indx] >= term })

		if indx < arrLen && array[indx] == term && array[indx] + array[i] == target {
			fmt.Fprintln(file, 1, "\n")
			//fmt.Println("sum :", term, array[i])

			os.Exit(0)
		}
	}

	fmt.Fprintln(file, 0, "\n")
}

func main() {

    //open input-file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data:= make([]byte, 100)
	var buf []byte

	// read all file
	for {
		n, er := file.Read(data)

		if er == io.EOF {
			break
		}else {
			for i:=0; i < n; i++ {
				buf = append(buf, data[i])
			}
		}
	}
    // create numbers from read strings
 	var allNumbersSlice []int = getNumbersFromFile(buf)
	// get target-number
	var target = allNumbersSlice[0]

	var numbers []int
	//fmt.Println("target:", target)

	// throw out all numbers, which are bigger, than target
	for i:=1; i < len(allNumbersSlice); i++ {
		if allNumbersSlice[i] <= target {
			numbers = append(numbers, allNumbersSlice[i])
		}
	}

	// sort numbers from small to big
	sort.Ints(numbers)

	// try to find sum
	findSum(numbers, target)


}
