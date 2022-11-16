package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomPoint(n int) int {
	return rand.Intn(n)

}

func randomPointCalculation(x, y int) {
	x = randomPoint(50)
	y = randomPoint(100)
	fmt.Println(2*x + 10)
	fmt.Println(-3*y - 5)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomPointCalculation(5, 10)
}
