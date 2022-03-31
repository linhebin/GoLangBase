package main

import "fmt"

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	fmt.Println(&w.i)
	fmt.Println("ShowA", w.i)
	w.i = w.i + 10
	return w.i
}

func (w Work) ShowB() int {
	fmt.Println(&w.i)
	fmt.Println("ShowB", w.i)
	w.i = w.i + 20
	return w.i
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
	fmt.Println(c.ShowA())
	fmt.Println(c.ShowB())
}
