package main

import "fmt"

func main() {
	var n, m, p, quantidade int
	var fruta string
	var valor, valorTotal float64
	var frutas map[string]float64

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&m)
		frutas = make(map[string]float64)
		for j := 0; j < m; j++ {
			fmt.Scanln(&fruta, &valor)
			frutas[fruta] = valor
		}

		fmt.Scanln(&p)
		valorTotal = 0
		for j := 0; j < p; j++ {
			fmt.Scanln(&fruta, &quantidade)
			valorTotal += float64(quantidade) * frutas[fruta]
		}

		fmt.Printf("R$ %.2f\n", valorTotal)
	}
}
