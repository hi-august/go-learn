package main

import "fmt"

// 基础的运算和取内存地址
func main() {
	name := "august"
	age := 23
	fmt.Println("+", 5+2)
	fmt.Println("-", 5-2)
	fmt.Println("*", 5*2)
	fmt.Println("/", 5/2)
	fmt.Println("%", 5%2)
	fmt.Println("bool", 5 > 2, 5 <= 2)
	fmt.Println("bool", 5 < 2)
	// fmt.Println("bool", ++age, --age) // todo累加
	fmt.Println(&name) // 取内存地址
}
