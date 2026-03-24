package main

import "fmt"

func add(x, y, z int) int {
	return x + y + z
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(20, 13, 3))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}