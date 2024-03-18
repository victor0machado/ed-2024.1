package main

import "fmt"

func main() {
	var numCasos int
	var pa, pb int
	var g1, g2 float64
	var resposta = ""

	fmt.Scanln(&numCasos)

	for i := 1; i <= numCasos; i++ {
		fmt.Scanln(&pa, &pb, &g1, &g2)
		if resposta != "" { resposta += "\n" }

		anos := tempoUltrapassagem(pa, pb, g1, g2)
		if anos > 100 {
			resposta += "Mais de 1 seculo."
		} else {
			resposta += fmt.Sprintf("%d anos.", anos)
		}
	}

	fmt.Println(resposta)
}

func tempoUltrapassagem(pa, pb int, g1, g2 float64) int {
	var anos = 0

	for  {
		if pa > pb { break }
		if anos > 100 { return 101 }
		pa = int(float64(pa) * (1 + g1 / 100))
		pb = int(float64(pb) * (1 + g2 / 100))
		anos++
	}

	return anos
}
