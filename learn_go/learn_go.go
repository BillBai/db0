package main

import (
	"fmt"
)

func main() {
	var a int
	var b string
	a = 1230
	b = "bonjour"
	var c int = 123
	d := "hi"

	fmt.Println("Hello World! ", a, b, c, d)

	beyondHello()
	learnTypes()
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("Sum:", sum, "Product:", prod)
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Learn GO!"
	s2 := `A "raw" string literal
can include line breaks.`

	g := 'Σ'

	fmt.Println("str: ", str, "\n s2:", s2, "\n g:", g)
}
