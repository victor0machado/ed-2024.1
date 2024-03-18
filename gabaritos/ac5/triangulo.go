package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	a, b, c = ordenarDescrescente(a, b, c)

	if a >= b + c {
		fmt.Println("NAO FORMA TRIANGULO")
		return
	}

	if a * a == b * b + c * c {
		fmt.Println("TRIANGULO RETANGULO")
	} else if a * a > b * b + c * c {
		fmt.Println("TRIANGULO OBTUSANGULO")
	} else {
		fmt.Println("TRIANGULO ACUTANGULO")
	}

	if a == b && b == c {
		fmt.Println("TRIANGULO EQUILATERO")
	} else if a == b || a == c || b == c {
		fmt.Println("TRIANGULO ISOSCELES")
	}
}

func ordenarDescrescente(a, b, c float64) (float64, float64, float64) {
	if a < b { a, b = b, a }
	if b < c { b, c = c, b }
	if a < b { a, b = b, a }
	return a, b, c
}
