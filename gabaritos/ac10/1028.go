package main

import "fmt"

func mdc(a, b int) int {
	var resto int
	for b != 0 {
		resto = a % b
		a = b
		b = resto
	}

	return a
}

func main() {
	var n int
	var x, y int

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&x, &y)
		fmt.Println(mdc(x, y))
	}
}