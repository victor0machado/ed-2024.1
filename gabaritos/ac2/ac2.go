package main

import (
	"ac2/geometria"
	"fmt"
	"math"
)

func ex1() {
	var vetor = [10]int{1, 3, 2, 0, -1, 2, 10, 5, 8, 9}

	for _, valor := range vetor {
		fmt.Println(valor)
	}
}

// Referência para o exercício 2
// https://medium.com/@onezino.moreira/entendendo-string-em-go-c4628055f60d
func ex2(texto string) string {
	novaString := ""

	for _, caractere := range texto {
		novaString = string(caractere) + novaString
	}

	return novaString
}

// Referência para o exercício 3
// https://golangbyexample.com/power-golang/
type Ponto struct {
	X, Y float64
}

func (p Ponto) DistanciaOrigem() float64 {
	return math.Pow(p.X * p.X + p.Y * p.Y, 0.5)
}

func main() {
	ex1()

	fmt.Println(ex2("Olá, mundo"))

	p := Ponto{3.0, 4.0}
	fmt.Println(p.DistanciaOrigem())

	ret := geometria.Retangulo{Largura: 4.5, Altura: 2.0}
	fmt.Println(ret.CalcularArea(), ret.CalcularPerimetro())
}
