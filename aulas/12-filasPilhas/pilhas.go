package main

import "fmt"

const M = 5

func exibir(pilha *[M]int, topo *int) {
	for i := 0; i <= *topo; i++ {
		fmt.Printf("%d ", pilha[i])
	}
	fmt.Println("")
}

func push(pilha *[M]int, topo *int, valor int) error {
	if *topo == M-1 {
		return fmt.Errorf("Overflow")
	}
	*topo++
	pilha[*topo] = valor
	return nil
}

func pop(pilha *[M]int, topo *int) (int, error) {
	if *topo == -1 {
		return -1, fmt.Errorf("Underflow")
	}
	valorRemovido := pilha[*topo]
	pilha[*topo] = 0
	*topo--
	return valorRemovido, nil
}

func main() {
	var pilha [M]int
	var topo = -1

	exibir(&pilha, &topo)  // []
	push(&pilha, &topo, 2) // [2]
	push(&pilha, &topo, 3) // [2, 3]
	push(&pilha, &topo, 4) // [2, 3, 4]
	exibir(&pilha, &topo)  // [2, 3, 4]
	pop(&pilha, &topo)     // [2, 3]
	pop(&pilha, &topo)     // [2]
	exibir(&pilha, &topo)  // [2]

	valor, err := pop(&pilha, &topo)
	if err != nil { fmt.Println("Erro:", err) }
	fmt.Println(valor)

	valor, err = pop(&pilha, &topo)
	if err != nil { fmt.Println("Erro:", err) }
	fmt.Println(valor)
}
