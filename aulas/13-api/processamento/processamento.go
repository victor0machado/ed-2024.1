package processamento

import (
	"api/dados"
	"fmt"
	"time"
)

func ProcessaDados() {
	for {
		time.Sleep(10 * time.Second)
		if len(dados.Dados) == 0 {
			fmt.Println("Lista de dados está vazia!")
		} else {
			fmt.Println("-----------------------")
			fmt.Println("Total de mensagens:", len(dados.Dados))
			for _, mensagem := range dados.Dados {
				fmt.Println(mensagem)
			}
			fmt.Println("-----------------------")
		}
	}
}
