package main

import (
    "fmt"
    "unsafe"
)

func main()  {
    var boolVar bool
    fmt.Println(boolVar)
    var number1 uint8 // 无符号 8 位整型 (0 到 255)
    number1 = 255
    fmt.Println(number1)
    var number2 uint16 // 无符号 16 位整型 (0 到 65535)
    number2 = 65535
    fmt.Println(number2)
    var number3 uint32 // 无符号 32 位整型 (0 到 4294967295) 42.9亿
    number3 = 4294967295
    fmt.Println(number3)
    var number4 uint64 // 无符号 64 位整型 (0 到 18446744073709551615) 20位数字
    number4 = 18446744073709551615
    fmt.Println(number4)

    var number5 int8 // 有符号 8 位整型 (-128 到 127)
    number5 = 127
    fmt.Println(number5)
    var number6 int16 // 有符号 16 位整型 (-32768 到 32767)
    number6 = 32767
    fmt.Println(number6)
    var number7 int32 // 有符号 32 位整型 (-2147483648 到 2147483647) 21.4亿
    number7 = 2147483647
    fmt.Println(number7)
    var number8 int64 // 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) 19位
    number8 = 9223372036854775807
    fmt.Println(number8)

    var number9 float32 // IEEE-754 32位浮点型数  39位[前340为最大]
    number9 = 340256789012345678901234567890123456789.999999
    fmt.Println(number9)
    var number10 float64 // IEEE-754 64位浮点型数  155位
    number10 = 12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345.243
    fmt.Println(number10)

    var number11 complex64 // 32 位实数和虚数 39位[前340为最大]
    number11 = 340256789012345678901234567890123456789.243
    fmt.Println(number11)
    var number12 complex128 // 32 位实数和虚数 155位
    number12 = 12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345.243
    fmt.Println(number12)

    var number13 byte  // 类似 uint8  无符号 8 位整型 (0 到 255)
    number13 = 255
    fmt.Println(number13)
    var number14 rune // 类似 int32 有符号 32 位整型 (-2147483648 到 2147483647) 21.4亿
    number14 = 2147483647
    fmt.Println(number14)
    var number15 uint // 无符号 64 位整型 (0 到 18446744073709551615) 20位数字
    number15 = 18446744073709551615
    fmt.Println(number15)
    var number16 int // 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) 19位
    number16 = 9223372036854775807
    fmt.Println(number16)
    var number17 uintptr // 无符号整型，用于存放一个指针, 16进制数 (0x0 到 0xffffffffffffffff) 16个f
    number17 = 0xffffffffffffffff
    fmt.Println(number17)

    const length int = 10
    fmt.Println(length)

    const (
        aa = "abc"
        bb = len(aa)
        cc = unsafe.Sizeof(aa)
    )

    fmt.Println(aa, bb, cc)

    const (
        a = iota   //0
        b          //1
        c          //2
        d = "ha"   //独立值，iota += 1
        e          //"ha"   iota += 1
        f = 100    //iota +=1
        g          //100  iota +=1
        h = iota   //7,恢复计数
        i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)

    const (
        m=1<<iota  // m=1<<0 （<< 表示左移的意思） 1 , 1
        j=3<<iota  // j=3<<1  11, 110
        k          // k=3<<2  11, 1100
        l          // k=3<<3  11, 11000
    )
    fmt.Println("m=",m)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}
