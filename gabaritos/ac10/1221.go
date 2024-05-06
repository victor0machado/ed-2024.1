package main

import "fmt"

func ePrimo(n int) {
	div := 2
	for div * div <= n {
		if n % div == 0 {
			fmt.Println("Not Prime")
			return
		}

		div++
	}

	fmt.Println("Prime")
}

func main() {
	var n int
	var numero int

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&numero)
		ePrimo(numero)
	}
}
