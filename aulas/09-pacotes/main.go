package main

import (
	"fmt"
	m "projeto/modelos" // python -> from projeto import modelos as m
	"projeto/modelos/academico"
)

func main() {
	fmt.Println("Olá, mundo")

	aluno := m.Aluno{}
	aluno.Nome = "abcd"
	aluno.Matricula = "1234"

	curso := academico.Curso{Nome: "Engenharia"}

	fmt.Println(aluno, curso)
}
