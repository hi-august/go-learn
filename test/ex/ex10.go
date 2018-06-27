package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}
func main() {
	_
}
