package pagamento

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/entity/pedido"
	"github.com/wellingtonpires/fast-food-tech-challenge/infrastructure/persistence"
)

func Confirmacao(c *gin.Context) {
	var pedido pedido.Pedido
	err := c.ShouldBindJSON(&pedido)
	if err != nil {
		fmt.Println(err.Error())
	}
	d := persistence.DadosPedido(pedido)
	pedido.IdCliente = d.IdCliente
	pedido.Lanche = d.Lanche
	pedido.Acompanhamento = d.Acompanhamento
	pedido.Bebida = d.Bebida
	pedido.Status = "Em preparação"
	pedido.CodPedido = d.CodPedido

	persistence.ConfirmaPagamento(pedido)
	c.IndentedJSON(http.StatusOK, gin.H{"resultado": "Pedido pago", "codPedido": pedido.CodPedido})
}
