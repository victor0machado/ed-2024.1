package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Lista struct {
	tamanho int
	inicio  *No
	fim     *No
}

// Percorrer a lista
func (l *Lista) percorrer() {
	no := l.inicio
	if no == nil {
		fmt.Println("Lista vazia!")
	} else {
		for no.prox != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println(no.valor)
	}
}

// Considerando que não há um campo fim no tipo Lista
// Complexidade O(n)
func (l *Lista) inserirFinal(novoValor int) {
	no := &No{valor: novoValor}
	if l.inicio == nil {
		l.inicio = no
	} else {
		no_prox := l.inicio
		for no_prox.prox != nil {
			no_prox = no_prox.prox
		}
		no_prox.prox = no
	}

	l.tamanho++
}

// Considerando o campo fim na Lista, podemos ter inserção com O(1)
func (l *Lista) inserirFinalO1(novoValor int) {
	no := &No{valor: novoValor}
	if l.inicio == nil {
		l.inicio = no
		l.fim = no
	} else {
		l.fim.prox = no
		l.fim = no
	}
	l.tamanho++
}

func (l *Lista) inserirMeio(novoValor, indice int) error {
	if l.tamanho < indice {
		return fmt.Errorf("lista não comporta índice %d", indice)
	}

	if l.tamanho == indice {
		l.inserirFinalO1(novoValor)
	} else {
		cont := 0
		no := l.inicio
		for cont < indice - 1 {
			no = no.prox
			cont++
		}

		novoNo := &No{valor: novoValor}
		novoNo.prox = no.prox
		no.prox = novoNo

		l.tamanho++
	}

	return nil
}

func (l *Lista) excluir(valorAExcluir int) error {
	if l.inicio == nil {
		return fmt.Errorf("Lista vazia!")
	}

	no := l.inicio
	if no.valor == valorAExcluir {
		l.inicio = no.prox
	} else {
		for no != nil {
			if no.prox != nil {
				if no.prox.valor == valorAExcluir {
					no.prox = no.prox.prox
					break
				} else {
					no = no.prox
				}
			} else {
				return fmt.Errorf("Valor não encontrado")
			}
		}
	}

	if no.prox == nil {
		l.fim = no
	}
	l.tamanho--
	return nil
}

func main() {
	var lista Lista
	lista.percorrer()

	// n1 := &No{valor: 10}
	// n2 := &No{valor: 16}
	// lista.inicio = n1
	// n1.prox = n2
	// lista.inserirFinal(10)
	// lista.inserirFinal(16)
	// lista.inserirFinal(23)
	lista.inserirFinalO1(10)
	lista.inserirFinalO1(16)
	lista.inserirFinalO1(23)

	lista.inserirMeio(13, 1)
	lista.inserirMeio(17, 3)
	lista.inserirMeio(27, 5)

	lista.percorrer()

	lista.excluir(16)
	lista.excluir(10)
	lista.excluir(27)

	lista.percorrer()

	fmt.Println(lista.fim.valor)
	fmt.Println(lista.tamanho)
}
