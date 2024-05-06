package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Fila struct {
	tamanho int
	inicio    *No
	fim     *No
}

func (f *Fila) percorre() {
	if f.inicio == nil {
		fmt.Println("Fila vazia")
	} else {
		no := f.inicio
		for no.prox != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println(no.valor)
	}
}

func (f *Fila) insere(novoValor int) {
	novoNo := &No{valor: novoValor}

	if f.inicio == nil {
		f.inicio = novoNo
		f.fim = f.inicio
	} else {
		f.fim.prox = novoNo
		f.fim = novoNo
	}

	f.tamanho++
}

func (f *Fila) remove() error {
	if f.inicio == nil {
		return fmt.Errorf("Fila vazia")
	}

	if f.tamanho == 1 {
		f.inicio = nil
		f.fim = nil
	} else {
		aux := f.inicio
		f.inicio = f.inicio.prox
		aux.prox = nil

		if f.inicio.prox == nil {
			f.fim = f.inicio
		}
	}

	f.tamanho--
	return nil
}

func main() {
	fila := Fila{}

	fila.insere(2)
	fila.insere(4)
	fila.insere(8)

	fila.percorre()
	fmt.Println(fila.tamanho)

	fila.remove()

	fila.percorre()
	fmt.Println(fila.tamanho)

	fila.remove()
	fila.remove()

	fila.percorre()
	fmt.Println(fila.tamanho)

	err := fila.remove()
	fmt.Println(err)
}
