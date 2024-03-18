package main

import "fmt"

const M = 1000

type Pilha struct {
	pilha [M]rune
	topo  int
}

func (p *Pilha) exibirTopo() (rune, error) {
	if p.topo == -1 { return '0', fmt.Errorf("Pilha vazia") }
	return p.pilha[p.topo], nil
}

func (p *Pilha) push(valor rune) error {
	if p.topo == M-1 { return fmt.Errorf("Overflow") }
	p.topo++
	p.pilha[p.topo] = valor
	return nil
}

func (p *Pilha) pop() (rune, error) {
	if p.topo == -1 { return -1, fmt.Errorf("Underflow") }

	valorRemovido := p.pilha[p.topo]
	p.pilha[p.topo] = '0'
	p.topo--
	return valorRemovido, nil
}

func main() {
	var pilha Pilha
	var n int
	var entrada string

	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		diamantes := 0
		pilha = Pilha{topo: -1}
		fmt.Scanln(&entrada)
		for _, char := range entrada {
			if char == '.' { continue }
			if char == '<' {
				pilha.push(char)
			} else {
				// sei que o char é '>'
				charTopo, _ := pilha.exibirTopo()
				if charTopo == '<' {
					pilha.pop()
					diamantes++
				} else {
					pilha.push(char)
				}
			}
		}
		fmt.Println(diamantes)
	}
}
