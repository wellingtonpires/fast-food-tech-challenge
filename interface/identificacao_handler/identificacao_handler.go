package identificacao_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/aggregate/identificacao"
)

func Routes(route *gin.Engine) {
	i := route.Group("/identificacao")
	i.POST("/login", identificacao.Login)
	i.POST("/cadastro", identificacao.Cadastro)
}
