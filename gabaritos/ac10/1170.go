package main

import "fmt"

func dias(peso float64) int {
	d := 0
	for peso > 1 {
		d++
		peso /= 2
	}

	return d
}

func main() {
	var n int
	var peso float64

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&peso)
		fmt.Println(dias(peso), "dias")
	}
}
