package geometria

type Retangulo struct {
	Largura, Altura float64
}

func (r Retangulo) CalcularArea() float64 {
	return r.Altura * r.Largura
}

func (r Retangulo) CalcularPerimetro() float64 {
	return 2 * (r.Altura + r.Largura)
}
