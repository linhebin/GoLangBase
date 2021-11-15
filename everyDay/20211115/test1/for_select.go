package main

import "fmt"

func produce(id int, ch chan string, quitSub chan int) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("msg:%d:%d", id, i)
		fmt.Println("ok")
	}

	quitSub <- id
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	quitForSelect := make(chan int, 2)

	var quitTag []int

	quitMain := make(chan bool)

	go produce(1, c1, quitForSelect)
	go produce(2, c2, quitForSelect)

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case q := <-quitForSelect:
				fmt.Println("got quit tag for goRouting id:", q)
				quitTag = append(quitTag, q)
				if len(quitTag) == 2 {
					fmt.Println("end to quit Main")
					quitMain <- true
				}
			}
		}
	}()

	<-quitMain
	fmt.Println("exit from main")
}
