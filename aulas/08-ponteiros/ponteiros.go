package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 2
	y := x

	fmt.Println(x, y)

	x = 3
	fmt.Println(x, y)

	z := &x // referência
	fmt.Println(x, z)
	fmt.Println(x, *z) // dereferência

	x = 4
	fmt.Println(x, z)
	fmt.Println(x, *z) // dereferência

	msg1 := "olá, mundo"
	alteraMensagem(&msg1)
	fmt.Println(msg1)

	n1 := Numero{Valor: 10}
	fmt.Println(n1) // 10
	n1.Incremento()
	fmt.Println(n1) // ?
}

type Numero struct {
	Valor int
}

func (n *Numero) Incremento() {
	n.Valor++
}

/*
Usamos ponteiros como parâmetros de funções quando:
1. é necessário reduzir o espaço consumido em memória;
2. queremos alterar o valor dos parâmetros
*/
func alteraMensagem(msg *string) {
	novaMensagem := strings.ReplaceAll(*msg, "mundo", "turma")
	*msg = novaMensagem
}
