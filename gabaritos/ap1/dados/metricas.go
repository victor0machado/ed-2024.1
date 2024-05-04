package dados

type Metricas struct {
	TotalProdutos      int     `json:"total_produtos"`
	PedidosEncerrados  int     `json:"pedidos_encerrados"`
	PedidosEmAndamento int     `json:"pedidos_em_andamento"`
	FaturamentoTotal   float64 `json:"faturamento_total"`
}
