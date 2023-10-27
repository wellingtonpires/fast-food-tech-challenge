package identificacao

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellingtonpires/fast-food-tech-challenge/domain/entity/cliente"
	"github.com/wellingtonpires/fast-food-tech-challenge/infrastructure/persistence"
)

func Cadastro(c *gin.Context) {
	var cli cliente.Cliente
	err := c.ShouldBindJSON(&cli)
	if err != nil {
		fmt.Println(err.Error())
	}
	persistence.CadastraCliente(cli)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Cliente criado!"})
}

func Login(c *gin.Context) {
	clienteOk := false
	var cli cliente.Cliente
	err := c.ShouldBindJSON(&cli)
	if err != nil {
		fmt.Println(err.Error())
	}
	dadosCliente := persistence.DadosCliente(cli)

	for _, d := range dadosCliente {
		if cli.Cpf == d.Cpf && cli.Senha == d.Senha {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Acesso OK"})
			clienteOk = true
		}
	}

	if !clienteOk {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Não autorizado"})
	}

}
