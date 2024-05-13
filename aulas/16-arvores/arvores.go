package main

import "fmt"

type No struct {
	valor int
	esq   *No
	dir   *No
}

func (n *No) preOrdem() {
	if n == nil { return }
	fmt.Println(n.valor)
	n.esq.preOrdem()
	n.dir.preOrdem()
}

func (n *No) posOrdem() {
	if n == nil { return }
	n.esq.posOrdem()
	n.dir.posOrdem()
	fmt.Println(n.valor)
}

func (n *No) simetrico() {
	if n == nil { return }
	n.esq.simetrico()
	fmt.Println(n.valor)
	n.dir.simetrico()
}

func (n *No) altura() int {
	if n == nil { return 0 }

	alt1 := n.esq.altura()
	alt2 := n.dir.altura()

	if alt1 > alt2 { return alt1 + 1 }

	return alt2 + 1
}

type Arvore struct {
	raiz *No
}

func (a *Arvore) percurso(tipo string) {
	if a.raiz == nil { fmt.Println("Árvore vazia") }
	if tipo == "pre" { a.raiz.preOrdem() }
	if tipo == "pos" { a.raiz.posOrdem() }
	if tipo == "sim" { a.raiz.simetrico() }
}

func (a *Arvore) altura() int {
	return a.raiz.altura()
}

func main() {
	var arv Arvore
	arv.raiz = &No{valor: 10}

	n1 := &No{valor: 8}
	n2 := &No{valor: 15}
	n3 := &No{valor: 5}
	n4 := &No{valor: 12}
	n5 := &No{valor: 11}

	arv.raiz.esq = n1
	arv.raiz.dir = n2
	n1.esq = n3
	n2.esq = n4
	n4.esq = n5

	arv.percurso("pre")
	fmt.Println("-----------------")
	arv.percurso("pos")
	fmt.Println("-----------------")
	arv.percurso("sim")
	fmt.Println("-----------------")
	fmt.Println(arv.altura())
}
