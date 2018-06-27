package main

import (
	"bufio"
	"fmt"
	"os"
)

// 输入输出
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a word:")
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err, 2333)
	} else {
		fmt.Println(line)
	}
}
