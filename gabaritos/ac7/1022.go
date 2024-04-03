package main

import "fmt"

var operacoes = map[string]func(int, int, int, int) (int, int) {
	"+": soma,
	"-": subtracao,
	"*": multiplicacao,
	"/": divisao,
}

func soma(n1, d1, n2, d2 int) (int, int) {
	return n1 * d2 + n2 * d1, d1 * d2
}

func subtracao(n1, d1, n2, d2 int) (int, int) {
	return n1 * d2 - n2 * d1, d1 * d2
}

func multiplicacao(n1, d1, n2, d2 int) (int, int) {
	return n1 * n2, d1 * d2
}

func divisao(n1, d1, n2, d2 int) (int, int) {
	return n1 * d2, n2 * d1
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func calculaMdc(a, b int) int {
	if a < 0 { a = -1 * a }
	if b < 0 { b = -1 * b }

	var mdc = 1
	for i := 1; i <= min(a, b); i++ {
		if a % i == 0 && b % i == 0 { mdc = i }
	}
	return mdc
}

func simplifica(dividendo, divisor int) (int, int) {
	mdc := calculaMdc(dividendo, divisor)
	return dividendo / mdc, divisor / mdc
}

func main() {
	var n, n1, n2, d1, d2 int
	var oper, tmp string

	resposta := ""

	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Scanln(&n1, &tmp, &d1, &oper, &n2, &tmp, &d2)
		dividendo, divisor := operacoes[oper](n1, d1, n2, d2)
		resposta += fmt.Sprintf("%d/%d = ", dividendo, divisor)
		dividendo, divisor = simplifica(dividendo, divisor)
		resposta += fmt.Sprintf("%d/%d\n", dividendo, divisor)
	}

	fmt.Print(resposta)
}
