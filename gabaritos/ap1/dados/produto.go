package dados

type Produto struct {
	Id        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}

func (p *Produto) RegistrarID() {
	p.Id = totalIdsProdutos
	totalIdsProdutos++
}
