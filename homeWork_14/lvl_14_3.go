package main

import "fmt"

func firstPoint(x int) (b, d int) {
	b = x + 5
	d = x * 10
	fmt.Println(b, d)
	return b, d
}

func main() {
	firstPoint(10)
}
