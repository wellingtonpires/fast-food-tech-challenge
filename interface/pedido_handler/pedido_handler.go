package pedido_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/aggregate/pedido"
)

func Routes(route *gin.Engine) {
	p := route.Group("/pedido")
	p.POST("/cadastra-menu", pedido.Cadastro)
	p.PATCH("/atualiza-menu", pedido.Atualizacao)
	p.DELETE("/exclui-menu", pedido.Exclusao)
	p.GET("/consulta", pedido.ConsultaCategoria)
	p.POST("/novo-pedido", pedido.NovoPedido)
	p.GET("/consulta-pagamento", pedido.ConsultaPagamento)
}
