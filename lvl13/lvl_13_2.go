package main

import "fmt"

func func1(a int) {
	fmt.Println(a)
}
func func2(b int) {
	fmt.Println(b)
}

func full(first func(int), second func(int), a int, b int) {
	second(b)
	first(a)
}
func main() {
	full(func1, func2, 5, 10)
}
