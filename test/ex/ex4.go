package main

import (
	"fmt"
	"os"
	"strings"
)

// 命令行参数
func main() {
	name := "august"
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}
	fmt.Println(name)
}
