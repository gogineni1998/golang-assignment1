package main

import "fmt"

func main() {
	a := "###5#4456#"
	prevHash := false
	result := ""
	temp := ""
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '#' {
			prevHash = true
		} else if !prevHash {
			temp = temp + string(a[i])
		} else {
			prevHash = false
		}
	}
	for i := len(temp) - 1; i >= 0; i-- {
		result = result + string(temp[i])
	}
	fmt.Println(result)
}
