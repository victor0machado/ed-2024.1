package dados

type Pedido struct {
	Id         int     `json:"id"`
	Delivery   bool    `json:"delivery"`
	IdProdutos []int   `json:"id_produtos"`
	ValorTotal float64 `json:"valor_total"`
}

func (p *Pedido) RegistrarID() {
	p.Id = totalIdsPedidos
	totalIdsPedidos++
}
