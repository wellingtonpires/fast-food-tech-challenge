package pedido

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/entity/produto"
	"github.com/wellingtonpires/fast-food-tech-challenge/infrastructure/persistence"
)

func Cadastro(c *gin.Context) {
	var prod produto.Produto
	err := c.ShouldBindJSON(&prod)
	if err != nil {
		fmt.Println(err.Error())
	}
	persistence.CadastraProduto(prod)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Produto criado!"})
}

func Atualizacao(c *gin.Context) {
	var prod produto.Produto
	err := c.ShouldBindJSON(&prod)
	if err != nil {
		fmt.Println(err.Error())
	}
	persistence.AtualizaProduto(prod)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Produto atualizado!"})
}

func Exclusao(c *gin.Context) {
	var prod produto.Produto
	err := c.ShouldBindJSON(&prod)
	if err != nil {
		fmt.Println(err.Error())
	}
	persistence.ExcluiProduto(prod)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Produto deletado!"})
}

func ConsultaCategoria(c *gin.Context) {
	categoria := c.Query("categoria")
	c.IndentedJSON(http.StatusOK, persistence.ConsultaProdutoPorCategoria(categoria))
}
