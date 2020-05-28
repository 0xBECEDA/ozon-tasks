package main

import (
	"fmt"
	"os"
	"io"
	"strconv"
)

func getNumbersFromFile (slice []byte) []int {

	var	curNum int = 0
	var numsSlice []int

	for i :=0; i< len(slice);  i++ {

		if slice[i] == 10 || slice[i] == 32 {
				numsSlice = append(numsSlice, curNum)
				curNum = 0
		}else {
			curN, errr := (strconv.Atoi(string(slice[i])))
			if errr != nil {
				fmt.Println(" errr : ",  errr)
				os.Exit(1)
			}
			curNum =  int(curN) + (curNum * 10)
		}
	}

	//fmt.Println("return getNumbersFromFile ")
	return numsSlice
}


func findSum (array[]int, target int) {

	var arrLen = len(array)

	//fmt.Println("arrLen: ", len(array))

	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}

	for i:=0; i < arrLen; i++ {
		for j:= i+1; j < arrLen; j++ {
			if target == array[i] + array[j]{
				fmt.Fprintln(file, 1, "\n")
				//fmt.Println("sum: ", array[i], array[j])
				os.Exit(0)
			}
		}
	}
	fmt.Fprintln(file, 0, "\n")
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

		if er == io.EOF {
			break
		}else {

			for i:=0; i < n; i++ {
				buf = append(buf, data[i])
			}
		}
	}

	var allNumbersSlice []int = getNumbersFromFile(buf)
	var target = allNumbersSlice[0]
	//var target = 5678
	var numbers []int

	//fmt.Println("len allNumbers:", len(allNumbersSlice))

	//fmt.Println("target:", target)
	for i:=1; i < len(allNumbersSlice); i++ {
		if allNumbersSlice[i] <= target {
			numbers = append(numbers, allNumbersSlice[i])
		}
	}

	//fmt.Println("len numbers:", len(numbers))

	findSum(numbers, target)
}
