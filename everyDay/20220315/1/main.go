package main

import "fmt"

func foo(x []int) {
	x = append(x, 1, 2, 2)
	x[0] = 200
	fmt.Println(len(x), cap(x), x)
}

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2)
	foo(x)
	fmt.Println(len(x), cap(x), x)
	y := []int{2: 1}
	fmt.Println(len(y), cap(y), y)
	var z = []int{4: 44, 55, 66, 1: 77, 88}
	fmt.Println(len(z), cap(z), z)
	var l = []int{4: 44, 55, 66, 3: 77}
	fmt.Println(len(l), cap(l), l)
}
