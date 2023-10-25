package entrega_handler

import (
    "github.com/gin-gonic/gin"
    "github.com/wellingtonpires/fast-food-tech-challenge/domain/aggregate/entrega"
)

func Routes(route *gin.Engine) {
    listaPedidos := route.Group("/listapedidos")
    listaPedidos.GET("/todos", entrega.ListaPedidos)
}