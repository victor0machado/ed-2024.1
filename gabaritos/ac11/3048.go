package main

import "fmt"

func main() {
	var n, n_lido int
	numeros := make([]int, 0)

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&n_lido)
		numeros = append(numeros, n_lido)
	}

	cont := 1
	for i := 0; i < n - 1; i++ {
		if numeros[i] != numeros[i + 1] {
			cont++
		}
	}

	fmt.Println(cont)
}
