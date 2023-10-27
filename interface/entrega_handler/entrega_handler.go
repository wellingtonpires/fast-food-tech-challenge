package entrega_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/aggregate/entrega"
)

func Routes(route *gin.Engine) {
	e := route.Group("/listapedidos")
	e.GET("/todos", entrega.ListaPedidos)
}
