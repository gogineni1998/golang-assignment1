package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n *int
var array []int
var key *int

func main() {
	result := getIndex()
	fmt.Println("Index of the required value is : ", result)
}

func getIndex() int {
	for i := 0; i < len(array); i++ {
		if array[i] == *key {
			return i
		}
	}
	return -1
}

func init() {
	fmt.Println("Enter the number of elements in the array : ")
	n = input()
	fmt.Println("Enter the values in the array : ")
	for i := 0; i < *n; i++ {
		fmt.Println("Enter the value ", i+1)
		array = append(array, *input())
	}
	fmt.Println("Enter the value to be found in the array : ")
	key = input()
}

func input() *int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hit Enter after entering the value")
	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	value, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		panic(err)
	}
	return &value
}
