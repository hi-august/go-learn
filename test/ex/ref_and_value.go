package main

import (
    "fmt"
)

func add(a int) int {
    a = a + 2
    return a
}

func add2(a *int) int {
    *a = *a + 2
    return *a
}

// 传递指针可以使多个函数操作同一个对象
// 传指针只是传递内存地址,节省时间和内存
// go语言中slice,channel,map这三类实现机制和指针类似,可以直接传递,不用取指针传递
func main() {
    x := 5
    x1 := add(x) // 传递x的一个copy
    fmt.Println("x is", x, "\r\nx1 is", x1) // out: 5, 7
    x2 := add2(&x) // 传递了一个引用(指针),这个会修改原引用的值
    fmt.Println("x is", x, "\r\nx2 is", x2) // out: 7, 7
}
