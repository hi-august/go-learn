package main

import "fmt"

func main() {
	i := 9
	j := 5
	i, j, k := add(i, j)
	fmt.Printf("%v %v %d", i, j, k)
	// var a, b, c int
	// a, b, c := add2(&i, &j)
	// fmt.Printf("%v %v %d", i, j, k)

}

func add(x, y int) (int, int, int) {
	return y, x, x + y
}

// func add2(x, y int) (int, int, int) {
// return *y, *x, *x+*y
// }
