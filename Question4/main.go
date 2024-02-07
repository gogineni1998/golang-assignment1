package main

import "fmt"

func main() {
	array := []int{4, 2, 3, 1, 6}
	fmt.Println("Input array", array)
	sortArray(array)
	fmt.Println("Sorted array", array)
}

func sortArray(array []int) {

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}
