package pagamento_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/aggregate/pagamento"
)

func Routes(route *gin.Engine) {
	p := route.Group("/pagamento")
	p.PATCH("/confirmacao", pagamento.Confirmacao)
}
