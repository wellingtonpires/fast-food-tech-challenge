package pedido

type Pedido struct {
	Lanche         string `json:"lanche"`
	Acompanhamento string `json:"acompanhamento"`
	Bebida         string `json:"bebida"`
	Status         string `json:"status"`
	CodPedido      string `json:"codpedido"`
	IdCliente      string `json:"idcliente"`
}
