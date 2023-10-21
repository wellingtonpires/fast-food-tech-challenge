package entrega

import(
	"net/http"
    "github.com/gin-gonic/gin"
)


func ListaPedidos(c *gin.Context) {
    //MOCK
    var pedidos = []pedido{
        {Numero: "123", NomeCliente: "Wellington Pires", Lanche: "cheeseburguer", Acompanhamento: "batata frita", Bebida: "coca-cola"},
        {Numero: "124", NomeCliente: "Wellington Pires", Lanche: "cheesebacon", Acompanhamento: "nuggets", Bebida: "fanta laranja"},
    }
	c.IndentedJSON(http.StatusOK, pedidos)
}