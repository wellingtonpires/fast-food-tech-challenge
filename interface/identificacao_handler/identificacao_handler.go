package identificacao_handler

import (
    "github.com/gin-gonic/gin"
    "github.com/wellingtonpires/fast-food-tech-challenge/domain/aggregate/identificacao"
)

func Routes(route *gin.Engine) {
    ident := route.Group("/identificacao")
    ident.POST("/cadastro", identificacao.Cadastro)
}