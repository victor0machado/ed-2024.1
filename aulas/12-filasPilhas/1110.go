package main

import "fmt"

const M = 50

type Fila struct {
	fila    [M]int
	in, fim int
	tam     int
}

func (f *Fila) inicializar(n int) {
	f.in, f.fim = -1, -1
	for i := 1; i <= n; i++ {
		f.inserir(i)
	}
}

func (f *Fila) inserir(valor int) error {
	aux := (f.fim + 1) % M
	if aux == f.in {
		return fmt.Errorf("Overflow")
	}

	f.fila[aux] = valor
	f.fim = aux
	if f.in == -1 {
		f.in = 0
	}
	f.tam++
	return nil
}

func (f *Fila) remover() (int, error) {
	if f.in == -1 {
		return -1, fmt.Errorf("Underflow")
	}

	valorRemovido := f.fila[f.in]
	f.fila[f.in] = 0

	if f.in == f.fim {
		f.in, f.fim = -1, -1
	} else {
		f.in = (f.in + 1) % M
	}

	f.tam--
	return valorRemovido, nil
}

func exibirDescartados(array []int) {
	fmt.Printf("Discarded cards: ")
	for i, num := range array {
		if i < len(array) - 1 {
			fmt.Printf("%d, ", num)
		} else {
			fmt.Println(num)
		}
	}
}

func main() {
	var n int

	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		}

		fila := Fila{}
		fila.inicializar(n)

		descartados := make([]int, 0)
		for fila.tam > 1 {
			valorDescartado, _ := fila.remover()
			descartados = append(descartados, valorDescartado)

			valorDescartado, _ = fila.remover()
			fila.inserir(valorDescartado)
		}

		exibirDescartados(descartados)
		valorDescartado, _ := fila.remover()
		fmt.Printf("Remaining card: %d\n", valorDescartado)
	}
}
