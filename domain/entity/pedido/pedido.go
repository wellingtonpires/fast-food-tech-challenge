package pedido

type Pedido struct {
	Lanche         string `json:"lanche"`
	Acompanhamento string `json:"acompanhamento"`
	Bebida         string `json:"bebida"`
	Status         string `json:"status"`
	IdCliente      string `json: "idcliente"`
	CodPedido      string `json: "codpedido"`
}
