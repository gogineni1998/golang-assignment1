package main

import (
	"fmt"

	stdinput "github.com/gogineni1998/golang-assignment1/Question1/input"
	"github.com/gogineni1998/golang-assignment1/Question1/operations"
)

func main() {
	fmt.Println("Follow the options to use the calculator : ")
	fmt.Println(" 1 : Add\n", "2 : Subtract\n", "3 : Divide\n", "4 : Multiply")

	options := stdinput.GetInput()
	fmt.Println("Enter first value")
	a := stdinput.GetInput()
	fmt.Println("Enter second value")
	b := stdinput.GetInput()

	switch options {
	case 1:
		result := operations.Add(a, b)
		fmt.Println(result)
	case 2:
		result := operations.Subtract(a, b)
		fmt.Println(result)
	case 3:
		result, err := operations.Divide(a, b)
		if err != "" {
			panic(err)
		}
		fmt.Println(result)
	case 4:
		result := operations.Multiply(a, b)
		fmt.Println(result)
	}

}
