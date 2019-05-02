package main

import "fmt"

func sum(values []int, resultChan chan int)  {
    sum := 0
    for _, value := range values{
        sum += value
    }
    resultChan <- sum  // 结果并发channel
}

func main()  {
    values := [] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(values)
    resultChan := make(chan int, 2)  // 建一个两个值的管道
    go sum(values[:len(values)/2], resultChan)  // 起一个协程
    go sum(values[len(values)/2:], resultChan)  // 起一个协程
    sum1, sum2 := <- resultChan, <- resultChan // 接收结果
    fmt.Println("Result:",  sum1, sum2, sum1 + sum2)
}
