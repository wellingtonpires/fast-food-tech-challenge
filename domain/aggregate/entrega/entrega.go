package entrega


import(
	"net/http"
    "github.com/gin-gonic/gin"
    "github.com/wellingtonpires/fast-food-tech-challenge/infrastructure/persistence"
)



func ListaPedidos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persistence.BuscaPedidos())
}
