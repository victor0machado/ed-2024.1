package dados

import "fmt"

type Produtos struct {
	Lista []Produto
}

func (p *Produtos) inicializar() {
	p.Lista = make([]Produto, 0)
}

func (p *Produtos) Adicionar(nome, descricao string, valor float64) error {
	if nome == "" || descricao == "" { return fmt.Errorf("nome ou descrição vazios") }
	if valor <= 0 { return fmt.Errorf("valor do produto inválido") }

	prod := Produto{
		Nome: nome,
		Descricao: descricao,
		Valor: valor,
	}

	prod.RegistrarID()
	p.Lista = append(p.Lista, prod)
	MetricasColetadas.TotalProdutos++
	return nil
}

func (p *Produtos) Remover(id int) error {
	for indice, prod := range p.Lista {
		if prod.Id == id {
			p.Lista = append(p.Lista[:indice], p.Lista[indice+1:]...)
			MetricasColetadas.TotalProdutos--
			return nil
		}
	}

	return fmt.Errorf("id não encontrado na lista de produtos")
}

func (p *Produtos) Buscar(id int) (Produto, error) {
	for _, prod := range p.Lista {
		if prod.Id == id {
			return prod, nil
		}
	}

	return Produto{}, fmt.Errorf("id não encontrado na lista de produtos")
}

func (p *Produtos) Listar() []Produto {
	return p.Lista
}
