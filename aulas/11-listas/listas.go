package main

import "fmt"

const M = 7

func main() {
	lista := [M]int{1, 3, 5, 6, 8}
	n := 5

	fmt.Println(busca1(lista, n, 5)) // 2
	fmt.Println(busca1(lista, n, 2)) // -1

	fmt.Println(busca2(lista, n, 5)) // 2
	fmt.Println(busca2(lista, n, 2)) // -1

	fmt.Println(buscaOrd(lista, n, 5)) // 2
	fmt.Println(buscaOrd(lista, n, 2)) // -1

	fmt.Println(buscaBin(lista, n, 5)) // 2
	fmt.Println(buscaBin(lista, n, 2)) // -1

	insere(&lista, &n, 5)
	fmt.Println(lista, n)
	insere(&lista, &n, 7)
	fmt.Println(lista, n)
	insere(&lista, &n, 2)
	fmt.Println(lista, n)
	insere(&lista, &n, 10)
	fmt.Println(lista, n)

	fmt.Println("---------------")
	fmt.Println(remove(&lista, &n, 10))
	fmt.Println(lista, n)
	fmt.Println(remove(&lista, &n, 8))
	fmt.Println(lista, n)
	fmt.Println(remove(&lista, &n, 2))
	fmt.Println(lista, n)
	fmt.Println(remove(&lista, &n, 1))
	fmt.Println(remove(&lista, &n, 3))
	fmt.Println(remove(&lista, &n, 5))
	fmt.Println(remove(&lista, &n, 6))
	fmt.Println(remove(&lista, &n, 7))
	fmt.Println(n)
	fmt.Println(remove(&lista, &n, 8))
}

func remove(v *[M]int, n *int, x int) int {
	if *n != 0 {
		ind := busca1(*v, *n, x)
		if ind == -1 {
			fmt.Println("Valor não encontrado!")
		} else {
			for i := ind; i < *n - 1; i++ {
				v[i] = v[i + 1]
			}
			*n--
			v[*n] = 0
			return x
		}
	} else {
		fmt.Println("Underflow")
	}
	return -1
}

func insere(v *[M]int, n *int, x int) {
	if *n < M {
		if busca1(*v, *n, x) == -1 {
			v[*n] = x
			*n++
		} else {
			fmt.Println("Erro! Elemento já existe.")
		}
	} else {
		fmt.Println("Overflow")
	}
}

func buscaBin(v [M]int, n, x int) int {
	inf, sup := 0, n - 1
	for inf <= sup {
		meio := int((inf + sup) / 2)
		if v[meio] == x {
			return meio
		}
		if v[meio] < x {
			inf = meio + 1
		} else {
			sup = meio - 1
		}
	}
	return -1
}

// Para rodar essa função, o array precisa ter pelo menos um espaço vazio.
func buscaOrd(v [M]int, n, x int) int {
	i := 0
	v[n] = x
	for v[i] < x {
		i++
	}
	if i == n + 1 || v[i] != x {
		return -1
	}
	return i
}

// Para rodar essa função, o array precisa ter pelo menos um espaço vazio.
func busca2(v [M]int, n, x int) int {
	i := 0
	v[n] = x
	for v[i] != x {
		i++
	}
	if i == n {
		return -1
	}
	return i
}

func busca1(v [M]int, n, x int) int {
	i := 0
	for i < n {
		if v[i] == x {
			return i
		}
		i++
	}
	return -1
}
