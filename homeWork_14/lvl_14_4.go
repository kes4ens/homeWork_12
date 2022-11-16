package main

var i = 10000

func first(x int) int {
	x = x + i
	return x
}

func second(y int) int {
	y = y + i
	return y
}

func third(z int) int {
	z = z + i
	return z
}

func main() {
	first(i)
	second(first(i))
	third(second(first(i)))
}
