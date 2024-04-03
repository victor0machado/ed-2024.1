package main

import "fmt"

func main() {
	var tmp, num1, num2 int
	var preco1, preco2 float64

	fmt.Scanln(&tmp, &num1, &preco1)
	fmt.Scanln(&tmp, &num2, &preco2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", float64(num1) * preco1 + float64(num2) * preco2)
}
