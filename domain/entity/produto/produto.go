package produto

type Produto struct {
	Id        string  `json:id`
	Nome      string  `json:nome`
	Valor     float32 `json:valor`
	Categoria string  `json:categoria`
	Descricao string  `json:descricao`
}
