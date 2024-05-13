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

func (n *No) busca(valor int) *No {
	if n == nil { return nil }
	if n.valor == valor { return n }

	noProcurado := n.esq.busca(valor)
	if noProcurado != nil { return noProcurado }

	return n.dir.busca(valor)
}

func (n *No) buscaPai(noPai *No, valor int) (*No, *No) {
	if n == nil { return nil, noPai }
	if n.valor == valor { return n, noPai }

	noProcurado, noPai := n.esq.buscaPai(n, valor)
	if noProcurado != nil { return noProcurado, noPai }

	return n.dir.buscaPai(n, valor)
}

type Arvore struct {
	raiz *No
}

func (a *Arvore) excluir(valor int) error {
	if a.raiz == nil { return fmt.Errorf("Árvore vazia") }

	noProcurado, noPai := a.raiz.buscaPai(nil, valor)
	if noProcurado == nil { return fmt.Errorf("Nó não se encontra na árvore") }

	if noPai == nil {
		a.raiz = nil
	} else if noPai.esq == noProcurado {
		noPai.esq = nil
	} else {
		noPai.dir = nil
	}

	return nil
}

func (a *Arvore) inserir(valorInserir, valorPai int, posicao string) error {
	if a.raiz == nil { return fmt.Errorf("Árvore vazia") }

	if a.busca(valorInserir) != nil { return fmt.Errorf("Árvore já possui esse valor") }

	noPai := a.busca(valorPai)
	if noPai == nil { return fmt.Errorf("O nó pai não existe na árvore") }

	switch posicao {
	case "esq":
		if noPai.esq != nil { return fmt.Errorf("O nó pai já possui filho à esquerda") }
		noPai.esq = &No{valor: valorInserir}
	case "dir":
		if noPai.dir != nil { return fmt.Errorf("O nó pai já possui filho à direita") }
		noPai.dir = &No{valor: valorInserir}
	default:
		return fmt.Errorf("Posição informada não é válida.")
	}

	return nil
}

func (a *Arvore) busca(valor int) *No {
	if a.raiz == nil { return nil }
	return a.raiz.busca(valor)
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

	// arv.percurso("pre")
	// fmt.Println("-----------------")
	// arv.percurso("pos")
	// fmt.Println("-----------------")
	// arv.percurso("sim")
	// fmt.Println("-----------------")
	// fmt.Println(arv.altura())

	// fmt.Println(arv.busca(8))
	// fmt.Println(arv.busca(25))

	fmt.Println(arv.inserir(13, 5, "esq")) // ok
	fmt.Println(arv.inserir(22, 5, "esq")) // já tem filho à esquerda
	fmt.Println(arv.inserir(22, 6, "esq")) // nó pai não existe
	fmt.Println(arv.inserir(13, 5, "dir")) // nó já existe
	fmt.Println(arv.inserir(22, 5, "abc")) // posição inválida
	fmt.Println(arv.inserir(22, 10, "dir")) // já tem filho à direita

	arv.percurso("sim")

	arv.excluir(10)
	fmt.Println("-----------------")
	arv.percurso("sim")
}
