package main

import "fmt"

const M = 10

func bubblesort(x *[M]int) {
	limite := M
	trocou := true

	for trocou == true {
		trocou = false
		limite--
		for i := 0; i < limite; i++ {
			if x[i] > x[i + 1] {
				x[i], x[i + 1] = x[i + 1], x[i]
				trocou = true
			}
		}
	}
}

func selectionSort(x *[M]int) {
	for i := 0; i < M - 1; i++ {
		min := i
		for j := i; j < M; j++ {
			if x[j] < x[min] { min = j }
		}

		if min != i {
			x[i], x[min] = x[min], x[i]
		}
	}
}

func insertionSort(x *[M]int) {
	for j := 1; j < M; j++ {
		i := j - 1
		chave := x[j]
		for i >= 0 && x[i] > chave {
			x[i + 1] = x[i]
			i--
		}

		x[i + 1] = chave
	}
}

func quicksort(x *[M]int, inicio, fim int) {
	i, j := inicio, fim
	pivo := x[i]

	for i <= j {
		for x[i] < pivo { i++ }
		for pivo < x[j] { j-- }
		if i <= j {
			x[i], x[j] = x[j], x[i]
			i++
			j--
		}
	}

	if inicio < j { quicksort(x, inicio, j) }
	if i < fim { quicksort(x, i, fim) }
}

func combina(x *[M]int, inicio, meio, fim int, aux *[M]int) {
	i, j := inicio, meio + 1
	for k := i; k <= fim; k++ {
		aux[k] = x[k]
	}
	k := i

	for i <= meio && j <= fim {
		if aux[i] < aux[j] {
			x[k] = aux[i]
			i++
		} else {
			x[k] = aux[j]
			j++
		}
		k++
	}

	for i <= meio {
		x[k] = aux[i]
		k++
		i++
	}

	for j <= fim {
		x[k] = aux[j]
		k++
		j++
	}
}

func mergesort(x *[M]int, inicio, fim int, aux *[M]int) {
	if inicio < fim {
		meio := int((inicio + fim) / 2)
		mergesort(x, inicio, meio, aux)
		mergesort(x, meio + 1, fim, aux)
		combina(x, inicio, meio, fim, aux)
	}
}

func iniciaMergesort(x *[M]int) {
	var aux [M]int
	mergesort(x, 0, M - 1, &aux)
}

func main() {
	lista := [M]int{3, 4, 6, 8, 1, 2, 10, 9, 7, 5}

	// bubblesort(&lista)
	// selectionSort(&lista)
	// insertionSort(&lista)
	// quicksort(&lista, 0, M - 1)
	iniciaMergesort(&lista)
	fmt.Println(lista)
}
