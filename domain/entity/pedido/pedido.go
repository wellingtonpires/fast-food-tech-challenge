package pedido

type Pedido struct {
    Lanche string  `json:lanche`
    Acompanhamento  string `json:"acompanhamento"`
	Bebida  string `json:"bebida"`
    IdCliente string `json: "idcliente"`
    NumPedido string `json: "numpedido"`
}