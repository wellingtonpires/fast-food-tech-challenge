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
	var baseClientes []cliente.Cliente
	clienteExiste := true
	err := c.ShouldBindJSON(&cli)
	if err != nil {
		fmt.Println(err.Error())
	}

	baseClientes = persistence.DadosCliente()

	for _, b := range baseClientes {
		if b.Cpf == cli.Cpf {
			clienteExiste = false
		}
	}

	if clienteExiste {
		persistence.CadastraCliente(cli)
		c.IndentedJSON(http.StatusCreated, gin.H{"resultado": "Cliente criado"})
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"resultado": "CPF já cadastrado na base"})
	}

}

func Login(c *gin.Context) {
	clienteLogin := false
	var cli cliente.Cliente
	err := c.ShouldBindJSON(&cli)
	if err != nil {
		fmt.Println(err.Error())
	}
	dadosCliente := persistence.DadosCliente()

	for _, d := range dadosCliente {
		if cli.Cpf == d.Cpf && cli.Senha == d.Senha {
			c.IndentedJSON(http.StatusOK, gin.H{"resultado": "Acesso OK"})
			clienteLogin = true
		}
	}

	if !clienteLogin {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"resultado": "Não autorizado"})
	}

}
