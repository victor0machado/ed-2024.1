package main

import "fmt"

/*
- Inteiros
int8, int16, int32, int64
uint8, uint16, uint32, uint64 -> unsigned
int -> tipo genérico, ocupa 32 ou 64bits, dependendo da arquitetura
uint

byte -> uint8
rune -> int32 -> caractere Unicode

- Ponto flutuante
float32 ou float64
float64 -> padrão

complex64 ou complex128

- Texto
string (cada caractere ocupa 8bits em memória)

- Booleanos
bool -> true ou false

- Outros tipos
	- nil
	- ponteiros
*/

var a int
var b = 15
// c := 10 -> não posso fazer!
const PI = 3.1415 // preciso fazer declaração implícita para constantes

func main() {
	fmt.Println(12.4 + 43.17i)
	msg := `texto
em bloco`

	fmt.Println(msg)

	var x int
	var msg1, msg2 string

	var y = 14.5
	z := 20

	x = 18
	msg1 = "oi"
	msg2 = "olá"
	fmt.Println(x, y, z)
	fmt.Println(msg1, msg2)
}
