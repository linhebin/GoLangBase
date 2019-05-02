package main

import (
    "fmt"
    "reflect"
)

type Bird struct {
    Name string
    LifeExpectance int
}

func (b Bird) Fly()  {
    fmt.Println("I am flying...")
}

func main()  {
    sparrow := &Bird{"Sparrow", 3}  // 初始化一结构体
    s := reflect.ValueOf(sparrow).Elem() // 反射， 得到对像
    typeOft := s.Type()
    for i := 0; i < s.NumField(); i ++ {
        f := s.Field(i)
        fmt.Println(fmt.Sprintf("%d: %s %s = %v\n", i, typeOft.Field(i).Name, f.Type(), f.Interface()))
    }
}