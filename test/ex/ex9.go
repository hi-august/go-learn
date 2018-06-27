package main

import (
	"fmt"
)

func main() {
	mass := make(map[string]int)
	mass["name"] = 1
	mass["age"] = 23
	mass["test"] = 0
	fmt.Println(mass)
}
