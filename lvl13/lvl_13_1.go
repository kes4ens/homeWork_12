package main

import "fmt"

func sumEven(a, b int) {
	if a > b {
		var c = a
		a = b
		b = c
	}
	for i := 0; a <= b; a++ {
		if a%2 == 0 {
			i = i + a
			fmt.Println(i)
		}
	}
}

func main() {
	sumEven(1, 10)
	sumEven(10, 1)
}
