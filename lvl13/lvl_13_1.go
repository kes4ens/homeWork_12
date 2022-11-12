package main

import "fmt"

func sumEven(a, b int) {
	var result int
	if a > b {
		a, b = b, a
	}
	for i := 0; a <= b; a++ {
		if a%2 == 0 {
			i += +a
			result = i
		}
	}
	fmt.Println(result)
}

func main() {

	sumEven(1, 10)
	sumEven(10, 1)
}
