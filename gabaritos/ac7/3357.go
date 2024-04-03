package main

import "fmt"

func leSlice(n int) []string {
	slice := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Scanln()
	return slice
}

func main() {
	var n, ultimo int
	var l, q, resto float64

	fmt.Scanln(&n, &l, &q)
	num_cuias := int(l / q)
	if int(l * 10) % int(q * 10) == 0 {
		resto = q
		ultimo = num_cuias - 1
	} else {
		resto = (l / q - float64(num_cuias)) * q
		ultimo = num_cuias
	}
	nomes := leSlice(n)
	fmt.Printf("%s %.1f\n", nomes[ultimo % n], resto)
}
