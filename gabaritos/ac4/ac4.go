package main

import "fmt"

var totalMovimentos = 0

func main() {
	hanoi(7, "x", "y", "z")

	fmt.Printf("Resolvido em %d movimentos.\n", totalMovimentos)

	fmt.Println(buscaPar([]int{1, 2, 4, 5, 7}, 3))
	fmt.Println(buscaPar([]int{1, 2, 4, 5, 7}, 12))
	fmt.Println(buscaPar([]int{1, 2, 4, 5, 7}, 25))
}

func hanoi(tamanho int, origem string, destino string, trabalho string) {
	if tamanho > 0 {
		hanoi(tamanho - 1, origem, trabalho, destino)
		fmt.Printf("Move peça %d: %s -> %s\n", tamanho, origem, destino)
		totalMovimentos++
		hanoi(tamanho - 1, trabalho, destino, origem)
	}
}

func buscaPar(arr []int, alvo int) (int, int) {
	esq, dir := 0, len(arr) - 1

	for esq < dir {
		soma := arr[esq] + arr[dir]
		if soma == alvo {
			return arr[esq], arr[dir]
		}

		if soma > alvo {
			dir--
		} else {
			esq++
		}
	}

	return -1, -1 // não encontrou nenhum par
}
