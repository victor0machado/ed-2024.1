package main

import "fmt"

type Contato struct {
	Nome, Email string
}

func (c *Contato) alterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func adicionarContato(novoContato Contato, contatos *[5]Contato) {
	for i, contato := range contatos {
		if (contato == Contato{}) {
			contatos[i] = novoContato
			return
		}
	}

	fmt.Println("Agenda cheia!")
}

func excluirContato(contatos *[5]Contato) {
	// Se lista é vazia, sobe erro e retorna
	if (contatos[0] == Contato{}) {
		fmt.Println("Agenda vazia!")
		return
	}

	for i, contato := range contatos {
		if (contato == Contato{}) {
			contatos[i - 1] = Contato{}
			return
		}
	}

	// Caso o for termine sem achar um contato vazio, deve apagar o último item
	contatos[4] = Contato{}
}

// Função apenas para visualizar os dados
func exibirContatos(contatos *[5]Contato) {
	for _, contato := range contatos {
		fmt.Println(contato)
	}
}

func exibirMensagem() {
	fmt.Println("Agenda de contatos!")
	fmt.Println("Selecione uma das opções a seguir:")
	fmt.Println("1 - Adicionar contato;")
	fmt.Println("2 - Excluir contato;")
	fmt.Println("3 - Exibir contatos;")
	fmt.Println("qualquer outra tecla - sair do programa.")
}

func main() {
	var op int
	var contatos [5]Contato

	for {
		var nome, email string

		exibirMensagem()
		fmt.Scanln(&op)
		switch op {
		case 1:
			fmt.Println("Informe o primeiro nome e email, separados por um espaço:")
			fmt.Scanln(&nome, &email)
			adicionarContato(Contato{Nome: nome, Email: email}, &contatos)
		case 2:
			excluirContato(&contatos)
		case 3:
			exibirContatos(&contatos)
		default:
			// o break não funciona pois ele sai do bloco do switch, mas permanece no for
			return
		}
	}
}
