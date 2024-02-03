package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int

func main() {
	first := 0
	if n >= 1 {
		fmt.Print(first, " ")
	}
	second := 1
	if n >= 2 {
		fmt.Print(second, " ")
	}
	if n < 3 {
		return
	}
	for i := 2; i < n; i++ {
		fmt.Print(first+second, " ")
		temp := first
		first = second
		second = temp + second
	}
}

func init() {
	fmt.Println("Enter the limit : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	value, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		panic(err)
	}
	n = value
}
