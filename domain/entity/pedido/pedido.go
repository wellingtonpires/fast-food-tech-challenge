package pedido

type Pedido struct {
	Lanche           string `json:"lanche,omitempty"`
	Acompanhamento   string `json:"acompanhamento,omitempty"`
	Bebida           string `json:"bebida,omitempty"`
	StatusPagamento  string `json:"statuspagamento,omitempty"`
	StatusPreparacao string `json:"statuspreparacao,omitempty"`
	CodPedido        string `json:"codpedido,omitempty"`
	IdCliente        string `json:"idcliente,omitempty"`
}
