package main

import (
	"fmt"
)

func main() {
	println("hi")
	fmt.Println("===%d,格式化整型")
	fmt.Printf("%d\n", 110)
	fmt.Println("===%s,格式化字符串")
	fmt.Printf("%s\n", "hi")
	fmt.Println("===%f,格式化浮点型")
	fmt.Printf("%0.3f\n", 3.14159)
	fmt.Println("===%v,格式化任意类型")
	fmt.Printf("%v, %v\n", "hi", 110)
}
