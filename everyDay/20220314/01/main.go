package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func(x int) {
		fmt.Println(x)
	}(<-ch1)
	ch1 <- 5
	time.Sleep(2 * time.Second)
}
