package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := fn(5)
	fmt.Println(111, n)
}

func fn(idx int) string {
	url := strconv.Itoa(idx)
	return url
}
