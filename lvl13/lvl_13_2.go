package main

import "fmt"

func swap(c, d int) (int, int) {
	return d, c
}
func main() {
	a, b := swap(10, 15)
	fmt.Println(a, b)
}
