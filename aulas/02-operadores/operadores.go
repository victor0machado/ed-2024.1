package main

import "fmt"

func main() {
	var x = 2
	y := 4

	// operadores aritméticos
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	x += 3
	// unários
	x++ // -> x = x + 1
	y--

	// operadores de comparação
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)

	var a, b = true, false
	// operadores lógicos
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!b)

	// Bitwise
	// * e & -> ponteiros
}
