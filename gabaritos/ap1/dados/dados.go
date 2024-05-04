package dados

import "fmt"

var ListaProdutos Produtos
var FilaPedidos Pedidos
var totalIdsProdutos, totalIdsPedidos = 1, 1
var lojaAberta = false
var MetricasColetadas Metricas

func InicializarDados() {
	ListaProdutos = Produtos{}
	ListaProdutos.inicializar()

	FilaPedidos = Pedidos{}
	FilaPedidos.inicializar()

	MetricasColetadas = Metricas{
		TotalProdutos:      0,
		PedidosEncerrados:  0,
		PedidosEmAndamento: 0,
		FaturamentoTotal:   0.0,
	}
}

func AbrirLoja() error {
	if lojaAberta {
		return fmt.Errorf("loja já está aberta")
	}

	lojaAberta = true
	return nil
}

func FecharLoja() error {
	if !lojaAberta {
		return fmt.Errorf("loja já está fechada")
	}

	lojaAberta = false
	return nil
}

func LojaEstaAberta() bool {
	return lojaAberta
}
