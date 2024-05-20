package main

import "fmt"

type No struct {
	valor int
	esq   *No
	dir   *No
}

func (n *No) simetrico() {
	if n == nil { return }
	n.esq.simetrico()
	fmt.Println(n.valor)
	n.dir.simetrico()
}

func (n *No) buscaABB(valor int) *No {
	if n == nil { return nil }

	if n.valor == valor { return n }
	if n.valor < valor { return n.dir.buscaABB(valor) }
	return n.esq.buscaABB(valor)
}

func (n *No) buscaABB2(valor int, noPai *No) (*No, *No) {
	if n == nil { return nil, noPai }

	if n.valor == valor { return n, noPai }
	if n.valor < valor { return n.dir.buscaABB2(valor, n) }
	return n.esq.buscaABB2(valor, n)
}

type Arvore struct {
	raiz *No
}

func (a *Arvore) percurso() {
	if a.raiz == nil { fmt.Println("Árvore vazia") }
	a.raiz.simetrico()
}

func (a *Arvore) buscaABB(valor int) *No {
	if a.raiz == nil { return nil }
	return a.raiz.buscaABB(valor)
}

func (a *Arvore) buscaABB2(valor int) (*No, *No) {
	if a.raiz == nil { return nil, nil }
	return a.raiz.buscaABB2(valor, nil)
}

func (a *Arvore) insercao(valor int) error {
	novoNo := &No{valor: valor}
	if a.raiz == nil {
		a.raiz = novoNo
		return nil
	}

	if a.buscaABB(valor) != nil { return fmt.Errorf("nó já existe na árvore") }

	no := a.raiz
	for no != nil {
		if no.valor < valor {
			if no.dir == nil {
				no.dir = novoNo
				break
			}
			no = no.dir
		} else {
			if no.esq == nil {
				no.esq = novoNo
				break
			}
			no = no.esq
		}
	}

	return nil
}

func (a *Arvore) remover(valor int) error {
	no, noPai := a.buscaABB2(valor)
	if no == nil { return fmt.Errorf("nó não existe") }

	if no.esq == nil && no.dir == nil {
		if noPai.valor < valor {
			noPai.dir = nil
		} else {
			noPai.esq = nil
		}
	} else if no.esq == nil {
		if noPai.valor < valor {
			noPai.dir = no.dir
		} else {
			noPai.esq = no.dir
		}
	} else if no.dir == nil {
		if noPai.valor < valor {
			noPai.dir = no.esq
		} else {
			noPai.esq = no.esq
		}
	} else {
		x := no
		n := no.dir
		for n.esq != nil {
			x = n
			n = n.esq
		}

		if noPai.valor < valor {
			noPai.dir = n
		} else {
			noPai.esq = n
		}
		n.esq = no.esq
		n.dir = no.dir
		x.esq = nil
	}

	return nil
}

func main() {
	var arv Arvore

	arv.percurso()

	fmt.Println(arv.buscaABB(5))
	arv.insercao(11)
	arv.insercao(7)
	arv.insercao(9)
	arv.insercao(2)
	arv.insercao(33)
	arv.insercao(20)
	arv.insercao(40)
	arv.insercao(37)
	arv.insercao(55)
	arv.insercao(35)
	arv.insercao(34)
	arv.insercao(36)
	arv.percurso()

	fmt.Println(arv.buscaABB(33))
	fmt.Println(arv.buscaABB2(33))

	arv.remover(33)
	arv.percurso()
}
