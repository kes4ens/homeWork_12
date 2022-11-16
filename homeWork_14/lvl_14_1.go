package main

import "fmt"

func trueOrfalse(n int) (int, bool) {
	if n%2 == 0 {
		fmt.Print("true")
		return n, true
	} else {
		fmt.Print("false")
		return n, false
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	trueOrfalse(n)
}
