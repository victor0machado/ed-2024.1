package main

import "fmt"

/*
Crie uma função calculaMedia que receba um número variável de parâmetros de
tipo ponto flutuante e retorne a média aritmética desses valores.

Dica: para fazer esse exercício, precisamos usar o conceito de "função variada"
https://www.geeksforgeeks.org/golang-program-that-uses-func-with-variable-argument-list/
*/
func calculaMedia(nums ...float64) float64 {
	soma := 0.0
	// nums funciona como um slice de dados
	for _, num := range nums {
		soma += num
	}

	return soma / float64(len(nums))
}

/*
Construa uma função verificaParidade que receba um inteiro e retorne se o
número é par ou ímpar.
*/
func verificaParidade(num int) bool {
	return num % 2 == 0
}

/*
Desenvolva uma função minhaPotencia que receba dois inteiros, base e expoente,
e retorne o resultado de base elevado ao expoente, sem usar funções prontas da
linguagem.
*/
func minhaPotencia(base, exp int) int {
	pot := 1
	for i := 0; i < exp; i++ {
		pot *= base
	}

	return pot
}

/*
Implemente uma função converteCelsiusParaFahrenheit que receba uma
temperatura em Celsius e retorne a temperatura convertida em Fahrenheit.
*/
func converteCelsiusParaFahrenheit(tempCelsius float64) float64 {
	return tempCelsius * 9.0 / 5.0 + 32.0
}

func main() {
	fmt.Println(calculaMedia(4.8, 7.5, -2.4, 8.8))
	fmt.Println(verificaParidade(18))
	fmt.Println(verificaParidade(3))
	fmt.Println(minhaPotencia(4, 4))
	fmt.Println(minhaPotencia(3, 5))
	fmt.Println(converteCelsiusParaFahrenheit(80.0))
	fmt.Println(converteCelsiusParaFahrenheit(100.0))
}
