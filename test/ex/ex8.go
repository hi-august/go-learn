package main

import (
	"fmt"
)

func main() {
	// 声明数组
	cities := [...]string{"Shanhai", "Fuzhou", "Congqin"}
	fmt.Printf("%-8T %2d, %q\n", cities, len(cities), cities)
	// 声明切片
	// []里有没有内容,如果有则为数组,没有则为切片
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := s[:5]
	u := s[3 : len(s)-1]
	fmt.Println(s, t, u)
	// 遍历切片
	for i, letter := range s {
		if letter != "G" {
			fmt.Println(i, letter)
		}
	}
	s2 := make([]string, 3)
	s2[0] = "hello"
	s2[1] = " "
	s2 = append(s2, "world!")
	// s2[2] = "world!"
	fmt.Println(s2)
}
