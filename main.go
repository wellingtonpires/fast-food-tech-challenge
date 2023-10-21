package main

import(
	"net/http"
    "github.com/gin-gonic/gin"
    "github.com/wellingtonpires/fast-food-tech-challenge/aggregates"
)

type pedido struct {
    Numero     string  `json:"numero"`
    NomeCliente  string  `json:nomeCliente`
    Lanche string  `json:lanche`
    Acompanhamento  string `json:"acompanhamento"`
	Bebida  string `json:"bebida"`
}

func main() {
    router := gin.Default()
    router.GET("/pedidos", aggregates.listaPedidos)
    router.Run("localhost:8080")
}