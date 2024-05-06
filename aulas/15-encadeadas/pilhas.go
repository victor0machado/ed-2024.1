package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Pilha struct {
	tamanho int
	topo    *No
}

func (p *Pilha) percorre() {
	if p.topo == nil {
		fmt.Println("Pilha vazia")
	} else {
		no := p.topo
		for no.prox != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println(no.valor)
	}
}

func (p *Pilha) insere(novoValor int) {
	novoNo := &No{valor: novoValor}

	if p.topo == nil {
		p.topo = novoNo
	} else {
		novoNo.prox = p.topo
		p.topo = novoNo
	}

	p.tamanho++
}

func (p *Pilha) remove() error {
	if p.topo == nil {
		return fmt.Errorf("Pilha vazia")
	}

	if p.tamanho == 1 {
		p.topo = nil
	} else {
		aux := p.topo
		p.topo = p.topo.prox
		aux.prox = nil
	}

	p.tamanho--
	return nil
}

func main() {
	pilha := Pilha{}

	pilha.insere(2)
	pilha.insere(4)
	pilha.insere(8)

	pilha.percorre()
	fmt.Println(pilha.tamanho)

	pilha.remove()

	pilha.percorre()
	fmt.Println(pilha.tamanho)

	pilha.remove()
	pilha.remove()

	pilha.percorre()
	fmt.Println(pilha.tamanho)

	err := pilha.remove()
	fmt.Println(err)
}
