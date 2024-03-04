package main

import "fmt"

func buscaMatriz(m [][]int, n int, x int) bool {
	i := 0
	var j int

	for i < n {
		j = 0
		for j < n {
			if m[i][j] == x { return true }
			j++
		}
		i++
	}

	return false
}

func main() {
	matriz := [][]int{{1, 2}, {3, 4}}
	fmt.Println(buscaMatriz(matriz, 2, 3))
	fmt.Println(buscaMatriz(matriz, 2, 5))
}
