package processamento

import (
	"api/dados"
	"time"
)

func ProcessaPedidos() {
	for {
		if dados.LojaEstaAberta() {
			time.Sleep(30 * time.Second)
			dados.FilaPedidos.Expedir()
		}
	}
}
