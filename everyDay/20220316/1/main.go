package main

import "fmt"

func main() {
	var s1 []int
	var s2 = []int{}
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println(s1)
		fmt.Println("no nil")
	}

	if s2 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println(s2)
		fmt.Println("no nil")
	}
}
